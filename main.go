package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"apiForSrs/api/router"
	"apiForSrs/global"
	"apiForSrs/poolservice"
	"apiForSrs/queue"
)

func main() {
	// 启动协程池
	if err := poolservice.InitPool(); err != nil {
		return
	}

	global.RunnerBox = queue.NewRunner()
	go global.RunnerBox.Start()

	r := router.InitRouter()
	// http 请求
	// 相关进程关闭时
	handleSignal()

	fmt.Println("[info] start http server listening", 9586)

	// 开启服务
	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", 9586)); err != nil {
		fmt.Println(err)
	}

}

func handleSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		<-c

		fmt.Println("清除协程池")
		poolservice.StopPool()

		fmt.Println("清除queue")
		global.RunnerBox.Stop()

		time.Sleep(1 * time.Second)
		fmt.Println("退出程序-----------------------------------**********************----------------------------程序退出")
		os.Exit(0)
	}()
}
