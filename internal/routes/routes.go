package routes

import (
	"koda-b6-backend/internal/cache"
	"koda-b6-backend/internal/di"
	"koda-b6-backend/internal/middleware"
	"koda-b6-backend/internal/models"
	"net/http"

	_ "koda-b6-backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Backend Coffeeshop
// @version 1.0.0
// @description this is basic backend apps with DI
// @host localhost:8888
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetupRoutes(r *gin.Engine, conn *pgxpool.Pool, c cache.Cache) {

	// INIT CONTAINER
	container := di.NewContainer(conn, c)

	userHandler := container.UserHandler()
	productHandler := container.ProductHandler()
	reviewHandler := container.ReviewHandler()
	categoryHandler := container.CategoryHandler()
	sizeHandler := container.SizeHandler()
	variantHandler := container.VariantHandler()
	discountHandler := container.DiscountHandler()
	cartHandler := container.CartHandler()
	proCatHandler := container.ProCatHandler()
	proVarHandler := container.ProVarHandler()
	proSizHandler := container.ProSizHandler()
	proDisHandler := container.ProDisHandler()
	trxHandler := container.TransactionHandler()
	trpHandler := container.TransactionProductHandler()
	forgotPassHandler := container.ForgotPassHandler()
	landingHandler := container.LandingHandler()

	// ADMIN ROUTES
	admin := r.Group("/admin")
	{

		users := admin.Group("/users", middleware.AuthMiddleware())
		{
			users.GET("/", userHandler.GetAll)
			users.GET("/:id", userHandler.GetById)
			users.GET("/email/:email", userHandler.GetByEmail)
			users.DELETE("/:id", userHandler.Delete)

			users.GET("/profile", userHandler.GetProfile)
			users.PATCH("/profile", userHandler.UpdateProfile)
			users.PATCH("/profile/password", userHandler.ChangePassword)
		}

		products := admin.Group("/products")
		{
			products.POST("/", productHandler.CreateProduct)
			products.GET("/", productHandler.GetAllProduct)
			products.GET("/:id", productHandler.GetProductById)
			products.PATCH("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}

		reviews := admin.Group("/reviews")
		{
			reviews.POST("/", reviewHandler.CreateReview)
			reviews.GET("/", reviewHandler.GetAllReview)
			reviews.GET("/:id", reviewHandler.GetReviewById)
			reviews.PATCH("/:id", reviewHandler.Update)
			reviews.DELETE("/:id", reviewHandler.Delete)
		}

		category := admin.Group("/category")
		{
			category.POST("/", categoryHandler.CreateCategory)
			category.GET("/", categoryHandler.GetAllCategory)
			category.GET("/:id", categoryHandler.GetCategoryById)
			category.PATCH("/:id", categoryHandler.Update)
			category.DELETE("/:id", categoryHandler.Delete)
		}

		size := admin.Group("/size")
		{
			size.POST("/", sizeHandler.CreateSize)
			size.GET("/", sizeHandler.GetAllSizes)
			size.GET("/:id", sizeHandler.GetSizeById)
			size.PATCH("/:id", sizeHandler.Update)
			size.DELETE("/:id", sizeHandler.Delete)
		}

		variant := admin.Group("/variant")
		{
			variant.POST("/", variantHandler.CreateVariant)
			variant.GET("/", variantHandler.GetAllVariants)
			variant.GET("/:id", variantHandler.GetVariantById)
			variant.PATCH("/:id", variantHandler.Update)
			variant.DELETE("/:id", variantHandler.Delete)
		}

		discount := admin.Group("/discount")
		{
			discount.POST("/", discountHandler.CreateDiscount)
			discount.GET("/", discountHandler.GetAllDiscount)
			discount.GET("/:id", discountHandler.GetDiscountById)
			discount.PATCH("/:id", discountHandler.Update)
			discount.DELETE("/:id", discountHandler.Delete)
		}

		cart := admin.Group("/cart", middleware.AuthMiddleware())
		{
			cart.POST("/", cartHandler.CreateCart)
			cart.GET("/", cartHandler.GetCartByUser)
			cart.PATCH("/:id", cartHandler.Update)
			cart.DELETE("/user", cartHandler.Delete)
		}

		proCat := admin.Group("/product_category")
		{
			proCat.POST("/", proCatHandler.Create)
			proCat.GET("/", proCatHandler.GetAll)
			proCat.GET("/:id", proCatHandler.GetById)
		}

		proVar := admin.Group("/product_variant")
		{
			proVar.POST("/", proVarHandler.Create)
			proVar.GET("/", proVarHandler.GetAll)
			proVar.GET("/:id", proVarHandler.GetById)
		}

		proSiz := admin.Group("/product_size")
		{
			proSiz.POST("/", proSizHandler.Create)
			proSiz.GET("/", proSizHandler.GetAll)
			proSiz.GET("/:id", proSizHandler.GetById)
		}

		proDis := admin.Group("/product_discount")
		{
			proDis.POST("/", proDisHandler.Create)
			proDis.GET("/", proDisHandler.GetAll)
			proDis.GET("/:id", proDisHandler.GetById)
		}

		trx := admin.Group("/transaction", middleware.AuthMiddleware())
		{
			trx.POST("/", trxHandler.CreateTransaction)
			trx.GET("/", trxHandler.GetAllTransaction)
			trx.GET("/:id", trxHandler.GetTransactionById)
		}

		trp := admin.Group("/transaction-product")
		{
			trp.POST("/", trpHandler.Create)
			trp.GET("/", trpHandler.GetAll)
			trp.GET("/:id", trpHandler.GetById)
		}
	}

	// AUTH
	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Create)
		auth.POST("/login", userHandler.Login)
		auth.POST("/forgot-password", forgotPassHandler.RequestForgotPass)
		auth.PATCH("/forgot-password", forgotPassHandler.ResetPass)
	}

	// PUBLIC
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "Backend running well",
		})
	})

	r.GET("/recommended-products", landingHandler.RecommendedProducts)
	r.GET("/reviews", landingHandler.GetAllReview)

	products := r.Group("/products")
	{
		products.POST("", productHandler.CreateProduct)
		products.GET("", productHandler.GetAllProduct)
		products.GET("/:id", productHandler.GetProductById)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
