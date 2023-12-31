package p_routes

import (
	"github.com/gin-gonic/gin"
	"starter-kit-go/config"
	"starter-kit-go/routes"
)

var routeRegistry = []IRoute{
	routes.Api{},
	routes.Web{},
}

// Run routers registered
func Run() {
	router := gin.Default()
	// bind route registered
	for _, r := range routeRegistry {
		r.Do(router)
	}

	// run router with specific port app configuration
	err := router.Run(":" + config.App().Port)
	if err != nil {
		panic("error run router : " + err.Error())
	}
}
