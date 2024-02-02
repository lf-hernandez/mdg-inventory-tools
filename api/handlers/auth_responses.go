package handlers

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type SignupResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
