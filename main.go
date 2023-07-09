package main

import (
	"github.com/gin-gonic/gin"

	sqlconnect "acfun-back-end/pubfun"
)

func setupRouter() *gin.Engine {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	sqlconnect.DbOpen()
	sqlconnect.Query()
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8080")
}
