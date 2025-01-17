package handler

import (
	"github.com/HadryanSilva/finance-control-api/config"
	"github.com/HadryanSilva/finance-control-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SaveAccountHandler
// @Base path /api/v1
// @Summary Create an account
// @Description Create a new account
// @Tags accounts
// @Accept json
// @Produce json
// @Param request body AccountRequest true "Request Body"
// @Success 201 {object} AccountPostResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /accounts [post]
func SaveAccountHandler(ctx *gin.Context) {
	request := AccountRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	account := model.Account{
		Name:    request.Name,
		Balance: request.Balance,
	}

	if err := db.Create(&account).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to create opening")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-account", account)
}

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostgres()
}
