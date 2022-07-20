package server

import (
	"github.com/RolandKun5/go-rest-api/src/routers"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	routers.Init(router)
	router.Run("localhost:9000")
}
