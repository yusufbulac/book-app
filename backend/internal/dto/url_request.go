package dto

type UrlProcessRequest struct {
	URL       string `json:"url" validate:"required,url"`
	Operation string `json:"operation" validate:"required,oneof=canonical redirection all"`
}
