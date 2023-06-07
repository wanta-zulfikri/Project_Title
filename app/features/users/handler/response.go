package handler 

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Password   string `json:"password"`
}