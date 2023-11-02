package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/api/route"
)

func main() {

	gin := gin.Default()

	route.Setup(gin)

	gin.Run("4000")
}
