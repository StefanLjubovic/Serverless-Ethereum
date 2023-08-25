package dto

type SignInResponseDto struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}
