package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starter-kit-go/app/models/dto_request/auth"
	"starter-kit-go/app/response"
	"starter-kit-go/helper"
)

type AuthController struct{}

func (controller AuthController) Register(c *gin.Context) {
	var input auth.Registerinput
	//input.Name = "nashir"
	err := helper.ValidateStructFormatted(input)
	fmt.Println("ERRORNYA ", err)
	c.JSON(http.StatusOK, response.Api().Success("Sukses", "joss"))
}
