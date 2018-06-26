package router

import "github.com/gin-gonic/gin"

func Register(g *gin.Engine) {

	// 这是一个测试的请求处理
	registerDemoRouter(g.Group("/demo"))

}
