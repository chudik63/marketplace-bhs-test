package service

import "time"

type SignUpInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken     string
	AccessTokenTTL  time.Duration
	RefreshToken    string
	RefreshTokenTTL time.Duration
}
