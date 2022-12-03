package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var OK = respTemplate{
	Status: 200,
	Info:   "Success",
}

var LoginOk = respTemplate{
	Status: 200,
	Info:   "Login success",
}

var ChangePasswordOK = respTemplate{
	Status: 200,
	Info:   "Change password success",
}

var ParamError = respTemplate{
	Status: 300,
	Info:   "paramError",
}

var InternalError = respTemplate{
	Status: 500,
	Info:   "internalError",
}

var LeaveMessageSuccess = respTemplate{
	Status: 200,
	Info:   "Leave message success",
}

var GetMessageSuccess = respTemplate{
	Status: 200,
	Info:   "读取成功",
}

var DeleteMessageSuccess = respTemplate{
	Status: 200,
	Info:   "删除成功",
}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
	return
}

func RespLoginSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, LoginOk)
	return
}

func RespChangePasswordSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, ChangePasswordOK)
	return
}

func RespParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
	return
}

func RespInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
	return
}

func RespNormalError(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
	return
}

func RespLeaveMessageSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, LeaveMessageSuccess)
	return
}

func RespGetMessageSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, GetMessageSuccess)
	return
}

func RespDeleteMessageSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, DeleteMessageSuccess)
	return

}
