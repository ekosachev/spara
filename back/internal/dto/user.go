package dto

type User struct {
	Model
	Email        string
	Username     string
	PasswordHash string
}
