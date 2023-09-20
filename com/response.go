package com

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseData struct {
	Code    ResCode     `json:"code"`
	Message interface{} `json:"message"`
	Result  interface{} `json:"result"`
}

/*
ResponseData 自定义带有数据的响应
*/
func ResponseData(context *gin.Context, code ResCode, result any, message ...string) {
	rd := &responseData{
		Code:    code,
		Message: GetResponseMessage(code, ""),
		Result:  result,
	}
	context.JSON(http.StatusOK, rd)
	context.Abort()
}
