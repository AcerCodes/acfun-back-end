package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
<<<<<<< HEAD
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
=======
	r := gin.Default()
>>>>>>> 3a1a0d8e999d682e585cf7de9a3ca0e239a6379e
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
<<<<<<< HEAD
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
=======
>>>>>>> 3a1a0d8e999d682e585cf7de9a3ca0e239a6379e
	r.Run(":8080")
}
