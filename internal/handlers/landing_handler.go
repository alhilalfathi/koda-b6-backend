package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LandingHandler struct {
	service *service.LandingService
}

func NewLandingHandler(s *service.LandingService) *LandingHandler {
	return &LandingHandler{
		service: s,
	}
}

func (h *LandingHandler) RecommendedProducts(ctx *gin.Context) {
	products, err := h.service.RecommendedProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Show recommended products success",
		Results: products,
	})
}
