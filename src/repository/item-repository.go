package repository

import (
	"context"
	"time"

	"github.com/aom31/GO-Inventory/src/constants"
	"github.com/aom31/GO-Inventory/src/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	Client *mongo.Client
}

func (repo *ItemRepository) getCollectionDb() *mongo.Collection {
	return repo.Client.Database(constants.DBNAME_ITEM).Collection(constants.Collection_ITEMS)
}

func (repo *ItemRepository) FindItems(ctx context.Context, items *[]models.Item) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	cursor, err := repo.getCollectionDb().Find(ctx, bson.D{}, nil)
	if err != nil {
		return errors.Wrap(err, "cannot find item")
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item models.Item
		if err := cursor.Decode(&item); err != nil {
			return errors.Wrap(err, "failed decode item")
		}
		*items = append(*items, item)
	}

	return nil
}

func (repo *ItemRepository) FindOneItem(ctx context.Context, itemId string) (*models.Item, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	objectItemId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectItemId,
	}
	item := new(models.Item)
	if err := repo.getCollectionDb().FindOne(ctx, filter, nil).Decode(&item); err != nil {
		return nil, errors.Wrap(err, "fail to decode item find with itemid")
	}

	return item, nil
}
