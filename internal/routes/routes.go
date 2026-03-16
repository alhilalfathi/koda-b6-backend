package routes

import (
	"koda-b6-backend/internal/di"
	"koda-b6-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(r *gin.Engine, conn *pgx.Conn) {
	container := di.NewContainer(conn)

	userHandler := container.UserHandler()
	productHandler := container.ProductHandler()
	reviewHandler := container.ReviewHandler()
	categoryHandler := container.CategoryHandler()
	sizeHandler := container.SizeHandler()
	variantHandler := container.VariantHandler()
	discountHandler := container.DiscountHandler()
	carthandler := container.CartHandler()
	proCatHandler := container.ProCatHandler()
	proVarHandler := container.ProVarHandler()
	proSizHandler := container.ProSizHandler()
	proDisHandler := container.ProDisHandler()
	trxHandler := container.TransactionHandler()
	trpHandler := container.TransactionProductHandler()
	forgotPassHandler := container.ForgotPassHandler()
	landingHandler := container.LandingHandler()

	a := r.Group("/admin")
	{
		auth := a.Group("/auth")
		{
			auth.POST("/register", userHandler.Create)
			auth.POST("/login", userHandler.Login)
			auth.POST("/forgot-password", forgotPassHandler.RequestForgotPass)
			auth.PATCH("/forgot-password", forgotPassHandler.ResetPass)
		}

		users := a.Group("/users", middleware.AuthMiddleware())
		{
			users.GET("/", userHandler.GetAll)
			users.GET("/:id", userHandler.GetById)
			users.GET("/email/:email", userHandler.GetByEmail)
			users.PATCH("/:email", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}

		products := a.Group("/products")
		{
			products.POST("/", productHandler.CreateProduct)
			products.GET("/", productHandler.GetAllProduct)
			products.GET("/:id", productHandler.GetProductById)
			products.PATCH("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}

		reviews := a.Group("/reviews")
		{
			reviews.POST("/", reviewHandler.CreateReview)
			reviews.GET("/", reviewHandler.GetAllReview)
			reviews.GET("/:id", reviewHandler.GetReviewById)
			reviews.PATCH("/:id", reviewHandler.Update)
			reviews.DELETE("/:id", reviewHandler.Delete)
		}

		category := a.Group("/category")
		{
			category.POST("/", categoryHandler.CreateCategory)
			category.GET("/", categoryHandler.GetAllCategory)
			category.GET("/:id", categoryHandler.GetCategoryById)
			category.PATCH("/:id", categoryHandler.Update)
			category.DELETE("/:id", categoryHandler.Delete)
		}

		size := a.Group("/size")
		{
			size.POST("/", sizeHandler.CreateSize)
			size.GET("/", sizeHandler.GetAllSizes)
			size.GET("/:id", sizeHandler.GetSizeById)
			size.PATCH("/:id", sizeHandler.Update)
			size.DELETE("/:id", sizeHandler.Delete)
		}

		variant := a.Group("/variant")
		{
			variant.POST("/", variantHandler.CreateVariant)
			variant.GET("/", variantHandler.GetAllVariants)
			variant.GET("/:id", variantHandler.GetVariantById)
			variant.PATCH("/:id", variantHandler.Update)
			variant.DELETE("/:id", variantHandler.Delete)
		}

		discount := a.Group("/discount")
		{
			discount.POST("/", discountHandler.CreateDiscount)
			discount.GET("/", discountHandler.GetAllDiscount)
			discount.GET("/:id", discountHandler.GetDiscountById)
			discount.PATCH("/:id", discountHandler.Update)
			discount.DELETE("/:id", discountHandler.Delete)
		}

		cart := a.Group("/cart")
		{
			cart.POST("/", carthandler.CreateCart)
			cart.GET("/", carthandler.GetAllCarts)
			cart.GET("/:id", carthandler.GetCartById)
			cart.PATCH("/:id", carthandler.Update)
			cart.DELETE("/:id", carthandler.Delete)
		}

		proCat := a.Group("/product_category")
		{
			proCat.POST("/", proCatHandler.Create)
			proCat.GET("/", proCatHandler.GetAll)
			proCat.GET("/:id", proCatHandler.GetById)
		}

		proVar := a.Group("/product_variant")
		{
			proVar.POST("/", proVarHandler.Create)
			proVar.GET("/", proVarHandler.GetAll)
			proVar.GET("/:id", proVarHandler.GetById)
		}

		proSiz := a.Group("/product_size")
		{
			proSiz.POST("/", proSizHandler.Create)
			proSiz.GET("/", proSizHandler.GetAll)
			proSiz.GET("/:id", proSizHandler.GetById)
		}

		proDis := a.Group("/product_discount")
		{
			proDis.POST("/", proDisHandler.Create)
			proDis.GET("/", proDisHandler.GetAll)
			proDis.GET("/:id", proDisHandler.GetById)
		}

		trx := a.Group("/transaction")
		{
			trx.POST("/", trxHandler.CreateTransaction)
			trx.GET("/", trxHandler.GetAllTransaction)
			trx.GET("/:id", trxHandler.GetTransactionById)
		}

		trp := a.Group("/transaction-product")
		{
			trp.POST("/", trpHandler.Create)
			trp.GET("/", trpHandler.GetAll)
			trp.GET("/:id", trpHandler.GetById)
		}
	}

	r.GET("/recommended-products", landingHandler.RecommendedProducts)
	r.GET("/reviews", landingHandler.GetAllReview)
}
