package task

type RequestAddTask struct {
	Name string `json:"name"`
}

type ResponseAddTask struct {
	Name string `json:"name"`
}
