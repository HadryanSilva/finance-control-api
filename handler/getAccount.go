package handler

import (
	"fmt"
	"github.com/HadryanSilva/finance-control-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAccountHandler
// @Base path /api/v1
// @Summary Find an account
// @Description Find an account data
// @Tags accounts
// @Accept json
// @Produce json
// @Param id query string true "Account Identification"
// @Success 200 {object} AccountGetResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /account [get]
func GetAccountHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	account := model.Account{}
	if err := db.First(&account, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Account with id %s not found", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "get-opening", account)
}
