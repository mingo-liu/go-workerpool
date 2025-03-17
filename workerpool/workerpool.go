package workerpool

import (
	"sync"
)

// WorkerPool 定义协程池
type WorkerPool struct {
	workerNum   int          // 协程池的大小
	workerQueue chan *Worker // Worker队列
	jobQueue    chan Job     // 任务队列
	wg          sync.WaitGroup
}

// NewWorkerPool 创建一个协程池
func NewWorkerPool(workerNum int, jobQueueNum int) *WorkerPool {
	return &WorkerPool{
		workerNum:   workerNum,
		workerQueue: make(chan *Worker, workerNum),
		jobQueue:    make(chan Job, jobQueueNum),
	}
}

// Start 启动协程池
func (wp *WorkerPool) Start() {
	// 初始化Worker
	for range wp.workerNum {
		worker := NewWorker()
		worker.Start(wp)
	}

	// 分配任务
	go func() {
		for job := range wp.jobQueue {
			worker := <-wp.workerQueue
			worker.job <- job
		}
	}()
}

// Submit 提交任务到协程池
func (wp *WorkerPool) Submit(job Job) {
	wp.wg.Add(1)
	wp.jobQueue <- job
}

// Wait 等待所有任务完成
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.jobQueue)
	close(wp.workerQueue)
}
