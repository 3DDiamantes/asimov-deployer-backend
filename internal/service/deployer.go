package service

import (
	"asimov-deployer-backend/internal/domain"
	"asimov-deployer-backend/internal/repository"
)

type DeployerService interface {
	Deploy(body domain.DeployBody) error
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

func (s *deployerService) Deploy(body domain.DeployBody) error {
	// Download binary from Github
	release, err := s.ghRepo.GetReleaseByTag(body.Owner, body.Repo, body.Tag)
	if err != nil {
		return err
	}

	for _, asset := range release.Assets {
		if asset.Name == body.Tag {
			s.ghRepo.DownloadAsset(body.Owner, body.Repo, body.Tag, asset.Name)
			break
		}
	}

	// Move the asset to the scope folder
	s.fsRepo.Move(body.Scope)

	// Run the binary
	s.fsRepo.Run(body.Scope, body.Tag)

	return nil
}
