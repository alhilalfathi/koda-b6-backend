package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPassHandler struct {
	service *service.ForgotPassService
}

func NewForgotPassHandler(sr *service.ForgotPassService) *ForgotPassHandler {
	return &ForgotPassHandler{
		service: sr,
	}
}

func (h *ForgotPassHandler) RequestForgotPass(ctx *gin.Context) {
	var req models.CreateForgotPassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.RequestForgotPass(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success sent request",
	})
}

func (h *ForgotPassHandler) ResetPass(ctx *gin.Context) {
	var req models.UpdateForgotPassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.ResetPass(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Password change success",
	})
}
