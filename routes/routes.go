package routes

import (
	"ecommerce/controllers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

// InitRoutes initializes routes
func initRoutes(e *echo.Echo) {
	authController := &controllers.AuthController{}
	initAuthRoutes(e, authController)

	userController := &controllers.UserController{}
	initUserRoutes(e, userController)

	productController := &controllers.ProductController{}
	initProductRoutes(e, productController)

	categoryController := &controllers.CategoryController{}
	initCategoryRoutes(e, categoryController)

	orderController := &controllers.OrderController{}
	initOrderRoutes(e, orderController)

	orderDetailController := &controllers.OrderDetailController{}
	initOrderDetailRoutes(e, orderDetailController)

	reviewController := &controllers.ReviewController{}
	initReviewRoutes(e, reviewController)

	couponController := &controllers.CouponController{}
	initCouponRoutes(e, couponController)
}

func initAuthRoutes(e *echo.Echo, uc *controllers.AuthController) {
	log.Info("initAuthenticationRoutes")
	e.POST("/auth/login", uc.Login)
	e.POST("/auth/register", uc.Register)
}

func initUserRoutes(e *echo.Echo, uc *controllers.UserController) {
	log.Info("initUserRoutes")
	e.GET("/users", uc.GetUsers)
	e.GET("/users/:id", uc.GetUserByID)
	e.POST("/users", uc.CreateUser)
	e.PUT("/users/:id", uc.UpdateUser)
	e.DELETE("/users/:id", uc.DeleteUser)
}

func initProductRoutes(e *echo.Echo, pc *controllers.ProductController) {
	e.GET("/products", pc.GetProducts)
	e.GET("/products/:id", pc.GetProductByID)
	e.POST("/products", pc.CreateProduct)
	e.PUT("/products/:id", pc.UpdateProduct)
	e.DELETE("/products/:id", pc.DeleteProduct)
}

func initCategoryRoutes(e *echo.Echo, cc *controllers.CategoryController) {
	e.GET("/categories", cc.GetAllCategories)
	e.GET("/categories/:id", cc.GetCategoryByID)
	e.POST("/categories", cc.CreateCategory)
	e.PUT("/categories/:id", cc.UpdateCategory)
	e.DELETE("/categories/:id", cc.DeleteCategory)
}

func initOrderRoutes(e *echo.Echo, oc *controllers.OrderController) {
	e.GET("/orders", oc.GetAllOrders)
	e.GET("/orders/:id", oc.GetOrderByID)
	e.POST("/orders", oc.CreateOrder)
	e.PUT("/orders/:id", oc.UpdateOrder)
	e.DELETE("/orders/:id", oc.DeleteOrder)
}

func initOrderDetailRoutes(e *echo.Echo, odc *controllers.OrderDetailController) {
	e.GET("/order_details", odc.GetAllOrderDetails)
	e.GET("/order_details/:id", odc.GetOrderDetailByID)
	e.POST("/order_details", odc.CreateOrderDetail)
	e.PUT("/order_details/:id", odc.UpdateOrderDetail)
	e.DELETE("/order_details/:id", odc.DeleteOrderDetail)
}

func initReviewRoutes(e *echo.Echo, rc *controllers.ReviewController) {
	e.GET("/reviews", rc.GetAllReviews)
	e.GET("/reviews/:id", rc.GetReviewByID)
	e.POST("/reviews", rc.CreateReview)
	e.PUT("/reviews/:id", rc.UpdateReview)
	e.DELETE("/reviews/:id", rc.DeleteReview)
}

func initCouponRoutes(e *echo.Echo, cc *controllers.CouponController) {
	e.GET("/coupons", cc.GetAllCoupons)
	e.GET("/coupons/:id", cc.GetCouponByID)
	e.POST("/coupons", cc.CreateCoupon)
	e.PUT("/coupons/:id", cc.UpdateCoupon)
	e.DELETE("/coupons/:id", cc.DeleteCoupon)
}

// StartServer starts the Echo server
func StartServer(db *gorm.DB) {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Set debug mode based on the environment variable
	debugMode := os.Getenv("DEBUG")
	if debugMode == "true" {
		e.Debug = true
	}

	initRoutes(e)

	e.GET("/ping", func(c echo.Context) error {
		c.JSON(http.StatusOK, "pong")
		return nil
	})

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
