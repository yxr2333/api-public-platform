package service_test

import (
	"api-public-platform/api/request"
	"api-public-platform/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAPIService struct {
	mock.Mock
}

func (m *MockAPIService) CreateAPI(data request.APICreateRequest) (*model.API, error) {
	args := m.Called(data)
	return args.Get(0).(*model.API), args.Error(1)
}

func TestCreateAPI(t *testing.T) {
	mockService := new(MockAPIService)
	data := request.APICreateRequest{
		APIDescription: "Test API",
	}
	mockService.On("CreateAPI", data).Return(&model.API{
		APIDescription: "Test API",
	}, nil)
	api, err := mockService.CreateAPI(data)
	assert.NoError(t, err)
	assert.NotNil(t, api)
	assert.Equal(t, "Test API", api.APIDescription)
}
