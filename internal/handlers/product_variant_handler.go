package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductVariantHandler struct {
	service *service.ProductVariantService
}

func NewProductVariantHandler(s *service.ProductVariantService) *ProductVariantHandler {
	return &ProductVariantHandler{
		service: s,
	}
}

func (h *ProductVariantHandler) Create(ctx *gin.Context) {
	var req models.CreateProductVariantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.Create(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Product_Variant created successfully",
	})
}

func (h *ProductVariantHandler) GetAll(ctx *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Show all Product_Variant success",
		Results: products,
	})
}

func (h *ProductVariantHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Variant id invalid",
		})
		return
	}

	product, err := h.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Product_Variant found",
		Results: product,
	})
}
