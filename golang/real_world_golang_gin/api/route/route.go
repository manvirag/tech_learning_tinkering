package route

import (
	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Engine) {
	// protectedRouter.Use(middleware.auth(env.AccessTokenSecret))
	publicRouter := gin.Group("")
	NewProfileRouter(publicRouter)
}
