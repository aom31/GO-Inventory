package repository

import (
	"context"
	"errors"

	"log"
	"time"

	"github.com/aom31/GO-Inventory/src/constants"
	"github.com/aom31/GO-Inventory/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Client *mongo.Client
}

func (repo *UserRepository) FindOneUser(ctx context.Context, userId string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	filterPipeline := filterQueryPipeline(userId)

	cursor, err := repo.Client.Database(constants.DBNAME_USER).Collection(constants.Collection_USERS).Aggregate(ctx, filterPipeline, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// decode result query to struct
	var userItem *models.UserItemID
	if cursor.Next(ctx) {
		if err := cursor.Decode(&userItem); err != nil {
			return nil, errors.New("user not found")
		}
	}

	// map result from db to struct user
	user := mapUserItemToUser(userItem)

	return user, nil
}

func filterQueryPipeline(userId string) bson.A {
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}

	pipeline := bson.A{
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "users_items"},
					{"localField", "_id"},
					{"foreignField", "user_id"},
					{"as", "itemofuser"},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 1},
					{"username", 1},
					{"itemofuser", "$itemofuser.item_id"},
				},
			},
		},
		bson.D{{"$match", bson.D{{"_id", objectID}}}},
	}

	return pipeline
}

func mapUserItemToUser(userItem *models.UserItemID) *models.User {
	items := make([]models.ItemId, 0)

	for _, idItem := range userItem.ItemOfUser {
		items = append(items, models.ItemId{
			ObjectId: idItem,
		})
	}

	user := models.User{
		Id:       userItem.Id,
		Username: userItem.Username,
		Items:    items,
	}

	return &user
}
