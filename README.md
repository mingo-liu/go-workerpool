# go-workerpool

## 简介
go-workerpool 是一个轻量级、高效的 Go 语言协程池库，简化高并发任务的调度和执行。
## 安装

```shell
go get -u github.com/mingo-liu/go-workerpool@latest
```
## 例子
```go
package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/mingo-liu/go-workerpool/workerpool"
)

// Task 定义具体的任务结构体
type Task struct {
	Number int
	Data   any // 函数参数
}

// Run 实现Job接口的Run方法
func (t Task) Run() {
	fmt.Printf("This is task: %d, Data: %v\n", t.Number, t.Data)
	time.Sleep(1 * time.Second) // 模拟任务执行时间
}

func main() {
	// 设置协程池的大小
	poolNum := 10
	jobQueueNum := 20
	workerPool := workerpool.NewWorkerPool(poolNum, jobQueueNum)
	workerPool.Start()

	// 模拟请求
	taskNum := 100

	// 提交任务
	for i := range taskNum {
		task := Task{Number: i, Data: fmt.Sprintf("Data from Task %d", i)}
		workerPool.Submit(task)
	}

	// 等待所有任务完成
	workerPool.Wait()

	// 打印当前Goroutine数量
	fmt.Println("All tasks completed. Current Goroutines:", runtime.NumGoroutine())
}
```
