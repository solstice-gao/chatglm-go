package entity

type TaskResponse struct {
	Message string `json:"message"`
	Result  struct {
		TaskID string `json:"task_id"`
	} `json:"result"`
	Status int `json:"status"`
}
