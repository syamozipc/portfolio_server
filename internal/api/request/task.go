package request

type GetTask struct {
	ID string `param:"id" validate:"required"`
}

type CreateTask struct {
	Title string `json:"title" validate:"required"`
}
