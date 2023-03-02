package dto

// DTO = Data Transfer Object

// TodoResponse はAPIレスポンスで渡すJSONに詰め替え
type TodoResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// TodoRequest はAPIリクエストで渡されたJSONから詰め替え
type TodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodosResponse struct {
	Todos []TodoResponse `json:"todos"`
}
