package routes

import (
	"github.com/gin-gonic/gin"
	"starter-kit-go/app/controllers"
	"starter-kit-go/app/controllers/auth"
)

type Api struct{}

func (a Api) Do(route *gin.Engine) {
	api := route.Group("/api")
	// example test api
	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world !!!",
		})
	})
	// example test api on controller
	api.GET("example", controllers.ExampleController{}.Index)

	/*GROUPING AUTH*/
	authRoute := api.Group("/auth")
	authRoute.POST("register", auth.AuthController{}.Register)

}
