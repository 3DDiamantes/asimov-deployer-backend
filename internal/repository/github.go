package repository

import (
	"asimov-deployer-backend/internal/apierror"
	"asimov-deployer-backend/internal/defines"
	"asimov-deployer-backend/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

type GithubRepository interface {
	DownloadAsset(owner string, repo string, assetID uint64, targetFile string, githubToken string) *apierror.ApiError
	GetReleaseByTag(owner string, repo string, tag string, githubToken string) (*domain.GithubGetReleaseByTagResponse, *apierror.ApiError)
}

type githubRepository struct {
	rc *resty.Client
}

func NewGithubRepository(rc *resty.Client) GithubRepository {
	return &githubRepository{
		rc: rc,
	}
}

func (r *githubRepository) GetReleaseByTag(owner string, repo string, tag string, githubToken string) (*domain.GithubGetReleaseByTagResponse, *apierror.ApiError) {
	resp, err := r.rc.R().
		SetHeader("Accept", defines.GithubHeaderAccept).
		SetHeader("Authorization", "bearer "+githubToken).
		SetPathParam(defines.GithubPathParamOwner, owner).
		SetPathParam(defines.GithubPathParamRepository, repo).
		SetPathParam(defines.GithubPathParamTag, tag).
		Get(defines.GithubURLGetReleaseByTag)

	if err != nil {
		return nil, apierror.New(http.StatusInternalServerError, err.Error())
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, apierror.New(resp.StatusCode(), string(resp.Body()))
	}

	var githubGetReleaseByTagResponse domain.GithubGetReleaseByTagResponse
	err = json.Unmarshal(resp.Body(), &githubGetReleaseByTagResponse)

	if err != nil {
		return nil, apierror.New(http.StatusInternalServerError, "failed to unmarshal GetReleaseByTag")
	}

	return &githubGetReleaseByTagResponse, nil
}

func (r *githubRepository) DownloadAsset(owner string, repo string, assetID uint64, targetFile string, githubToken string) *apierror.ApiError {
	// Descarga la el binario de la nueva versión
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/repos/%s/%s/releases/assets/%d", defines.GithubURLBase, owner, repo, assetID), nil)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}
	req.Header.Set("Accept", "application/octet-stream")
	req.Header.Set("Authorization", "bearer "+githubToken)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	// Crea el archivo
	out, err := os.Create(targetFile)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}
	defer out.Close()

	// Escribe en el archivo
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}
	return nil
}
