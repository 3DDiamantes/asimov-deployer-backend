package service

import (
	"asimov-deployer-backend/internal/apierror"
	"asimov-deployer-backend/internal/domain"
	"github.com/stretchr/testify/mock"
)

type DeployerServiceMock struct {
	mock.Mock
}

func (m *DeployerServiceMock) Deploy(body *domain.DeployBody, githubToken *string) *apierror.ApiError {
	args := m.Called(body, githubToken)

	err := args.Get(0)

	if err != nil {
		return err.(*apierror.ApiError)
	}

	return nil
}