package poolservice

import (
	"fmt"

	"github.com/panjf2000/ants/v2"
)

var (
	p   *ants.PoolWithFunc
	err error
)

// 初始化协程池
func InitPool() error {
	// 注册工作池，传入任务
	// 参数1 初始化worker并发个数
	if p, err = ants.NewPoolWithFunc(100, func(i interface{}) {
		switch i.(type) {
		case func():
			i.(func())()
		default:
			fmt.Println("暂无可用协程处理方法")
		}
	}); err != nil {
		fmt.Println("初始化协程池失败:", err)
		return err
	}

	return nil
}

func StopPool() {
	p.Release()
}

func AddFuncWork(work func()) {
	_ = p.Invoke(work)
}
