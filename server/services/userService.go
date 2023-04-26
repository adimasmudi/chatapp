package services

import (
	"chatapp/helper"
	"chatapp/inputs"
	"chatapp/models"
	"chatapp/repositories"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


type UserService interface {
	Register(ctx context.Context, input inputs.RegisterUserInput) (*mongo.InsertOneResult,error)
	Login(ctx context.Context, input inputs.LoginUserInput) (models.User, string, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Register(ctx context.Context, input inputs.RegisterUserInput) (*mongo.InsertOneResult, error){
	
	userExist, _ := s.repository.IsUserExist(ctx,input.Username )

	if userExist{
		return nil, errors.New("User already exist")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil{
		return nil, err
	}


	newUser := models.User{
		Username : input.Username,
		Email : input.Email,
		Password : string(passwordHash),
	}

	registeredUser, err := s.repository.Save(ctx,newUser)

	if err != nil{
		return nil, err
	}

	return registeredUser, nil
}

func (s *userService) Login(ctx context.Context, input inputs.LoginUserInput) (models.User, string, error){

	user, err := s.repository.FindByUsername(ctx,input.Username)

	if err != nil{
		return user, "", errors.New("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil{
		return user, "", errors.New("wrong Password")
	}

	token, err := helper.GenerateToken(user.Email)

	if err != nil{
		return user, "", errors.New("can't generate token")
	}

	return user, token, nil
}