package service

import (
	"awesomeProject/internal/adapter/config"
	"awesomeProject/internal/core/domain"
	"awesomeProject/internal/core/models"
	"awesomeProject/internal/core/port"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"
)

type ItemService struct {
	repo     port.ItemRepository
	fileType port.FileReader
}

func NewItemService(repo port.ItemRepository, svc port.FileReader) *ItemService {
	return &ItemService{
		repo,
		svc,
	}
}

const (
	bufferedChannels = 1000
	workersNum       = 10
	batchSize        = 1000
)

func getFileType(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func getFileName(fileHeader *multipart.FileHeader) string {
	return fileHeader.Filename
}

func (is *ItemService) Register(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*domain.Item, error) {

	startTime := time.Now()

	fileName := getFileName(fileHeader)
	fileType := getFileType(fileName)
	var fileReader port.FileReader

	switch fileType {
	case "csv":
		fileReader = &port.CSVFileReader{
			Config: config.FileReaderConfig{
				Format:   "csv",
				Encoding: "utf-8",
				Details: config.CSVConfig{
					Separator: ",",
					Delimiter: "\n",
				},
			},
		}
	case "jsonl":
		fileReader = &port.JSONLFileReader{
			Config: config.FileReaderConfig{
				Format:   "jsonl",
				Encoding: "utf-8",
				Details: config.JSONLConfig{
					Delimiter: "\n",
				},
			},
		}
	case "txt":
		fileReader = &port.TXTFileReader{
			Config: config.FileReaderConfig{
				Format:   "txt",
				Encoding: "utf-8",
				Details: config.TXTConfig{
					Separator: ",",
				},
			},
		}
	default:
		return nil, errors.New("unsupported file format")
	}

	reader := bufio.NewReader(file)

	// Crear canales para enviar datos leídos y procesados
	linesChannel := make(chan string)
	processedDataChannel := make(chan *domain.Item)
	done := make(chan struct{})

	// Iniciar el productor (lector del archivo) en una goroutine
	go func() {
		defer close(linesChannel)
		for {
			line, err := fileReader.ReadLine(reader) //reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println("Error al leer línea del archivo:", err)
				continue
			}
			linesChannel <- line
		}
	}()

	var wg sync.WaitGroup

	// Iniciar el pool de workers
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processLines(linesChannel, processedDataChannel)
		}()
	}

	// Gorutina para esperar a que todas las gorutinas de trabajo terminen
	go func() {
		wg.Wait()
		close(processedDataChannel)
		close(done)
	}()

	processAndInsert(ctx, is, processedDataChannel, batchSize)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	hours := int(elapsedTime.Hours())
	minutes := int(elapsedTime.Minutes()) % 60
	seconds := int(elapsedTime.Seconds()) % 60

	log.Printf("Tiempo transcurrido: %02d:%02d:%02d", hours, minutes, seconds)
	return nil, nil
}

