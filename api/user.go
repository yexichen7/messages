package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func Register(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.RespParamError(c)
		return
	}
	u, err := service.SearchUserByUserName(userName)
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalError(c)
		return
	}

	if u.UserName != "" {
		util.RespNormalError(c, 300, "账号已处在")
		return
	}

	err = service.CreateUser(model.User{
		UserName: userName,
		Password: password,
	})
	if err != nil {
		util.RespInternalError(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.RespParamError(c)
		return
	}

	u, err := service.SearchUserByUserName(userName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespNormalError(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalError(c)
			return
		}

		return
	}
	if u.Password != password {
		util.RespNormalError(c, 300, "密码错误")
		return
	}
	util.RespLoginSuccess(c)
	c.SetCookie("name", password, 0, "", "/", false, false)
}

func ChangePassword(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.RespParamError(c)
		return
	}

	_, err := service.SearchUserByUserName(userName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespNormalError(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalError(c)
			return
		}
		return
	}
	err = service.ChangePasswordByName(userName, password)
	if err != nil {
		util.RespNormalError(c, 300, "密码修改失败")
		return
	}
	util.RespChangePasswordSuccess(c)
}
