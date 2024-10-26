package service

type SignUpInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
