package controllers

import (
	"autoposter/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// SignIn - login
func (auth *AuthController) SignIn(ctx *gin.Context) {

	_, err := ctx.Writer.Write([]byte("SignIn"))
	if err != nil {
		fmt.Println(err)
	}
}

// SignUp - register
func (auth *AuthController) SignUp(ctx *gin.Context) {
	var input models.User

	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println(err)
	}

	fmt.Println(input)

	//_, err := ctx.Writer.Write()
	//if err != nil {
	//	fmt.Println(err)
	//}
}
