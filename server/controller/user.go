package controller

import (
	"fmt"
	"server/entity"
	"server/service"
)

func GetAllUser() []*entity.User {
	userList, err := service.GetAllUser()
	if err != nil {
		fmt.Println((err))
	}
	return userList
}

func CreateUser(user entity.User) error {
	err := service.CreateUser(&user)
	return err
}

func CheckAuth(username string, password string) bool {
	isValidUser := service.FindUser(username, password)
	return isValidUser
}
