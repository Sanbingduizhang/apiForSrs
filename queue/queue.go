package queue

import (
	"fmt"
)

type Runner struct {
	stop chan int

	tasks     chan func()
	mailTasks chan func()
}

func NewRunner() *Runner {
	return &Runner{
		stop:  make(chan int),
		tasks: make(chan func(), 100),
	}
}

func (r *Runner) AddTask(task func()) {
	r.tasks <- task
}

// 显式执行go
// 开始
func (r *Runner) Start() {
	for {
		select {
		case fn := <-r.tasks:
			fmt.Println("开始执行一个普通任务")
			fn()
		case <-r.stop:
			close(r.tasks)
			fmt.Println("关闭任务执行")
			return
		}
	}
}

// 关闭

func (r *Runner) Stop() {
	r.stop <- 1
}
