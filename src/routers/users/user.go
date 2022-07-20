package users

import (
	"github.com/RolandKun5/go-rest-api/src/routers/users/handlers"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("users/:userid", handlers.GetUserById)
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateNewUser)
	router.PATCH("/users", handlers.UpdateUser)
}
