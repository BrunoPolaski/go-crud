package request

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Age      int8   `json:"age" binding:"required"`
}
