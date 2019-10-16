package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/shadowlink0122/sakura_test/pkg/di"
)

type UserCreateResponse struct {
	Token string `json:token`
}

func UserCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get form data
		ctx.Request.ParseForm()
		username := ctx.Request.Form["id"][0]
		// password := ctx.Request.Form["password"][0]

		// check form-data format
		// if err != nil {
		// ctx.String(
		// 	// http.Status
		// )
		// }

		// Create TokenID
		// token, err := di.User.Create(username, password)
		// if err != nil {
		// ctx.String(
		// 	http.StatusBadRequest,
		// 	err
		// )
		// }

		// ok
		ctx.JSON(
			http.StatusOK,
			UserCreateResponse{
				Token: username,
			},
		)
	}
}

func UserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func UserData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
