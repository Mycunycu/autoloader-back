package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthController struct {}

// SignIn - login
func (auth *AuthController) SignIn(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	_, err := ctx.Writer.Write([]byte ("SignIn"))
	if err != nil {
		fmt.Println(err)
	}
}

// SignUp - register
func (auth *AuthController) SignUp(ctx *gin.Context) {

}
