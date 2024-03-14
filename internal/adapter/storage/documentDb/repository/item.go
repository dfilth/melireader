package repository

import (
	"awesomeProject/internal/core/domain"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "log"
	_ "time"
)

type ItemRepository struct {
	collection *mongo.Collection
}

func NewItemRepository(db *mongo.Database) *ItemRepository {
	return &ItemRepository{
		collection: db.Collection("items"),
	}
}

func (i *ItemRepository) InsertItem(ctx context.Context, item *domain.Item) (*domain.Item, error) {
	_, err := i.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, errors.Wrap(err, "error inserting item into collection")
	}
	return item, nil
}

func (repo *ItemRepository) InsertItemsBatch(ctx context.Context, items []*domain.Item) error {
	var documents []interface{}
	for _, item := range items {
		documents = append(documents, item)
	}

	_, err := repo.collection.InsertMany(ctx, documents)
	return err
}

func (i *ItemRepository) GetItems(ctx context.Context, page int, pageSize int) ([]*domain.Item, error) {
	// Calcular el índice de inicio y el límite de la página
	startIndex := (page - 1) * pageSize
	limit := pageSize

	// Definir la consulta para la paginación
	options := options.Find().SetSkip(int64(startIndex)).SetLimit(int64(limit))

	cursor, err := i.collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching items from collection")
	}
	defer cursor.Close(ctx)

	// Iterar sobre el cursor y almacenar los resultados en un slice
	var items []*domain.Item
	for cursor.Next(ctx) {
		var item domain.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, errors.Wrap(err, "error decoding item from cursor")
		}
		items = append(items, &item)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}

	return items, nil
}
