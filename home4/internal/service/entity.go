package service

// TaskRequest - структура, представляющая тело запроса
type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type TaskReqId struct {
	ID int `json:"id"`
}
