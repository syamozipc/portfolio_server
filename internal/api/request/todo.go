package request

type CreateTodo struct {
	Title string `json:"title" validate:"required"`
}
