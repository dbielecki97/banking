package dto

type VerifyResponse struct {
	IsAuthorized bool   `json:"is_authorized,omitempty"`
	Message      string `json:"message,omitempty"`
}
