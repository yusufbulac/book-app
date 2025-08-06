package dto

type CreateBookRequest struct {
	Title  string `json:"title" validate:"required,min=1,max=255"`
	Author string `json:"author" validate:"required,min=1,max=255"`
	Year   int    `json:"year" validate:"required,min=0"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" validate:"omitempty,min=1,max=255"`
	Author string `json:"author" validate:"omitempty,min=1,max=255"`
	Year   int    `json:"year" validate:"omitempty,min=0"`
}
