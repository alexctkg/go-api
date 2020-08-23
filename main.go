package main

import (
	"os"
	controllers "tdez/controllers"
	"tdez/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))

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
	router.POST("/superuser", controllers.SuperUserStore)
	router.POST("/externalapp", controllers.ExternalUserStore)

	// router.POST("/login")

	superUserRoutes := router.Group("")
	superUserRoutes.Use(middleware.Jwt(0)){

	}

	externalAppRoutes := router.Group("")
	externalAppRoutes.Use(middleware.Jwt(1)){
		
	}

	return router

}
