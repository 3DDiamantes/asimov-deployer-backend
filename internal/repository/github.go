package repository

import (
	"asimov-deployer-backend/internal/apierror"
	"asimov-deployer-backend/internal/defines"
	"asimov-deployer-backend/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	errUnmarshalGetReleaseByTag = apierror.New(http.StatusInternalServerError, "failed to unmarshal GetReleaseByTag")
	errUnmarshalGetAssetByID = apierror.New(http.StatusInternalServerError, "failed to unmarshal GetAssetByID")
)

type GithubRepository interface {
	DownloadAsset(owner string, repo string, assetID uint64, targetFile string, githubToken string) *apierror.ApiError
	GetReleaseByTag(owner string, repo string, tag string, githubToken string) (*domain.GithubGetReleaseByTagResponse, *apierror.ApiError)
	GetAssetByID(owner string, repo string, assetID uint64, githubToken string) (*domain.Asset, *apierror.ApiError)
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
		SetHeader("Accept", defines.GithubHeaderAcceptJSON).
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
		return nil, errUnmarshalGetReleaseByTag
	}

	return &githubGetReleaseByTagResponse, nil
}

func (r *githubRepository) DownloadAsset(owner string, repo string, assetID uint64, targetFile string, githubToken string) *apierror.ApiError {
	asset, apiErr := r.GetAssetByID(owner, repo, assetID, githubToken)
	if apiErr != nil {
		return apiErr
	}

	// Descarga la el binario de la nueva versión
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/repos/%s/%s/releases/assets/%d", defines.GithubURLBase, owner, repo, assetID), nil)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}
	req.Header.Set("Accept", defines.GithubHeaderAcceptOctetStream)
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
	counter := &writeCounter{Total: asset.Size}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (r *githubRepository) GetAssetByID(owner string, repo string, assetID uint64, githubToken string) (*domain.Asset, *apierror.ApiError) {
	resp, err := r.rc.R().
		SetHeader("Accept", defines.GithubHeaderAcceptJSON).
		SetHeader("Authorization", "bearer "+githubToken).
		SetPathParam(defines.GithubPathParamOwner, owner).
		SetPathParam(defines.GithubPathParamRepository, repo).
		SetPathParam(defines.GithubPathParamAssetID, strconv.FormatUint(assetID, 10)).
		Get(defines.GithubURLGetReleaseAsset)

	if err != nil {
		return nil, apierror.New(http.StatusInternalServerError, err.Error())
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, apierror.New(resp.StatusCode(), string(resp.Body()))
	}

	var asset domain.Asset
	err = json.Unmarshal(resp.Body(), &asset)

	if err != nil {
		return nil, errUnmarshalGetAssetByID
	}

	return &asset, nil
}

// Utils
type writeCounter struct {
	Current uint64
	Total uint64
}
func (wc *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Current += uint64(n)
	wc.PrintProgress()

	return n, nil
}
func (wc *writeCounter) PrintProgress() {
	percent := float32(wc.Current) * float32(100) / float32(wc.Total)
	log.Printf("Downloading... %.2f%% \n", percent)
}