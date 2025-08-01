package persistence

import (
	"api/internal/domain"
	"api/internal/ports"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

// Ensure MongoUserRepository implements ports.UserRepository
var _ ports.UserRepository = (*MongoUserRepository)(nil)

func NewMongoUserRepository(db *mongo.Database, collectionName string) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoUserRepository) Save(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if user.ID == "" {
		user.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, bson.M{
		"_id":      user.ID,
		"name":     user.Name,
		"email":    user.Email,
		"whatsapp": user.WhatsApp,
		"password": user.Password,
	})
	return err
}

func (r *MongoUserRepository) FindByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *MongoUserRepository) FindAll() ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*domain.User
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, cursor.Err()
}

func (r *MongoUserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}
