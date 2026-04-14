package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: s,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var req models.CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized: User session not found",
		})
		return
	}
	req.UserId = userId.(int)

	if err := h.service.CreateTransaction(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Transaction created",
		Results: req,
	})
}

func (h *TransactionHandler) GetAllTransaction(ctx *gin.Context) {
	tr, err := h.service.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Transaction list",
		Results: tr,
	})
}

func (h *TransactionHandler) GetDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	data, err := h.service.GetDetail(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Results: data,
	})
}

func (h *TransactionHandler) GetByUser(ctx *gin.Context) {

	userIdRaw, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	userId := userIdRaw.(int)

	data, err := h.service.GetByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Results: data,
	})
}
