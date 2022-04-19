package main

// import (
// 	"log"
// 	"time"

// 	"github.com/robfig/cron"

// 	"github.com/EDDYCJY/go-gin-example/models"
// )

// func main() {
// 	log.Println("Starting...")
// 	//会根据本地时间创建一个新（空白）的 Cron job runner
// 	//c := cron.New() 应该改为 c := cron.New(cron.WithSeconds())，否则不支持秒级别的定时任务。
// 	c := cron.New()
// 	//AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行
// 	c.AddFunc("* * * * * *", func() {
// 		log.Println("Run models.CleanAllTag...")
// 		models.CleanAllTag()
// 	})
// 	c.AddFunc("* * * * * *", func() {
// 		log.Println("Run models.CleanAllArticle...")
// 		models.CleanAllArticle()
// 	})
// 	//c.Start()
// 	//在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制哦
// 	c.Start()

// 	t1 := time.NewTimer(time.Second * 10)
// 	for {
// 		select {
// 		case <-t1.C:
// 			t1.Reset(time.Second * 10)
// 		}
// 	}
// }
