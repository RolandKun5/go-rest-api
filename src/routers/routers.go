package routers

import (
	"github.com/RolandKun5/go-rest-api/src/routers/users"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	users.Init(router)
}
