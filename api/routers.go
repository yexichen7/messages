package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.PUT("/password", ChangePassword)
	}
	m := r.Group("/message")
	{
		m.POST("/leaveMessage", LeaveMessage)
		m.GET("/getOneMessage", GetOneMessage)
		m.GET("/getAllMessage", GetAllMessage)
		m.DELETE("/deleteMessage", DeleteMessage)
	}
	r.Run()
}