func (is *ItemService) GetItemsPage(ctx context.Context, page int, pageSize int) ([]*domain.Item, error) {
	items, err := is.repo.GetItems(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (is *ItemService) InsertItemsBatch(ctx context.Context, items []*domain.Item) error {
	// Aquí ya no necesitas transformar items a un slice de interface{}
	// porque InsertItemsBatch espera []*domain.Item

	err := is.repo.InsertItemsBatch(ctx, items)
	if err != nil {
		log.Printf("Error al insertar el lote: %v\n", err)
		return err
	}

	return nil
}

func processAndInsert(ctx context.Context, service *ItemService, processedDataChannel <-chan *domain.Item, batchSize int) {
	var batch []*domain.Item
	var wg sync.WaitGroup

	// Función auxiliar para insertar un batch y manejar la concurrencia.
	flushBatch := func(batchToInsert []*domain.Item) {
		wg.Add(1)
		go func(batch []*domain.Item) {
			defer wg.Done()
			// Intentar insertar el lote en la base de datos.
			if err := service.InsertItemsBatch(ctx, batch); err != nil {
				// Si ocurre un error al insertar el lote, se registra el error.
				log.Printf("Error al insertar un batch: %v\n", err)
			} else {
				// Si la inserción es exitosa, se registra un mensaje indicándolo.
				log.Printf("Lote de %d ítems insertado satisfactoriamente.\n", len(batch))
			}
		}(batchToInsert)
	}

	for item := range processedDataChannel {
		batch = append(batch, item)
		if len(batch) >= batchSize {
			flushBatch(batch)
			batch = nil // Crear un nuevo batch
		}
	}

	// Asegurarse de insertar cualquier dato restante
	if len(batch) > 0 {
		flushBatch(batch)
	}

	wg.Wait() // Esperar a que todas las inserciones en lote terminen
}

func processLines(input <-chan string, output chan<- *domain.Item) {
	for line := range input {
		processedData, err := processLine(line)
		if err != nil {
			log.Println("Error al procesar la línea:", err)
			continue
		}
		if processedData != nil {
			output <- processedData
		}
	}
}

func processLine(line string) (*domain.Item, error) {
	parts := strings.Split(line, ",")

	if len(parts) >= 2 {
		site := strings.TrimSpace(parts[0])
		itemID := strings.TrimSpace(parts[1])

		// Verificar si la línea contiene encabezados de columna
		if itemID == "id" && site == "site" {
			// Ignorar la primera línea y devolver nil para evitar procesarla
			return nil, nil
		}

		info, err := getItemInfo(site, itemID)
		if err != nil {
			fmt.Println("No se pudo obtener información para el ítem", itemID, "del sitio", site, ":", err)
			return nil, nil // Continuar con la siguiente línea
		}

		return info, nil
	} else {
		return nil, errors.New("Error: La línea no tiene el formato esperado.")
	}
}

func getItemInfo(site string, itemID string) (*domain.Item, error) {
	itemResponse, err := getItemMELI(site, itemID)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	price := itemResponse.Price
	startTime := itemResponse.StartTime
	categoryID := itemResponse.CategoryID
	currencyID := itemResponse.CurrencyID
	sellerID := itemResponse.SellerID

	categoryName, err := getCategoryName(categoryID)
	if err != nil {
		return nil, err
	}

	var currencyDesc = "Moneda Desconocida"

	if currencyID != "" {
		currencyDesc, err = getCurrencyDescription(currencyID)
		if err != nil {
			log.Println("Error al obtener la descripción de la moneda:", err)
		}
	}

	sellerNickname, err := getSellerNickname(sellerID)
	if err != nil {
		return nil, err
	}

	itemInfo := &domain.Item{
		Site:      site,
		Id:        itemID,
		StartTime: startTime,
		Price:     price,
		Category:  categoryName,
		Currency:  currencyDesc,
		Seller:    sellerNickname,
	}

	return itemInfo, nil
}

func getItemMELI(site, itemID string) (*models.ItemResponse, error) {
	itemKey := site + itemID
	url := fmt.Sprintf("https://api.mercadolibre.com/items/%s", itemKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, errors.New("Ítem no encontrado")
	}

	var item models.ItemResponse
	if err := json.NewDecoder(response.Body).Decode(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

func getCategoryName(categoryID string) (string, error) {
	// Realiza la llamada a la API de MercadoLibre para obtener la categoría
	url := fmt.Sprintf("https://api.mercadolibre.com/categories/%s", categoryID)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Decodifica el JSON de la respuesta en la estructura CategoryResponse
	var category models.CategoryResponse
	if err := json.NewDecoder(response.Body).Decode(&category); err != nil {
		return "", err
	}

	return category.Name, nil
}

// Función para obtener la descripción de una moneda de MercadoLibre
func getCurrencyDescription(currencyID string) (string, error) {
	// Realiza la llamada a la API de MercadoLibre para obtener la moneda
	url := fmt.Sprintf("https://api.mercadolibre.com/currencies/%s", currencyID)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	var currency models.CurrencyResponse
	if err := json.NewDecoder(response.Body).Decode(&currency); err != nil {
		return "", err
	}

	return currency.Description, nil
}

func getSellerNickname(userID int) (string, error) {
	url := fmt.Sprintf("https://api.mercadolibre.com/users/%d", userID)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	var user models.UserResponse
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return "", err
	}

	return user.Nickname, nil
}
