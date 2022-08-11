package data

import (
	"http/example/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user User) ToCore() users.Core {
	userCore := users.Core{
		UserID:   int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return userCore
}

func fromCore(userCore users.Core) User {
	userModel := User{
		Name:     userCore.Name,
		Email:    userCore.Email,
		Password: userCore.Password,
	}
	return userModel
}
