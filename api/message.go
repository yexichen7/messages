package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func LeaveMessage(c *gin.Context) {
	authorID := c.PostForm("authorID")
	receiverID := c.PostForm("receiverID")
	detail := c.PostForm("detail")
	if authorID == "" || receiverID == "" || detail == "" {
		util.RespParamError(c)
		return
	}
	err := service.CreateMessage(model.Message{
		AuthorID: authorID,
		RecID:    receiverID,
		Detail:   detail,
	})
	if err != nil {
		util.RespNormalError(c, 300, "发送留言失败")
		return
	}
	util.RespLeaveMessageSuccess(c)
}
func GetOneMessage(c *gin.Context) {
	authorID := c.Query("authorID")
	receiverID := c.Query("receiverID")
	if authorID == "" || receiverID == "" {
		util.RespParamError(c)
		return
	}
	err := service.GetOneMessage(receiverID)
	if err != nil {
		util.RespNormalError(c, 300, "接收留言失败")
		return
	}
	util.RespGetMessageSuccess(c)
}

func GetAllMessage(c *gin.Context) {
	authorID := c.Query("authorID")
	receiverID := c.Query("receiverID")
	if authorID == "" || receiverID == "" {
		util.RespParamError(c)
		return
	}
	err := service.GetMessage(receiverID)
	if err != nil {
		util.RespNormalError(c, 300, "接收留言失败")
		return
	}
	util.RespGetMessageSuccess(c)
}

func DeleteMessage(c *gin.Context) {
	authorID := c.PostForm("authorID")
	receiverID := c.PostForm("receiverID")
	detail := c.PostForm("detail")
	if authorID == "" || receiverID == "" || detail == "" {
		util.RespParamError(c)
		return
	}
	u, err := service.SearchUserByReceiverID(receiverID)
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
	err = service.DeleteMessage(u)
	if err != nil {
		util.RespInternalError(c)
		return
	}
	util.RespDeleteMessageSuccess(c)
}
