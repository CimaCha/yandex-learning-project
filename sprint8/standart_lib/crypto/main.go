package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Username string
	Password string
}

func CreateUser(username, rawPassword string) (User, error) {
	// Реализуйте создание структуры пользователя.
	// Структура должна содержать оригинальное имя пользователя и хешированный пароль.
	// Используйте логин пользователя в качестве соли при хешировании.
	bytes, err := bcrypt.GenerateFromPassword([]byte(username+rawPassword), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}
	return User{Username: username, Password: string(bytes)}, nil
}

func Authenticate(u User, rawPassword string) (bool, error) {
	// Реализуйте проверку пароля для пользователя.
	// Обратите внимание: в структуре должен быть хеш пароля, а в аргументе функции — оригинальный пароль.
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.Username+rawPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}

func main() {
	rawPassword := "554eB5ZbwV5ck2s8"
	user, err := CreateUser("anonymous", rawPassword)
	if err != nil {
		log.Fatal(err)
	}

	if user.Password == rawPassword {
		log.Fatal("user.Password is raw")
	}

	isAuthenticated, err := Authenticate(user, rawPassword)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User %s has authenticated ? %v", user.Username, isAuthenticated)
}
