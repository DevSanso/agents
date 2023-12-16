package worker


type IWorker interface{
	Work(args ...interface{}) ([]byte, error)
}