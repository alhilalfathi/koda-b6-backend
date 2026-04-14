package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"koda-b6-backend/internal/cache"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductService struct {
	repo  *repository.ProductRepository
	cache cache.Cache
}

func NewProductService(rp *repository.ProductRepository, c cache.Cache) *ProductService {
	return &ProductService{
		repo:  rp,
		cache: c,
	}
}

// CREATE
func (s *ProductService) CreateProduct(req *models.CreateProductRequest) error {
	if req.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	product := models.Product{
		ProductName: req.ProductName,
		Desc:        req.Desc,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	err := s.repo.CreateProduct(product)
	if err != nil {
		return err
	}

	// invalidate cache
	s.cache.Delete(context.Background(), "products:all")

	return nil
}

// GET ALL
func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	cacheKey := "products:all"

	var products []models.Product

	//ambil dari cache
	err := s.cache.Get(ctx, cacheKey, &products)
	if err == nil {
		return products, nil
	}

	//fallback ke DB
	products, err = s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	//simpan ke cache
	_ = s.cache.Set(ctx, cacheKey, products, 5*time.Minute)

	return products, nil
}

// GET BY ID
func (s *ProductService) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)

	var product models.Product

	// cache
	err := s.cache.Get(ctx, cacheKey, &product)
	if err == nil {
		return &product, nil
	}

	// DB
	productPtr, err := s.repo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	// set cache
	_ = s.cache.Set(ctx, cacheKey, productPtr, 5*time.Minute)

	return productPtr, nil
}

// UPDATE
func (s *ProductService) Update(id int, req models.UpdateProductRequest) error {
	existing, err := s.repo.GetProductById(id)
	if err != nil {
		return errors.New("product not found")
	}

	if req.ProductName != "" {
		existing.ProductName = req.ProductName
	}
	if req.Desc != "" {
		existing.Desc = req.Desc
	}
	if req.Price > 0 {
		existing.Price = req.Price
	}
	if req.Stock > 0 {
		existing.Stock = req.Stock
	}

	err = s.repo.Update(id, *existing)
	if err != nil {
		return err
	}

	// invalidate cache
	ctx := context.Background()
	s.cache.Delete(ctx,
		"products:all",
		fmt.Sprintf("product:%d", id),
	)

	return nil
}

// DELETE
func (s *ProductService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	// invalidate cache
	ctx := context.Background()
	s.cache.Delete(ctx,
		"products:all",
		fmt.Sprintf("product:%d", id),
	)

	return nil
}

func (s *ProductService) UpdateProductImage(productId int, path string) error {
	return s.repo.UpdateImage(productId, path)
}
