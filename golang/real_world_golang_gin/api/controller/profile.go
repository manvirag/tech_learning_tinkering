package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manvirag982/tech_learning_tinkering/golang/real_world_golang_gin/domain"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	c.JSON(http.StatusOK, profile)
}
