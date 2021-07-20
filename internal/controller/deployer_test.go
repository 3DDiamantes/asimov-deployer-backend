package controller

import (
	"asimov-deployer-backend/internal/apierror"
	"asimov-deployer-backend/internal/defines"
	"asimov-deployer-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeployerController_Deploy_OK(t *testing.T) {
	// Given
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	srv := new(service.DeployerServiceMock)
	srv.On("Deploy", mock.Anything, mock.Anything).Return(nil)
	ctrl := NewDeployerController(srv)

	json := `{"owner":"a","repo":"b","tag":"c","scope":"d"}`

	headers := http.Header{defines.HeaderGithubToken:{"****"}}

	jsonReader := strings.NewReader(json)
	ioReadCloser := ioutil.NopCloser(jsonReader)
	ctx.Request = &http.Request{
		Body: ioReadCloser,
		Header: headers,
	}

	// When
	ctrl.Deploy(ctx)

	// Then
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `{"message":"deployed correctly"}`, response.Body.String())
}
func TestDeployerController_Deploy_InvalidBodyError(t *testing.T) {
	// Given
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	srv := new(service.DeployerServiceMock)
	srv.On("Deploy", mock.Anything, mock.Anything).Return(nil)
	ctrl := NewDeployerController(srv)

	json := `{}`

	headers := http.Header{defines.HeaderGithubToken:{"****"}}

	jsonReader := strings.NewReader(json)
	ioReadCloser := ioutil.NopCloser(jsonReader)
	ctx.Request = &http.Request{
		Body: ioReadCloser,
		Header: headers,
	}

	// When
	ctrl.Deploy(ctx)

	// Then
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, `{"Status":400,"Message":"invalid body"}`, response.Body.String())
}
func TestDeployerController_Deploy_MissingGitHubToken(t *testing.T) {
	// Given
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	srv := new(service.DeployerServiceMock)
	srv.On("Deploy", mock.Anything, mock.Anything).Return(nil)
	ctrl := NewDeployerController(srv)

	json := `{"owner":"a","repo":"b","tag":"c","scope":"d"}`

	jsonReader := strings.NewReader(json)
	ioReadCloser := ioutil.NopCloser(jsonReader)
	ctx.Request = &http.Request{
		Body: ioReadCloser,
	}

	// When
	ctrl.Deploy(ctx)

	// Then
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, `{"Status":400,"Message":"missing token"}`, response.Body.String())
}
func TestDeployerController_Deploy_ServiceError(t *testing.T) {
	// Given
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	srv := new(service.DeployerServiceMock)
	apiErr := apierror.New(http.StatusInternalServerError, "service error")
	srv.On("Deploy", mock.Anything, mock.Anything).Return(apiErr)
	ctrl := NewDeployerController(srv)

	json := `{"owner":"a","repo":"b","tag":"c","scope":"d"}`

	headers := http.Header{defines.HeaderGithubToken:{"****"}}

	jsonReader := strings.NewReader(json)
	ioReadCloser := ioutil.NopCloser(jsonReader)
	ctx.Request = &http.Request{
		Body: ioReadCloser,
		Header: headers,
	}

	// When
	ctrl.Deploy(ctx)

	// Then
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, `{"Status":500,"Message":"service error"}`, response.Body.String())
}