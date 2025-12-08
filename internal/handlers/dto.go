package handlers
type User struct{
	id uint `gorm:"primaryKey" json:"id"`
	username string `json:"username"`
	email string `json:"email"`
	password string `json:"password"`
	isActive bool `json:"is_active"`	
	createdAt  int64 `json:"created_at"`
	updatedAt  int64 `json:"updated_at"`
}
