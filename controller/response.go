package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

// success
func (r *Response) Success(data gin.H) {
	respH := gin.H{
		"msg": "ok",
		"code": 0,
	}

	if len(data) > 1 {
		respH["data"] = data
	} else {
		respH["data"] = data["data"]
	}

	r.Ctx.JSON(http.StatusOK, respH)
}

// error
func (r *Response) Error(err error) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"msg": "error",
		"code": 1,
		"data": err.Error(),
	})
}

// forbidden
func (r *Response) Forbidden() {
	r.Ctx.JSON(http.StatusForbidden, gin.H{
		"msg": "forbidden",
		"code": http.StatusForbidden,
		"data": nil,
	})
	r.Ctx.Abort()
}