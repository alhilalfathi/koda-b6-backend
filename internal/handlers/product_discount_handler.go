package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductDiscountHandler struct {
	service *service.ProductDiscountService
}

func NewProductDiscountHandler(s *service.ProductDiscountService) *ProductDiscountHandler {
	return &ProductDiscountHandler{
		service: s,
	}
}

func (h *ProductDiscountHandler) Create(ctx *gin.Context) {
	var req models.CreateProductDiscountRequest
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
		Message: "Product_Discount created successfully",
	})
}

func (h *ProductDiscountHandler) GetAll(ctx *gin.Context) {
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
		Message: "Show all Product_Discount success",
		Results: products,
	})
}

func (h *ProductDiscountHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Discount id invalid",
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
		Message: "Product_Discount found",
		Results: product,
	})
}
