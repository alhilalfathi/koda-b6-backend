package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductSizeHandler struct {
	service *service.ProductSizeService
}

func NewProductSizeHandler(s *service.ProductSizeService) *ProductSizeHandler {
	return &ProductSizeHandler{
		service: s,
	}
}

func (h *ProductSizeHandler) Create(ctx *gin.Context) {
	var req models.CreateProductSizeRequest
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
		Message: "Product_Size created successfully",
	})
}

func (h *ProductSizeHandler) GetAll(ctx *gin.Context) {
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
		Message: "Show all Product_Size success",
		Results: products,
	})
}

func (h *ProductSizeHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Size id invalid",
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
		Message: "Product_Size found",
		Results: product,
	})
}
