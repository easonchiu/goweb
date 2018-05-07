package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func registerUserRouter(g *gin.RouterGroup) {

	g.GET("", /*middleware.Jwt,*/ controller.GetUsersList)

	g.GET("/:id", /*middleware.Jwt,*/ controller.GetUserInfo)

}