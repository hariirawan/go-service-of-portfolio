package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transService transaction.Service
}

func NewTransactionHandler(transService transaction.Service) *transactionHandler {
	return &transactionHandler{transService}
}

func (h *transactionHandler) GetTransactionByCampaignID(ctx *gin.Context) {
	var param transaction.ParamTransaction

	err := ctx.ShouldBindUri(&param)
	if err != nil {
		response := helper.APIResponse("Error to get detail of campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	param.User = currentUser

	trans, err := h.transService.GetTransactionByCampaignID(param)

	if err != nil {
		response := helper.APIResponse("Error to get transaction", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatCampaignTransactions(trans)
	response := helper.APIResponse("List of transactions", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)
}
