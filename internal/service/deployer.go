package service

import (
	"asimov-deployer-backend/internal/domain"
	"asimov-deployer-backend/internal/repository"
)

type DeployerService interface {
	Deploy(body domain.DeployBody)
}

type deployerService struct {
	ghRepo repository.GithubRepository
	fsRepo repository.FilesystemRepository
}

func NewDeployerService(ghr repository.GithubRepository, fsr repository.FilesystemRepository) DeployerService {
	return &deployerService{
		ghRepo: ghr,
		fsRepo: fsr,
	}
}

func (s *deployerService) Deploy(body domain.DeployBody) {
	// Download binary from Github
	s.ghRepo.DownloadAsset(body.Repo, body.Tag)

	// Move the asset to the scope folder
	s.fsRepo.Move(body.Scope)

	// Run the binary
	s.fsRepo.Run(body.Scope, body.Tag)
}
