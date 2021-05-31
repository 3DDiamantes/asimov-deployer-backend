package http

import (
	"asimov-deployer-backend/internal/controller"
	"asimov-deployer-backend/internal/repository"
	"asimov-deployer-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	MapRoutes(router)

	return router
}

func MapRoutes(r *gin.Engine) {
	restClient := resty.New()

	// Repository
	githubRepo := repository.NewGithubRepository(restClient)
	filesystemRepo := repository.NewFilesystemRepository()

	// Service
	deployService := service.NewDeployerService(githubRepo, filesystemRepo)

	// Controller
	deployController := controller.NewDeployerController(deployService)

	// Endpoint
	r.POST("/deploy", deployController.Deploy)
	r.GET("/ping", ping)
}

func ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
