package controller

import (
	"github.com/gin-gonic/gin"
	"web/service"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func GetUsersList(c *gin.Context) {
	ctx := CreateCtx(c)

	skip, _ := c.GetQuery("skip")
	limit, _  := c.GetQuery("limit")

	intSkip, err := strconv.Atoi(skip)

	if err != nil {
		intSkip = 0
	}

	intLimit, err := strconv.Atoi(limit)

	if err != nil {
		intLimit = 10
	}

	userInfo, err := service.GetUsersList(intSkip, intLimit)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(userInfo)
}

func GetUserInfo(c *gin.Context) {
	ctx := CreateCtx(c)

	id := ctx.getParam("id")

	userInfo, err := service.GetUserInfoById(bson.ObjectIdHex(id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(userInfo)
}