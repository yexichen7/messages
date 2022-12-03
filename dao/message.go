package dao

import (
	"fmt"
	"log"
	"message-board/model"

	"strconv"
)

func InsertMessage(u model.Message) (err error) {
	_, err = DB.Exec("insert into message(authorID,receiverID,message) values (?,?,?)", u.AuthorID, u.RecID, u.Detail)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func GetOneMessage(receiverID string) (err error) {

	row := DB.QueryRow("select * from message where receiverID=?", receiverID)
	var m model.Message
	err = row.Scan(&m.Mid, &m.AuthorID, &m.RecID, &m.Detail)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(m)
	return
}

func GetMessage(receiverID string) (err error) {

	rows, err := DB.Query("select * from message where receiverID= ?", receiverID)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	var m model.Message
	for rows.Next() {
		err = rows.Scan(&m.Mid, &m.AuthorID, &m.RecID, &m.Detail)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(m)
	}

	return
}

func DeleteMessage(u model.Message) (err error) {
	_, err = DB.Exec("delete from message where authorID=?", u.AuthorID)
	return
}

func SearchUserByReceiverID(receiverID string) (u model.Message, err error) {
	receiverId, _ := strconv.Atoi(receiverID)
	fmt.Println(receiverId)
	row := DB.QueryRow("select id,name,password from users where id=?", receiverId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Mid, u.AuthorID, u.RecID, u.Detail)
	return
}
