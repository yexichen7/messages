package service

import (
	"message-board/dao"
	"message-board/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}

func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

func ChangePasswordByName(username, password string) (err error) {
	err = dao.ChangePasswordByName(username, password)
	return err
}
