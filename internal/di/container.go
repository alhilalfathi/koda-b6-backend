package di

import (
	"koda-b6-backend/internal/handlers"
	"koda-b6-backend/internal/repository"
	"koda-b6-backend/internal/service"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	db *pgx.Conn

	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler

	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handlers.ProductHandler

	reviewRepo    *repository.ReviewRepository
	reviewService *service.ReviewService
	reviewHandler *handlers.ReviewHandler

	categoryRepo    *repository.CategoryRepository
	categoryService *service.CategoryService
	categoryHandler *handlers.CategoryHandler

	sizeRepo    *repository.SizeRepository
	sizeService *service.SizeService
	sizeHandler *handlers.SizeHandler

	variantRepo    *repository.VariantRepository
	variantService *service.VariantService
	variantHandler *handlers.VariantHandler

	discountRepo    *repository.DiscountRepository
	discountService *service.DiscountService
	discountHandler *handlers.DiscountHandler

	cartRepo    *repository.CartRepository
	cartService *service.CartService
	cartHandler *handlers.CartHandler

	forgotPassRepo    *repository.ForgotPassRepository
	forgotPassService *service.ForgotPassService
	forgotPassHandler *handlers.ForgotPassHandler

	landingRepo    *repository.ProductRepository
	landingService *service.LandingService
	landingHandler *handlers.LandingHandler
}

func NewContainer(db *pgx.Conn) *Container {

	container := Container{
		db: db,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.db)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)

	c.productRepo = repository.NewProductRepository(c.db)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handlers.NewProductHandler(c.productService)

	c.reviewRepo = repository.NewReviewRepository(c.db)
	c.reviewService = service.NewReviewService(c.reviewRepo)
	c.reviewHandler = handlers.NewReviewHandler(c.reviewService)

	c.categoryRepo = repository.NewCategoryRepository(c.db)
	c.categoryService = service.NewCategoryService(c.categoryRepo)
	c.categoryHandler = handlers.NewCategoryHandler(c.categoryService)

	c.sizeRepo = repository.NewSizeRepository(c.db)
	c.sizeService = service.NewSizeService(c.sizeRepo)
	c.sizeHandler = handlers.NewSizeHandler(c.sizeService)

	c.variantRepo = repository.NewVariantRepository(c.db)
	c.variantService = service.NewVariantService(c.variantRepo)
	c.variantHandler = handlers.NewVariantHandler(c.variantService)

	c.discountRepo = repository.NewDiscountRepository(c.db)
	c.discountService = service.NewDiscountService(c.discountRepo)
	c.discountHandler = handlers.NewDiscountHandler(c.discountService)

	c.cartRepo = repository.NewCartRepository(c.db)
	c.cartService = service.NewCartService(c.cartRepo)
	c.cartHandler = handlers.NewCartHandler(c.cartService)

	c.forgotPassRepo = repository.NewForgotPassRepository(c.db)
	c.forgotPassService = service.NewForgotPassService(c.forgotPassRepo, c.userRepo)
	c.forgotPassHandler = handlers.NewForgotPassHandler(c.forgotPassService)

	c.landingRepo = repository.NewProductRepository(c.db)
	c.landingService = service.NewLandingService(c.landingRepo)
	c.landingHandler = handlers.NewLandingHandler(c.landingService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}

func (c *Container) ReviewHandler() *handlers.ReviewHandler {
	return c.reviewHandler
}

func (c *Container) CategoryHandler() *handlers.CategoryHandler {
	return c.categoryHandler
}

func (c *Container) SizeHandler() *handlers.SizeHandler {
	return c.sizeHandler
}

func (c *Container) VariantHandler() *handlers.VariantHandler {
	return c.variantHandler
}

func (c *Container) DiscountHandler() *handlers.DiscountHandler {
	return c.discountHandler
}

func (c *Container) CartHandler() *handlers.CartHandler {
	return c.cartHandler
}

func (c *Container) ForgotPassHandler() *handlers.ForgotPassHandler {
	return c.forgotPassHandler
}

func (c *Container) LandingHandler() *handlers.LandingHandler {
	return c.landingHandler
}
