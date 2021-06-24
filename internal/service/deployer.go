package service

import (
	"asimov-deployer-backend/internal/apierror"
	"asimov-deployer-backend/internal/defines"
	"asimov-deployer-backend/internal/domain"
	"asimov-deployer-backend/internal/repository"
	"net/http"
	"path/filepath"
)

var (
	errAssetNotFound = apierror.New(http.StatusNotFound, "asset not found")
)

type DeployerService interface {
	Deploy(body *domain.DeployBody, githubToken *string) *apierror.ApiError
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

func (s *deployerService) Deploy(body *domain.DeployBody, githubToken *string) *apierror.ApiError {
	// Get release
	release, apiErr := s.ghRepo.GetReleaseByTag(body.Owner, body.Repo, body.Tag, *githubToken)
	if apiErr != nil {
		return apiErr
	}

	// Find the asset
	var asset *domain.Asset
	for i := 0; i<len(release.Assets); i++{
		if release.Assets[i].Name == body.Tag {
			asset = &release.Assets[i]
			break
		}
	}

	if asset == nil {
		return errAssetNotFound
	}

	// Create a temp folder
	tmpDir, apiErr := s.fsRepo.CreateTempDir()
	if apiErr != nil {
		return apiErr
	}

	// Download the binary
	downloadPath := filepath.Join(tmpDir, body.Scope)
	apiErr = s.ghRepo.DownloadAsset(body.Owner, body.Repo, asset.ID, downloadPath, *githubToken)
	if apiErr != nil {
		return apiErr
	}

	// Move the binary and name it as scope
	binPath := filepath.Join(defines.BinariesRoot, body.Repo, body.Scope)
	s.fsRepo.Move(downloadPath, binPath)

	// Delete the temp folder
	apiErr = s.fsRepo.DeleteDir(tmpDir)
	if apiErr != nil {
		return apiErr
	}

	// Run the binary
	apiErr = s.fsRepo.Run(binPath)

	return nil
}
