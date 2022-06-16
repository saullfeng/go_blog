package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Count interface{} `json:"count,omitempty"`
}

//返回请求数据
func (res *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, res)
	return
}
