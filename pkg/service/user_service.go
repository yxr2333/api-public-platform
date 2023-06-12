package service

import (
	"api-public-platform/api/request"
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"api-public-platform/pkg/security"
	"api-public-platform/pkg/utils"
	"crypto/md5"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UserService interface {
	RegisterUser(data request.UserRegisterRequest) (*model.User, error)
	LoginUser(data request.UserLoginRequest) (string, error)
	GetUserByID(id uint) (*model.User, error)
	CreateUser(data request.UserCreateRequest) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	GenerateAPIToken(id uint) (string, error)
	UpdateAPIToken(id uint) (string, error)
}

type UserServiceImpl struct{}

func NewUserService() UserService {
	return &UserServiceImpl{}
}
func (us *UserServiceImpl) UpdateAPIToken(id uint) (string, error) {
	return us.GenerateAPIToken(id)
}

func (us *UserServiceImpl) GenerateAPIToken(id uint) (string, error) {
	var user model.User
	if err := db.MySQLDB.First(&user, id).Error; err != nil {
		log.Printf("get user failed: %v", err)
		return "", fmt.Errorf("get user failed: %v", err)
	}
	token, err := utils.GenerateAPIToken(32)
	if err != nil {
		log.Printf("generate api token failed: %v", err)
		return "", fmt.Errorf("generate api token failed: %v", err)
	}
	user.APIToken = token
	if err := db.MySQLDB.Save(&user).Error; err != nil {
		log.Printf("save user failed: %v", err)
		return "", fmt.Errorf("save user failed: %v", err)
	}
	return token, nil
}

func (us *UserServiceImpl) RegisterUser(data request.UserRegisterRequest) (*model.User, error) {
	var userCheck model.User
	if err := db.MySQLDB.Where("user_name = ? OR email = ?", data.UserName, data.Email).First(&userCheck).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("check user exist failed: %v", err)
			return nil, fmt.Errorf("check user exist failed: %v", err)
		}
	} else {
		return nil, fmt.Errorf("username or email already exist")
	}

	pwd, err := utils.GeneratePasswordHash(data.Password)
	if err != nil {
		log.Printf("generate password hash failed: %v", err)
		return nil, fmt.Errorf("generate password hash failed: %v", err)
	}
	mailHash := md5.Sum([]byte(data.Email))
	user := model.User{
		UserName: data.UserName,
		Password: pwd,
		Email:    data.Email,
		Gender:   data.Gender,
		Avatar:   fmt.Sprintf("https://www.gravatar.com/avatar/%x?identicon", mailHash),
		RoleID:   1, // 默认为普通用户
	}
	if err = db.MySQLDB.Create(&user).Error; err != nil {
		log.Printf("create user failed: %v", err)
		return nil, fmt.Errorf("create user failed: %v", err)
	}
	return &user, nil
}

func (us *UserServiceImpl) LoginUser(data request.UserLoginRequest) (string, error) {
	var user model.User

	// 从数据库中检索用户
	if err := db.MySQLDB.Where("user_name = ?", data.UserName).First(&user).Error; err != nil {
		log.Printf("get user by username failed: %v", err)
		return "", fmt.Errorf("user name or password incorrect")
	}

	// 检查密码
	if ok, err := utils.CheckPasswordHash(data.Password, user.Password); err != nil || !ok {
		log.Printf("check password hash failed: %v", err)
		return "", fmt.Errorf("user name or password incorrect")
	}
	jwtService := security.NewJWTService()
	token := jwtService.GenerateToken(user.UserName, user.ID, true)
	return token, nil
}

func (us *UserServiceImpl) GetUserByID(id uint) (user *model.User, err error) {
	err = db.MySQLDB.First(&user, id).Error
	return
}
func (us *UserServiceImpl) CreateUser(data request.UserCreateRequest) error {
	pwd, err := utils.GeneratePasswordHash(data.Password)
	if err != nil {
		log.Printf("generate password hash failed: %v", err)
		return fmt.Errorf("generate password hash failed: %v", err)
	}
	user := model.User{
		UserName: data.UserName,
		Password: pwd,
		Email:    data.Email,
		Gender:   data.Gender,
	}
	return db.MySQLDB.Create(user).Error
}
func (us *UserServiceImpl) UpdateUser(user *model.User) error {
	return db.MySQLDB.Save(user).Error
}
func (us *UserServiceImpl) DeleteUser(id uint) error {
	user, err := us.GetUserByID(id)
	if err != nil {
		return err
	}
	return db.MySQLDB.Delete(&user).Error
}
