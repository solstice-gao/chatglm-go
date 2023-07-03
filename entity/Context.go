package entity

type ContextResponse struct {
	Message string `json:"message"`
	Result  struct {
		ContextID string `json:"context_id"`
	} `json:"result"`
	Status int `json:"status"`
}
