package handler

type RegisterInput struct {
	Name      string `json:"name"`
	Email     string `json:"email"` 
	Country   string `json:"country"` 
	Password  string `json:"password"`
}

type LoginInput struct {
Email     string `json:"email"` 
Password  string `json:"password"` 
} 

type UpdateInput struct {
	Name     string `json:"name"` 
	Email    string `json:"email"` 
	Country  string `json:"country"`
	Password string `json:"password"`
} 