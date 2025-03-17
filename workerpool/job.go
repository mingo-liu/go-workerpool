package workerpool

// Job 定义任务接口
type Job interface {
	Run(request any)
}