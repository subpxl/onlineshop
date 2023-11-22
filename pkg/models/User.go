package models

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func UserRegister(user User) (string, error) {

	fmt.Println("username, password", user.Username, " ", user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "cannot hash password", err
	}

	// var user User
	user.RoleID = 1
	user.Password = string(hashedPassword)
	resp := DB.Create(&user)

	if resp.Error != nil {
		return "cannot create user", err
	}
	return "user created successfully", nil
}

func UserLogin(username, password string) (bool, error) {
	fmt.Println("username, password", username, " ", username)
	var user User
	DB.Where("username=?", username).Find(&user)
	fmt.Println("found user ", &user)

	status := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if status != nil {
		fmt.Println("password donot match")
		return false, errors.New("invalid password")
	} else {
		return status == nil, nil
	}
}

func GetUserByUsername(username string) User {
	var user User
	DB.Where("username=?", username).Find(&user)
	return user
}
