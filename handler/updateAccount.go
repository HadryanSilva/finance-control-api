package handler

import (
	"github.com/HadryanSilva/finance-control-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateAccountHandler
// @Base path /api/v1
// @Summary Update an account
// @Description Update an account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id query string true "Account Identification"
// @Param request body AccountRequest true "Request Body"
// @Success 200 {object} AccountPutResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /account [put]
func UpdateAccountHandler(ctx *gin.Context) {
	request := AccountRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}

	account := model.Account{}

	if err := db.First(&account, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Account not found")
		return
	}

	if request.Name != "" {
		account.Name = request.Name
	}

	if request.Balance != 0 || request.Balance != account.Balance {
		account.Balance = request.Balance
	}

	if err := db.Save(&account).Error; err != nil {
		logger.Errf("Failed to update account: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Failed to update account")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-account", account)

}
