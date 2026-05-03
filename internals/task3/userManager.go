package task3

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

type UserManager struct {
	AllUsers map[int]*User
}

func (um *UserManager) AddUser(id int, name string) {
	_, ok := um.AllUsers[id]
	if ok {
		fmt.Println("user dengan id:", id, "sudah ada")
	} else {
		um.AllUsers[id] = &User{id, name}
	}
}

func (um *UserManager) GetUser(id int) (*User, error) {
	user, ok := um.AllUsers[id]
	if ok {
		return user, nil
	} else {
		return nil, fmt.Errorf("data tidak ditemukan")
	}
}

func NewUserManager() *UserManager {
	return &UserManager{AllUsers: map[int]*User{}}
}
