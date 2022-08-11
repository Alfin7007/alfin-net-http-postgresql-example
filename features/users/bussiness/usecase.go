package bussiness

import (
	"errors"
	"fmt"
	"http/example/features/users"
	"http/example/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBussiness(userData users.Data) users.Bussiness {
	return &userUseCase{
		userData: userData,
	}
}

func (uc userUseCase) Login(userCore users.Core) (id int, token string, err error) {
	result, errLogin := uc.userData.FindUser(userCore.Email)
	if errLogin != nil {
		return 0, "", errLogin
	}
	passCompare := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userCore.Password))
	if passCompare != nil {
		return 0, "", errors.New("wrong password")
	}

	token, _ = middlewares.GenerateToken(result.UserID, result.Email)
	return result.UserID, token, nil
}

func (uc userUseCase) Register(userCore users.Core) (err error) {
	_, userCheck := uc.userData.FindUser(userCore.Email)
	if userCheck == nil {
		return errors.New("email existing")
	} else {
		fmt.Println(userCheck.Error())
	}
	bytePass := []byte(userCore.Password)
	hashPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	userCore.Password = string(hashPass)
	errInsert := uc.userData.InsertUser(userCore)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (uc userUseCase) GetData(id int) (users.Core, error) {

	userCore, nil := uc.userData.SelectUser(id)
	return userCore, nil
}
