package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(s *service.CartService) *CartHandler {
	return &CartHandler{
		service: s,
	}
}

func (h *CartHandler) CreateCart(ctx *gin.Context) {
	var req models.CreateCartRequest
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
			Message: "Unauthorized: User ID not found",
		})
		return
	}

	req.UserId = userId.(int)

	if err := h.service.CreateCart(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Cart created",
	})
}

func (h *CartHandler) GetAllCarts(ctx *gin.Context) {
	cart, err := h.service.GetAllCarts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Cart List",
		Results: cart,
	})
}

func (h *CartHandler) GetCartById(ctx *gin.Context) {
	val, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Response{Success: false, Message: "Unauthorized"})
		return
	}
	userId := val.(int)

	carts, err := h.service.GetCartById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Your cart list",
		Results: carts,
	})
}

func (h *CartHandler) GetCartByUser(ctx *gin.Context) {
	val, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized: User ID not found in context",
		})
		return
	}

	userID := val.(int)

	results, err := h.service.GetCartDetails(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch cart details: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cart details retrieved successfully",
		"results": results,
	})
}

func (h *CartHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Cart id invalid",
		})
		return
	}

	var req models.UpdateCartRequest

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
		Message: "Cart updated successfully",
	})
}

func (h *CartHandler) Delete(ctx *gin.Context) {

	val, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	userId := val.(int)

	if err := h.service.Delete(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cart cleared successfully",
	})
}
