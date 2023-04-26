package repositories

import (
	"chatapp/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	FindByUsername(ctx context.Context,username string) (models.User, error)
	Save(ctx context.Context,user models.User) (*mongo.InsertOneResult, error)
	IsUserExist(ctx context.Context, username string) (bool, error)
}

type userRepository struct{
	DB *mongo.Collection
}

func NewUserRepository(DB *mongo.Collection) *userRepository{
	return &userRepository{DB}
}

func (r *userRepository) Save(ctx context.Context,user models.User) (*mongo.InsertOneResult, error) {
	r.DB.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys : bson.D{{Key: "email", Value: 1},{Key:"username", Value:1}},
			Options : options.Index().SetUnique(true),
		},
	)
	
	result,err := r.DB.InsertOne(ctx, user)

	if err != nil {
		return result, err
	}

	return result, nil
}


func (r *userRepository) FindByUsername(ctx context.Context, username string) (models.User,  error){

	var user models.User

	err := r.DB.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil{
		return user, err
	}

	return user, nil
}

func (r *userRepository) IsUserExist(ctx context.Context, username string) (bool, error){
	var user models.User

	err := r.DB.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil{
		return false, err
	}

	return true, nil
}