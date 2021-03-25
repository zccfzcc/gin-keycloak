package app

import (
	"github.com/gin-gonic/gin"

	"gin-keycloak/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	successFlage := true
	if errCode != e.SUCCESS {
		successFlage = false
	}
	g.C.JSON(httpCode, Response{
		Success: successFlage,
		Code:    errCode,
		Msg:     e.GetMsg(errCode),
		Data:    data,
	})
	return
}
