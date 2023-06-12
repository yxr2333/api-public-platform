package handlers

import (
	"api-public-platform/api/request"
	"api-public-platform/api/response"
	"api-public-platform/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService service.UserService = service.NewUserService()

func UserRegisterHandler(c *gin.Context) {
	var userRegisterRequest *request.UserRegisterRequest
	reqBody, _ := c.Get("reqBody")
	userRegisterRequest, ok := reqBody.(*request.UserRegisterRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "reqBody type error",
		})
		return
	}
	dbRes, err := userService.RegisterUser(*userRegisterRequest)
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

func UserLoginHandler(c *gin.Context) {
	var userLoginRequest *request.UserLoginRequest
	reqBody, _ := c.Get("reqBody")
	userLoginRequest, ok := reqBody.(*request.UserLoginRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "reqBody type error",
		})
		return
	}
	token, err := userService.LoginUser(*userLoginRequest)
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
