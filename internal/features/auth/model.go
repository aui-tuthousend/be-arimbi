package auth

import "be-arimbi/internal/features/user"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User user.UserLoginResponse `json:"user"`
}