package dto

import "leg3nd-agora/internal/core/domain"

type UpdateAccountRequest struct {
	Id            int64                 `json:"id"`
	Email         *string               `json:"email"`
	Nickname      *string               `json:"nickname"`
	FullName      *string               `json:"full_name"`
	OAuthProvider *domain.OAuthProvider `json:"o_auth_provider"`
	Status        *domain.Status        `json:"status"`
}
