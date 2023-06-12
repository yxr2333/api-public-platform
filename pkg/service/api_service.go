package service

import (
	"api-public-platform/api/request"
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type APIService interface {
	CreateAPI(data request.APICreateRequest) (*model.API, error)
	UpdateAPI(data request.APIUpdateRequest) (*model.API, error)
	DeleteAPI(id uint) error
	GetAPIByID(id uint) (*model.API, error)
	GetAllAPIs(page int, size int) ([]*model.API, error)
	EnableAPI(id uint) error
	DisableAPI(id uint) error
}

type APIServiceImpl struct {
}

func NewAPIService() APIService {
	return &APIServiceImpl{}
}

// 此函数根据 `request.APICreateRequest` 参数中提供的数据在数据库中创建新的 API
// 记录。如果未提供请求数据中的“IsOpen”字段，则默认为“true”。该函数返回一个指向新创建的 `model.API` 记录的指针，如果创建记录有问题则返回一个错误。
func (as *APIServiceImpl) CreateAPI(data request.APICreateRequest) (*model.API, error) {
	// 如果data的IsOpen字段没有传值，那么默认修改为true
	if data.IsOpen == nil {
		*data.IsOpen = true
	}
	newAPI := model.API{
		APIDescription: data.APIDescription,
		APIEndpoint:    data.APIEndpoint,
		RequestMethod:  data.RequestMethod,
		IsOpen:         *data.IsOpen,
		LastCalled:     nil,
	}
	if err := db.MySQLDB.Create(&newAPI).Error; err != nil {
		return nil, fmt.Errorf("failed to create api: %w", err)
	}
	return &newAPI, nil
}

// 此函数使用“data”参数中提供的数据更新数据库中现有的 API 记录。它首先使用 `data.ID` 字段从数据库中检索 API 记录，如果找不到记录则返回错误。然后，它使用 GORM
// 库的“更新”方法，使用“数据”参数中提供的数据更新检索到的记录。最后，它返回指向更新的 API 记录的指针以及更新过程中遇到的任何错误。
func (as *APIServiceImpl) UpdateAPI(data request.APIUpdateRequest) (*model.API, error) {
	api := model.API{
		Model: gorm.Model{
			ID: data.ID,
		},
		APIDescription: data.APIDescription,
		APIEndpoint:    data.APIEndpoint,
		RequestMethod:  data.RequestMethod,
		IsOpen:         *data.IsOpen,
	}
	if err := db.MySQLDB.First(&api, data.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find api with id %d", data.ID)
		}
		return nil, fmt.Errorf("failed to get api: %w", err)
	}
	if err := db.MySQLDB.Model(&api).Updates(data).Error; err != nil {
		return nil, fmt.Errorf("failed to update api: %w", err)
	}
	return &api, nil
}

func (as *APIServiceImpl) DeleteAPI(id uint) error {
	if err := db.MySQLDB.Delete(&model.API{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (as *APIServiceImpl) GetAPIByID(id uint) (*model.API, error) {
	api := model.API{}
	if err := db.MySQLDB.First(api, id).Error; err != nil {
		return nil, fmt.Errorf("could not find api with id %d", id)
	}
	return &api, nil
}
func (as *APIServiceImpl) GetAllAPIs(page int, size int) ([]*model.API, error) {
	var apis []*model.API
	if err := db.MySQLDB.Offset((page - 1) * size).Limit(size).Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("could not find all apis")
	}
	return apis, nil
}
func (as *APIServiceImpl) EnableAPI(id uint) error {
	api := &model.API{}
	if err := db.MySQLDB.First(api, id).Error; err != nil {
		return fmt.Errorf("could not find api with id %d", id)
	}
	api.IsOpen = true
	if err := db.MySQLDB.Save(api).Error; err != nil {
		return fmt.Errorf("could not enable api with id %d", id)
	}
	return nil

}
func (as *APIServiceImpl) DisableAPI(id uint) error {
	api := &model.API{}
	if err := db.MySQLDB.First(api, id).Error; err != nil {
		return fmt.Errorf("could not find api with id %d", id)
	}
	api.IsOpen = false
	if err := db.MySQLDB.Save(api).Error; err != nil {
		return fmt.Errorf("could not disable api with id %d", id)
	}
	return nil
}
