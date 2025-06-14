package services

import (
	"myapp/configs"
	"myapp/models"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Users = []User{
	{ID: "1", Name: "Yamm"},
	{ID: "2", Name: "All"},
}

func GetAllUsers() []User {
	return Users
}

func CreateUser(user *models.User) (*models.User, error) {
    if err:= configs.DB.Create(&user).Error; err != nil {
        return nil, err
    }
    return user, nil
}