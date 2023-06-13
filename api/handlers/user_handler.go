package handlers

import (
	"api-public-platform/api/request"
	"api-public-platform/api/response"
	"api-public-platform/pkg/model"
	"api-public-platform/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler() UserHandler {
	return UserHandler{
		userService: service.NewUserService(),
	}
}

func (uh *UserHandler) UserRegisterHandler(c *gin.Context) {
	var userRegisterRequest *request.UserRegisterRequest
	reqBody, _ := c.Get("reqBody")
	userRegisterRequest, ok := reqBody.(*request.UserRegisterRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "reqBody type error",
		})
		return
	}
	dbRes, err := uh.userService.RegisterUser(*userRegisterRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	userResponse := response.UserRegisterResponse{
		UserName: dbRes.UserName,
		Email:    dbRes.Email,
		Gender:   dbRes.Gender,
		Avatar:   dbRes.Avatar,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
		"msg":  "success",
		"code": 200,
	})
}

func (uh *UserHandler) UserLoginHandler(c *gin.Context) {
	var userLoginRequest *request.UserLoginRequest
	reqBody, _ := c.Get("reqBody")
	userLoginRequest, ok := reqBody.(*request.UserLoginRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "reqBody type error",
		})
		return
	}
	token, err := uh.userService.LoginUser(*userLoginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	userResponse := response.UserLoginResponse{
		UserName: userLoginRequest.UserName,
		Token:    token,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
		"msg":  "success",
		"code": 200,
	})
}

func (uh *UserHandler) GetOneUserHandler(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id type error",
		})
		return
	}
	dbRes, err := uh.userService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res := response.UserBaseInfo{
		UserName: dbRes.UserName,
		Email:    dbRes.Email,
		Gender:   dbRes.Gender,
		Avatar:   dbRes.Avatar,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"msg":  "success",
		"code": 200,
	})
}

// CreateUserHandler handles the creation of a user.
func (uh *UserHandler) CreateUserHandler(c *gin.Context) {
	var userCreateRequest *request.UserCreateRequest
	reqBody, _ := c.Get("reqBody")
	userCreateRequest, ok := reqBody.(*request.UserCreateRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "reqBody type error",
		})
		return
	}
	if err := uh.userService.CreateUser(*userCreateRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "User created successfully",
		"code": 200,
	})
}

// UpdateUserHandler handles the updating of a user.
func (uh *UserHandler) UpdateUserHandler(c *gin.Context) {
	var user *model.User
	reqBody, _ := c.Get("reqBody")
	user, ok := reqBody.(*model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "reqBody type error",
		})
		return
	}
	if err := uh.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "User updated successfully",
		"code": 200,
	})
}

// DeleteUserHandler handles the deletion of a user.
func (uh *UserHandler) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := uh.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "User deleted successfully",
		"code": 200,
	})
}

// GenerateAPITokenHandler handles the generation of an API token for a user.
func (uh *UserHandler) GenerateAPITokenHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	token, err := uh.userService.GenerateAPIToken(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"msg":   "API Token generated successfully",
		"code":  200,
	})
}

// UpdateAPITokenHandler handles the updating of an API token for a user.
func (uh *UserHandler) UpdateAPITokenHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	token, err := uh.userService.UpdateAPIToken(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"msg":   "API Token updated successfully",
		"code":  200,
	})
}
