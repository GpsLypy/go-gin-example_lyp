// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/EDDYCJY/go-gin-example/pkg/setting"
// 	"github.com/EDDYCJY/go-gin-example/routers"
// )

// //关键词：都为处理器但责任划分不同
// func main() {
// 	//多路复用器：根据请求的URL将请求重定向到不同的处理器
// 	//相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
// 	//router := gin.Default()
// 	//Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，
// 	//在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
// 	// router.GET("/test", func(c *gin.Context) {
// 	// 	//gin.H{…}：就是一个map[string]interface{}
// 	// 	c.JSON(200, gin.H{
// 	// 		"message": "test",
// 	// 	})
// 	// })
// 	router := routers.InitRouter()
// 	s := &http.Server{
// 		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
// 		Handler:        router, //为server对象指定处理器，多路复用器(可以理解为总的请求处理器)
// 		ReadTimeout:    setting.ReadTimeout,
// 		WriteTimeout:   setting.WriteTimeout,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	fmt.Println("......")
// 	s.ListenAndServe()
// }

package main

import (
	_ "context"
	"fmt"
	"log"
	_ "net/http"
	_ "os"
	_ "os/signal"
	"syscall"
	_ "time"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers"
	"github.com/fvbock/endless"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	//节点
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"time"

// 	"github.com/EDDYCJY/go-gin-example/models"
// 	"github.com/EDDYCJY/go-gin-example/pkg/logging"
// 	"github.com/EDDYCJY/go-gin-example/pkg/setting"
// 	"github.com/EDDYCJY/go-gin-example/routers"
// )

// func main() {
//     // setting.Setup()
// 	// models.Setup()
// 	// logging.Setup()

// 	router := routers.InitRouter()

// 	s := &http.Server{
// 		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
// 		Handler:        router,
// 		ReadTimeout:    setting.ReadTimeout,
// 		WriteTimeout:   setting.WriteTimeout,
// 		MaxHeaderBytes: 1 << 20,
// 	}

// 	go func() {
// 		if err := s.ListenAndServe(); err != nil {
// 			log.Printf("Listen: %s\n", err)
// 		}
// 	}()

// 	quit := make(chan os.Signal)
// 	signal.Notify(quit, os.Interrupt)
// 	<-quit

// 	log.Println("Shutdown Server ...")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := s.Shutdown(ctx); err != nil {
// 		log.Fatal("Server Shutdown:", err)
// 	}

// 	log.Println("Server exiting")
// }
