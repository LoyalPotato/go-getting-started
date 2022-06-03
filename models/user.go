package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include ID")
	}
	u.ID = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with id %v not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID  == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id %v not found", id) 
}