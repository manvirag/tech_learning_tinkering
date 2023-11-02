package main

import (
	"time"

	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/init"

	"github.com/gin-gonic/gin"
)

func main() {

	env := init.NewEnv()
	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, gin)

	gin.Run(env.ServerAddress)
}
