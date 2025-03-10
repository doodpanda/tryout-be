package user

import (
	"github.com/doodpanda/tryout-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func RegisterUserToParam(req RegisterUserRequest, param *repository.InsertUserParams) error {
	param.Email = req.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	param.Password = string(hashedPassword)
	param.FirstName = req.FirstName
	param.LastName = req.LastName
	return nil
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
