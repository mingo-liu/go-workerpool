package main

import (
	"fmt"
	"go-workerpool/workerpool"
	"runtime"
	"time"
)

// Task 定义具体的任务结构体
type Task struct {
	Number int
}

// Run 实现Job接口的Run方法
func (t Task) Run(request any) {
	fmt.Println("This is task: ", t.Number)
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
		task := Task{Number: i}
		workerPool.Submit(task)
	}

	// 等待所有任务完成
	workerPool.Wait()

	// 打印当前Goroutine数量
	fmt.Println("All tasks completed. Current Goroutines:", runtime.NumGoroutine())
}
