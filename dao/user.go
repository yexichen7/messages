package dao

import (
	"message-board/model"
)

func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into users(name,password) values (?,?)", u.UserName, u.Password)
	return
}

func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from users where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}

func ChangePasswordByName(username, password string) (err error) {
	DB.Exec("update users set password=? where name=?", password, username)
	return
}
