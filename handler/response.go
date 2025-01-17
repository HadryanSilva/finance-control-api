package handler

import (
	"fmt"
	"github.com/HadryanSilva/finance-control-api/model"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"error": message,
		"code":  code,
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s successfully", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type AccountPostResponse struct {
	Message string                `json:"message"`
	Data    model.AccountResponse `json:"data"`
}

type AccountGetResponse struct {
	Message string                `json:"message"`
	Data    model.AccountResponse `json:"data"`
}

type AccountDeleteResponse struct {
	Message string                `json:"message"`
	Data    model.AccountResponse `json:"data"`
}

type AccountPutResponse struct {
	Message string                `json:"message"`
	Data    model.AccountResponse `json:"data"`
}
