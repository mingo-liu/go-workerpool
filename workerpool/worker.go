package workerpool

// Worker 定义工作者
type Worker struct {
	job  chan Job   // WorkerPool为Worker分配任务
	quit chan bool // 用于停止Worker
}

// NewWorker 创建一个Worker对象
func NewWorker() *Worker {
	return &Worker{
		job:  make(chan Job),
		quit: make(chan bool),
	}
}

// Start 启动Worker
func (w *Worker) Start(wp *WorkerPool) {
	go func() {
		for {
			wp.workerQueue <- w // 将自己注册到Worker队列中
			select {
			case job := <-w.job: // 从任务队列中获取任务
				job.Run(nil)
				wp.wg.Done() // 任务完成后调用Done
			case <-w.quit:
				return
			}
		}
	}()
}