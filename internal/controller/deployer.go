package controller

import (
	"asimov-deployer-backend/internal/defines"
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

	if err != nil || !body.IsValid() {
		ctx.String(http.StatusBadRequest, "Invalid body.")
		return
	}

	githubToken := ctx.Request.Header.Get(defines.HeaderGithubToken)

	if githubToken == "" {
		ctx.String(http.StatusBadRequest, "Missing token.")
		return
	}

	apiErr := c.svc.Deploy(&body, &githubToken)

	if apiErr != nil {
		ctx.AbortWithStatusJSON(apiErr.Status, apiErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deployed correctly"})
}
