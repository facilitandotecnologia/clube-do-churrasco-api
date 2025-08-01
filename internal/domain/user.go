package domain

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	WhatsApp string `json:"whatsapp"`
	Password string `json:"password" binding:"required,min=6"`
}
