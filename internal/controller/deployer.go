package controller

import (
	"asimov-deployer-backend/internal/apierror"
	"asimov-deployer-backend/internal/defines"
	"asimov-deployer-backend/internal/domain"
	"asimov-deployer-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	errInvalidBody = apierror.New(http.StatusBadRequest, "invalid body")
	errMissingToken = apierror.New(http.StatusBadRequest, "missing token")
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
		ctx.AbortWithError(errInvalidBody.Status, errInvalidBody)//nolint
		return
	}

	githubToken := ctx.Request.Header.Get(defines.HeaderGithubToken)

	if githubToken == "" {
		ctx.AbortWithError(errMissingToken.Status, errMissingToken)//nolint
		return
	}

	apiErr := c.svc.Deploy(&body, &githubToken)

	if apiErr != nil {
		ctx.AbortWithError(apiErr.Status, apiErr)//nolint
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deployed correctly"})
}
