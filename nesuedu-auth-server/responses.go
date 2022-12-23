package main

type JwtResponse struct {
	Id           uint   `json:"id"`
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refreshToken"`
}

type UserResponse struct {
	Id        uint    `json:"id"`
	Username  string  `json:"username"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Rate      float32 `json:"rate"`
	VkUrl     string  `json:"vkUrl"`
	TgUrl     string  `json:"tgUrl"`
}

type SingleJwtResponse struct {
	Id  uint   `json:"id"`
	JWT string `json:"jwt"`
}
