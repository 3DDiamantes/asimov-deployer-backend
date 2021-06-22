package repository

import (
	"asimov-deployer-backend/internal/apierror"
	"net/http"
	"os"
	"path/filepath"
)

type FilesystemRepository interface {
	Move(fromPath string, toPath string) *apierror.ApiError
	Run(path string)
}

type filesystemRepository struct {
}

func NewFilesystemRepository() FilesystemRepository {
	return &filesystemRepository{}
}

func (r *filesystemRepository) Move(fromPath string, toPath string) *apierror.ApiError{
	err := os.MkdirAll(filepath.Dir(toPath), os.ModePerm)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}

	err = os.Rename(fromPath, toPath)
	if err != nil {
		return apierror.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}
func (r *filesystemRepository) Run(path string) {

}
