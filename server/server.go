package main

import (
	"net/http"
	"server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// routhing
	router := gin.Default()

	// top
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(
			http.StatusOK,
			"TopPage",
		)
	})

	// authorication
	router.POST("/user/create", handler.UserCreate())
	// router.POST("/user/login", handler.AuthLogin())

	// user data
	// router.GET("/user", handler.UserData())
	// router.GET("/user/:user_id", handler.UserData())

	// run
	router.Run(":8090")
}
