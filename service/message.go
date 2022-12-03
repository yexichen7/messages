package service

import (
	"log"
	"message-board/dao"
	"message-board/model"
)

func SearchUserByReceiverID(receiverID string) (u model.Message, err error) {
	u, err = dao.SearchUserByReceiverID(receiverID)
	return u, err
}

func CreateMessage(u model.Message) (err error) {
	err = dao.InsertMessage(u)
	return err
}

func GetOneMessage(receiverID string) (err error) {
	err = dao.GetOneMessage(receiverID)
	if err != nil {
		log.Println(err)
		return
	}
	return err
}

func GetMessage(receiverID string) (err error) {
	err = dao.GetMessage(receiverID)
	return err
}

func DeleteMessage(u model.Message) (err error) {
	dao.DeleteMessage(u)
	return err
}
