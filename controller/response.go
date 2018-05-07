package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

// success handle
func (r *Response) Success(data gin.H) {
	respH := gin.H{
		"msg": "ok",
		"code": 0,
	}

	if len(data) > 1 { // Almost the length is more than 1, so just check it first.
		respH["data"] = data
	} else if data["data"] != nil {
		respH["data"] = data["data"]
	} else {
		respH["data"] = data
	}

	r.Ctx.JSON(http.StatusOK, respH)
}

// error handle
func (r *Response) Error(err error) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"msg": "error",
		"code": 1,
		"data": err.Error(),
	})
}

// forbidden handle
func (r *Response) Forbidden() {
	r.Ctx.JSON(http.StatusForbidden, gin.H{
		"msg": "forbidden",
		"code": http.StatusForbidden,
		"data": nil,
	})
	r.Ctx.Abort()
}