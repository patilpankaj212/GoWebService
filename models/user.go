package models

import (
	"errors"
	"fmt"
)

// User struct defines a user object
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers method is used to get the list of users
func GetUsers() []*User {
	return users
}

//AddUser method is used to add a new user to the list
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include id or it must be zero")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID returns the user with id passed
func GetUserByID(id int) (User, error) {
	for _, v := range users {
		if v.ID == id {
			return *v, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//UpdateUser updates the user object in the users list
func UpdateUser(u User) (User, error) {
	for i, v := range users {
		if v.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

//RemoveUserByID removes a user with passed id
func RemoveUserByID(id int) error {
	for i, v := range users {
		if v.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
