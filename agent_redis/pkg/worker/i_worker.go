package worker

type WorkerResponse struct{
	DType int
	Data []byte
}

type IWorker interface{
	Work(args ...interface{}) (*WorkerResponse, error)
	GetName() string
}