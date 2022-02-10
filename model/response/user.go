package response

type UserResponseCreatedByUsernamePwd struct {
	Username string `json:"username"`
}

type UserResponseProfile struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      int    `json:"role"`
	Gender    int    `json:"gender"`
	Avatar    string `json:"avatar"`
	CreatedAt int    `json:"created_at"`
}

type LoginResponse struct {
	User      UserResponseProfile `json:"user"`
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expires_at"`
}
