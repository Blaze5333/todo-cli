package user

import (
	"errors"
	"fmt"

	"github.com/Blaze5333/todo-cli/internal/storage"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var strg = storage.NewStorage[[]User]("users.json")

func Login(username, password string) (*User, error) {
	users := []User{}
	err := strg.Load(&users)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		} else if user.Username == username && user.Password != password {
			return nil, errors.New("incorrect password")
		}
	}
	return nil, errors.New("user not found")
}

func Register(username, password string) (*User, error) {
	users := []User{}
	err := strg.Load(&users)

	if err != nil {
		fmt.Println("Error loading users:", err)
		return nil, err
	}
	for _, user := range users {
		if user.Username == username {
			return nil, errors.New("username already exists")
		}
	}
	newuser := &User{
		Username: username,
		Password: password,
	}
	users = append(users, *newuser)
	err = strg.Save(users)
	if err != nil {
		fmt.Println("Error saving users:", err)
		return nil, err
	}
	return newuser, nil
}
