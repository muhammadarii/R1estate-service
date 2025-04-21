package schemas

type RoleRequest struct {
	RoleName string `json:"roleName" binding:"required,min=3"`
}
