package model

import (
	"fmt"
	"net/http"

	"github.com/XecFardin/go_tamplate/utils"
)

var (
	users = map[int64]*User{
		1: &User{
			Id:        1,
			FirstName: "Abdulla",
			LastName:  "Shaikh",
			Email:     "abdullaxec@gmail.com",
		},
		2: &User{
			Id:        2,
			FirstName: "Fardin",
			LastName:  "Shaikh",
			Email:     "abdullaxec@gmail.com",
		},
		3: &User{
			Id:        3,
			FirstName: "Ibtisaam",
			LastName:  "Shaikh",
			Email:     "ibtixec@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {

	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
