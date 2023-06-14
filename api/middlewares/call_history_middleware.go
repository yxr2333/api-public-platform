package middlewares

import (
	"api-public-platform/config"
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriterWrapper struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriterWrapper) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func APICallHistoryMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		callTime := time.Now()
		wrapper := &responseWriterWrapper{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = wrapper // 替换gin框架原来的writer
		c.Next()           // 执行其他中间件或者处理程序

		// 执行其他中间件或者处理程序之后
		// 从响应中获取状态码
		statusCode := c.Writer.Status()
		// 从响应中获取响应体
		responseBody := wrapper.body.String()
		userId := c.GetUint("userId")
		endpoint := strings.TrimPrefix(c.Request.URL.Path, config.ServerCfg.API.Outer.Prefix)
		method := c.Request.Method

		var api model.API
		if err := db.MySQLDB.Where("api_endpoint = ? AND request_method = ?", endpoint, method).First(&api).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// 根据statusCode判断调用是否成功
		var callStatus string
		switch statusCode / 100 {
		case 2:
			callStatus = "success"
		case 4:
			callStatus = "client error"
		case 5:
			callStatus = "server error"
		default:
			callStatus = "unknown error"
		}
		history := model.APICallHistory{
			APIID:        api.ID,
			CalledBy:     userId,
			CalledAt:     callTime,
			CallStatus:   callStatus,
			CallResponse: responseBody,
		}
		if err := db.MySQLDB.Create(&history).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
