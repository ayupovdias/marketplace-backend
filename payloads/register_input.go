package payloads

type RegisterInput struct {
	Username        string `json:"username" binding:"required,min=3"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=4"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
