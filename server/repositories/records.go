package repositories

import (
	"context"
	"github.com/mateusrlopez/url-shortener-server/models"
	"github.com/mateusrlopez/url-shortener-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Records struct {
	collection *mongo.Collection
}

func NewRecords(collection *mongo.Collection) *Records {
	return &Records{
		collection: collection,
	}
}

func (r Records) Insert(url string) (string, error) {
	record := models.Record{
		Key:       utils.GenerateUniqueKey(),
		URL:       url,
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, &record)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}

func (r Records) FindOneByKey(key string) (models.Record, error) {
	var record models.Record

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"_id": bson.M{"$eq": key}}

	if err := r.collection.FindOne(ctx, filter).Decode(&record); err != nil {
		return models.Record{}, err
	}

	return record, nil
}

func (r Records) FindOneAndUpdateAccessesByKey(key string) (models.Record, error) {
	var record models.Record

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"_id": bson.M{"$eq": key}}
	update := bson.M{
		"$set": bson.M{"last_access": time.Now()},
		"$inc": bson.M{"accesses": 1},
	}

	after := options.After
	opts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := r.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if err := result.Err(); err != nil {
		return models.Record{}, err
	}

	if err := result.Decode(&record); err != nil {
		return models.Record{}, err
	}

	return record, nil
}
