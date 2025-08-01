package persistence

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"api/internal/domain"
)

type MongoEventRepository struct {
	collection *mongo.Collection
}

func NewMongoEventRepository(db *mongo.Database, collectionName string) *MongoEventRepository {
	collection := db.Collection(collectionName)
	return &MongoEventRepository{collection: collection}
}

func (r *MongoEventRepository) Save(ctx context.Context, event *domain.Event) error {
	_, err := r.collection.InsertOne(ctx, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *MongoEventRepository) FindByID(ctx context.Context, id string) (*domain.Event, error) {
	var event domain.Event
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&event)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &event, nil
}

func (r *MongoEventRepository) FindAll(ctx context.Context) ([]*domain.Event, error) {
	var events []*domain.Event
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var event domain.Event
		err := cur.Decode(&event)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}