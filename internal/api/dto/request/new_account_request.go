package dto

import (
	"leg3nd-agora/internal/core/domain"
)

type NewAccountRequest struct {
	Email         string               `json:"email"`
	FullName      string               `json:"full_name"`
	OAuthProvider domain.OAuthProvider `json:"o_auth_provider"`
}
