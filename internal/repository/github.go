package repository

import (
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
	DownloadAsset(owner string, repo string, assetID uint64, filename string, path string) error
	GetReleaseByTag(owner string, repo string, tag string) (*domain.GithubGetReleaseByTagResponse, error)
}

type githubRepository struct {
	rc *resty.Client
}

func NewGithubRepository(rc *resty.Client) GithubRepository {
	return &githubRepository{
		rc: rc,
	}
}

func (r *githubRepository) GetReleaseByTag(owner string, repo string, tag string) (*domain.GithubGetReleaseByTagResponse, error) {
	resp, err := r.rc.R().
		SetHeader("Authorization", "bearer "+defines.GithubToken).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetPathParam(defines.GithubPathParamOwner, owner).
		SetPathParam(defines.GithubPathParamRepository, repo).
		SetPathParam(defines.GithubPathParamTag, tag).
		Get(defines.GithubURLGetReleaseByTag)

	if err != nil {
		return nil, err
	}

	var githubGetReleaseByTagResponse domain.GithubGetReleaseByTagResponse
	err = json.Unmarshal(resp.Body(), &githubGetReleaseByTagResponse)

	if err != nil {
		return nil, err
	}

	return &githubGetReleaseByTagResponse, nil
}

func (r *githubRepository) DownloadAsset(owner string, repo string, assetID uint64, filename string, path string) error {
	// Descarga la el binario de la nueva versión
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/repos/%s/%s/releases/assets/%d", defines.GithubURLBase, owner, repo, assetID), nil)
	req.Header.Set("Accept", "application/octet-stream")
	req.Header.Set("Authorization", "bearer "+defines.GithubToken)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Crea el archivo
	out, err := os.Create(path + "/" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Escribe en el archivo
	_, err = io.Copy(out, resp.Body)

	return err
}
