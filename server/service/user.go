package service

import (
	"server/dao"
	"server/entity"
)

func GetAllUser() (userList []*entity.User, err error) {
	if err := dao.SqlConnection.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

/*
*
新建User信息
*/
func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlConnection.Create(user).Error; err != nil {
		return err
	}
	return
}

func FindUser(username, password string) bool {
	var auth entity.User
	dao.SqlConnection.Select("id").Where(entity.User{Name: username, Password: password}).First(&auth)
	if auth.Id > 0 {
		return true
	}
	return false
}
