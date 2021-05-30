package controller

import (
	"asimov-deployer-backend/internal/domain"
	"asimov-deployer-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeployerController interface {
	Deploy(ctx *gin.Context)
}

type deployerController struct {
	svc service.DeployerService
}

func NewDeployerController(s service.DeployerService) DeployerController {
	return &deployerController{
		svc: s,
	}
}

func (c *deployerController) Deploy(ctx *gin.Context) {
	var body domain.DeployBody
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid body.")
	}

	c.svc.Deploy(body)
	ctx.String(http.StatusOK, "Deployed correctly.")
}
