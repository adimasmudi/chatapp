package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository interface {
	SaveRoom(ctx context.Context,roomName string) (*mongo.InsertOneResult, error)
}

type chatRepository struct{
	DB *mongo.Collection
}

func NewChatRepository(DB *mongo.Collection) *chatRepository{
	return &chatRepository{DB}
}

func (r *userRepository) SaveRoom(ctx context.Context,roomName string) (*mongo.InsertOneResult, error) {
	
	result,err := r.DB.InsertOne(ctx, roomName)

	if err != nil {
		return result, err
	}

	return result, nil
}