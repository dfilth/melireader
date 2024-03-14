package port

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"melireader/internal/adapter/config"
	"melireader/internal/core/domain"
	"melireader/internal/core/models"
	"mime/multipart"
	"strings"
)

//go:generate mockgen -source=item.go -destination=mock/item.go -package=mock

type ItemRepository interface {
	InsertItem(ctx context.Context, item *domain.Item) (*domain.Item, error)
	InsertItemsBatch(ctx context.Context, items []*domain.Item) error
	GetItems(ctx context.Context, page int, pageSize int) ([]*domain.Item, error)
}

type ItemService interface {
	Register(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*domain.Item, error)
	GetItemsPage(ctx context.Context, page int, pageSize int) ([]*domain.Item, error)
}

type FileReader interface {
	ReadLine(reader *bufio.Reader) (string, error)
}

type CSVFileReader struct {
	Config config.FileReaderConfig
}

func (c *CSVFileReader) ReadLine(reader *bufio.Reader) (string, error) {
	csvConfig, ok := c.Config.Details.(config.CSVConfig)
	if !ok {
		return "", errors.New("invalid configuration details for CSV file reader")
	}

	delimiterByte := []byte(csvConfig.Delimiter)[0]

	for {
		line, err := reader.ReadString(delimiterByte)
		if err != nil && err != io.EOF {
			return "", err
		}

		line = strings.TrimSuffix(line, string(delimiterByte))

		if len(strings.TrimSpace(line)) > 0 {
			return line, nil
		}

		if err == io.EOF {
			return "", io.EOF
		}
	}
}

type JSONLFileReader struct {
	Config config.FileReaderConfig
}

func (j *JSONLFileReader) ReadLine(reader *bufio.Reader) (string, error) {
	jsonLConfig, ok := j.Config.Details.(config.JSONLConfig)
	if !ok {
		return "", errors.New("invalid configuration details for JsonL file reader")
	}

	delimiterByte := []byte(jsonLConfig.Delimiter)[0]
	line, err := reader.ReadString(delimiterByte)
	if err != nil {
		return "", err
	}

	var item models.ItemLine

	err = json.Unmarshal([]byte(line), &item)
	if err != nil {
		log.Fatalf("Error al deserializar el JSON: %v", err)
	}

	return fmt.Sprintf("%s, %d", item.Site, item.ID), nil
}

type TXTFileReader struct {
	Config config.FileReaderConfig
}

func (t *TXTFileReader) ReadLine(reader *bufio.Reader) (string, error) {
	txtConfig, ok := t.Config.Details.(config.TXTConfig)
	if !ok {
		return "", errors.New("invalid configuration details for TXT file reader")
	}

	line, err := reader.ReadString(txtConfig.Separator[0])
	if err != nil && err != io.EOF {
		return "", err
	}
	line = strings.TrimSuffix(line, txtConfig.Separator)
	return line, nil
}
