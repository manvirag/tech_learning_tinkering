package route

import (
	"github.com/gin-gonic/gin"
	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/api/controller"
	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/repository"
	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/usecases"
)

func NewProfileRouter(group *gin.RouterGroup) {
	ur := repository.NewUserRepository("abcd", "xyzs")
	pc := &controller.ProfileController{
		ProfileUsecase: usecases.NewProfileUsecase(ur),
	}
	group.GET("/profile", pc.Fetch)
}
