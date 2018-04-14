package tasks

// ITask is the interface for tasks
type ITask interface {
	Start(interval uint)
	Run()
	Stop()
}
