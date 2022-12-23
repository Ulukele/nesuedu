package main

type userAuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type userSetParamsRequest struct {
	userAuthRequest
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type userGetRequest struct {
	Id uint `json:"id" validate:"required"`
}

type AuthRequest struct {
	JWT string `json:"jwt" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
