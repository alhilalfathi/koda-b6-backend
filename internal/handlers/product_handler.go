package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

// CreateProduct godoc
// @Summary      Create product
// @Description  Create a Product
// @Tags         Product
// @Produce      json
// @Success      200      {object}  models.CreateProductRequest
// @Failure      500      {object}  models.CreateProductRequest
// @Router       /admin/products [post]
func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req models.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.CreateProduct(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Product created successfully",
	})
}

// GetProducts godoc
// @Summary      Get products
// @Description  Show All Products
// @Tags         Products
// @Produce      json
// @Success      200      {object}  models.Product
// @Failure      500      {object}  models.Product
// @Router       /admin/products [get]
func (h *ProductHandler) GetAllProduct(ctx *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Show products success",
		Results: products,
	})
}

func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Product id invalid",
		})
		return
	}

	product, err := h.service.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Product found",
		Results: product,
	})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Product id invalid",
		})
		return
	}

	var req models.UpdateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.Update(id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Product updated successfully",
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Product id invalid",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Product deleted",
	})
}
