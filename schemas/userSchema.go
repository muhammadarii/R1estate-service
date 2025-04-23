package schemas

type UserRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=6"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty" `
	Avatar    *string `json:"avatar,omitempty" `
	RoleID    uint    `json:"role_id" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}