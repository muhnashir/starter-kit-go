package controllers

import (
	"github.com/gin-gonic/gin"
	"starter-kit-go/app/response"
)

type ExampleController struct{}

func (controller ExampleController) Index(c *gin.Context) {
	c.JSON(200, response.Api().Success("sukses", "coba data"))
}
