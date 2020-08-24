package main

import (
	"os"

	entity "tdez/controllers/entity"
	product "tdez/controllers/product"

	"tdez/docs"
	"tdez/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Swagger configuration
	docs.SwaggerInfo.Title = "T10 Test"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"https"}

	router := SetupRouter()
	router.Run()
}

//SetupRouter ..
func SetupRouter() *gin.Engine {
	gotenv.Load()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:    []string{"Content-type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length", "Content-type"},
		MaxAge:          36000,
	}))

	//routes do insert super users and enterprises (no authentication)
	router.POST("/superuser", entity.SuperUserStore)
	router.POST("/externalapp", entity.ExternalUserStore)

	router.POST("/login", entity.Login)

	superUserRouter := router.Group("admin")
	superUserRouter.Use(middleware.Jwt(0)) // 0- to super users
	superUserRouter.PUT("/reject", product.RejectActivation)
	superUserRouter.PUT("/aprove", product.AproveActivation)
	superUserRouter.GET("/index", product.IndexAll)

	externalAppRouter := router.Group("external") //1- to external
	externalAppRouter.Use(middleware.Jwt(1))
	externalAppRouter.POST("/product", product.IssueActivation)
	externalAppRouter.GET("/index", product.IndexExternal)

	swagger := router.Group("/docs")
	swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // api documentation HOST/docs/index.html

	return router

}
