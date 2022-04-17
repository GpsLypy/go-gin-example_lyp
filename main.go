package main

import (
	"fmt"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers"
	_ "github.com/gin-gonic/gin"
)

//关键词：都为处理器但责任划分不同
func main() {
	//多路复用器：根据请求的URL将请求重定向到不同的处理器
	//相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
	//router := gin.Default()
	//Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，
	//在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
	// router.GET("/test", func(c *gin.Context) {
	// 	//gin.H{…}：就是一个map[string]interface{}
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router, //为server对象指定处理器，多路复用器(可以理解为总的请求处理器)
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("......")
	s.ListenAndServe()
}
