package routes

import (
	appconfig "gin-gonic-gorm/configs/app_config"
	bookcontroller "gin-gonic-gorm/controllers/book_controller"
	filecontroller "gin-gonic-gorm/controllers/file_controller"
	usercontroller "gin-gonic-gorm/controllers/user_controller"
	"gin-gonic-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine)  {
	route := app

	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)

	// Middleware untuk mengizinkan CORS dari semua domain
	route.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	//API Versioning
	v1 := route.Group("/api/v1")

	// Route User
	v1.GET("/user", usercontroller.GetAllUser)
	v1.GET("/user/paginate", usercontroller.GetUserPaginate)
	v1.GET("/user/:id", usercontroller.GetById)
	v1.POST("/user", usercontroller.Store)
	v1.PATCH("/user/:id", usercontroller.UpdateById)
	v1.DELETE("/user/:id", usercontroller.DeleteById)
	v1.POST("/user/register", usercontroller.RegisterUser)


	// Route Book
	v1.GET("/book", bookcontroller.GetAllBook)

	// Route file
	authRoute := v1.Group("file", middleware.AuthMiddleware)
	authRoute.POST("/", filecontroller.HandleUploadFile)
	authRoute.DELETE("/:filename", filecontroller.HandleRemoveFile)

}
