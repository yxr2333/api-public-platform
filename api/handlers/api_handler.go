package handlers

import (
	"api-public-platform/api/request"
	"api-public-platform/api/response"
	"api-public-platform/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	apiService service.APIService
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		apiService: service.NewAPIService(),
	}
}

func (ah *APIHandler) CreateAPI(c *gin.Context) {
	var req *request.APICreateRequest
	reqBody, _ := c.Get("reqBody")
	req, ok := reqBody.(*request.APICreateRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "reqBody type error",
		})
		return
	}
	dbRes, err := ah.apiService.CreateAPI(*req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res := response.APICreateResponse{
		APIDescription: dbRes.APIDescription,
		APIEndpoint:    dbRes.APIEndpoint,
		RequestMethod:  dbRes.RequestMethod,
		IsOpen:         dbRes.IsOpen,
		CreatedAt:      dbRes.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"msg":  "success",
		"code": 200,
	})
}

func (ah *APIHandler) UpdateAPI(c *gin.Context) {
	var api *request.APIUpdateRequest
	reqBody, _ := c.Get("reqBody")
	api, ok := reqBody.(*request.APIUpdateRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "reqBody type error",
		})
		return
	}
	_, err := ah.apiService.UpdateAPI(*api)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": nil,
		"msg":  "success",
		"code": 200,
	})
}

func (ah *APIHandler) DeleteAPI(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "id type error",
		})
		c.Abort()
		return
	}
	err = ah.apiService.DeleteAPI(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": nil,
		"msg":  "success",
		"code": 200,
	})
}

func (ah *APIHandler) GetAPIByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "id type error",
		})
		c.Abort()
		return
	}
	dbRes, err := ah.apiService.GetAPIByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res := response.APIGetResponse{
		ID:             dbRes.ID,
		APIDescription: dbRes.APIDescription,
		APIEndpoint:    dbRes.APIEndpoint,
		RequestMethod:  dbRes.RequestMethod,
		IsOpen:         dbRes.IsOpen,
		CreatedAt:      dbRes.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      dbRes.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"msg":  "success",
		"code": 200,
	})
}

func (ah *APIHandler) GetAllAPIs(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	dbRes, err := ah.apiService.GetAllAPIs(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	var res []response.APIGetResponse
	for _, v := range dbRes {
		res = append(res, response.APIGetResponse{
			ID:             v.ID,
			APIDescription: v.APIDescription,
			APIEndpoint:    v.APIEndpoint,
			RequestMethod:  v.RequestMethod,
			IsOpen:         v.IsOpen,
			CreatedAt:      v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"msg":  "success",
		"code": 200,
	})
}

func (ah *APIHandler) EnableAPI(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "id type error",
		})
		c.Abort()
		return
	}
	err = ah.apiService.EnableAPI(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": nil,
		"msg":  "success",
		"code": 200,
	})
}

func (ah *APIHandler) DisableAPI(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "id type error",
		})
		c.Abort()
		return
	}
	err = ah.apiService.DisableAPI(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": nil,
		"msg":  "success",
		"code": 200,
	})
}
