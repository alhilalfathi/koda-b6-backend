package handlers

import (
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

// CREATE

// CreateProduct godoc
// @Summary Create product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body models.CreateProductRequest true "Create Product Request"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /admin/products [post]
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

// GET ALL

// GetAllProduct godoc
// @Summary Get all products
// @Description Retrieve all products
// @Tags Products
// @Produce json
// @Success 200 {object} models.Response{results=[]models.Product}
// @Failure 500 {object} models.Response
// @Router /admin/products [get]
func (h *ProductHandler) GetAllProduct(ctx *gin.Context) {
	products, err := h.service.GetAllProducts(ctx.Request.Context())
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

// GET BY ID

// GetProductById godoc
// @Summary Get product by ID
// @Description Retrieve product detail by ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Response{results=models.Product}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /admin/products/{id} [get]
func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Product id invalid",
		})
		return
	}

	product, err := h.service.GetProductById(ctx.Request.Context(), id)
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

// UPDATE

// UpdateProduct godoc
// @Summary Update product
// @Description Update product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body models.UpdateProductRequest true "Update Product Request"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /admin/products/{id} [patch]
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

// DELETE

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /admin/products/{id} [delete]
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

func (h *ProductHandler) UploadProductImage(ctx *gin.Context) {
	idParam := ctx.Param("id")
	productId, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid product id",
		})
		return
	}

	// ambil file
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Image is required",
		})
		return
	}

	// validasi tipe file
	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "File must be image",
		})
		return
	}

	// buat folder kalau belum ada
	uploadDir := "uploads/products"
	os.MkdirAll(uploadDir, os.ModePerm)

	// buat nama file unik
	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("product_%d_%d%s", productId, time.Now().Unix(), ext)

	filePath := filepath.Join(uploadDir, fileName)

	// simpan file
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to save image",
		})
		return
	}

	// simpan path ke DB
	err = h.service.UpdateProductImage(productId, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Image uploaded",
		Results: map[string]string{
			"path": filePath,
		},
	})
}
