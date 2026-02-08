package handlers

import "interview-copilot/backend/internal/service"

type Handlers struct {
	Auth *service.AuthService
}

func New(auth *service.AuthService) *Handlers {
	return &Handlers{Auth: auth}
}
