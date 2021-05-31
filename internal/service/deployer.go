package service

import (
	"asimov-deployer-backend/internal/domain"
	"asimov-deployer-backend/internal/repository"
	"os"
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

	tmpDir, err := os.MkdirTemp("", "asimov-deployer-*")
	if err != nil {
		return err
	}

	for _, asset := range release.Assets {
		if asset.Name == body.Tag {
			err = s.ghRepo.DownloadAsset(body.Owner, body.Repo, asset.ID, asset.Name, tmpDir)
			if err != nil {
				return err
			}
			break
		}
	}

	// Move the asset to the scope folder
	s.fsRepo.Move(tmpDir, body.Tag, body.Scope)

	// Delete the temp folder
	err = os.RemoveAll(tmpDir)
	if err != nil {
		return err
	}

	// Run the binary
	s.fsRepo.Run(body.Scope, body.Tag)

	return nil
}
