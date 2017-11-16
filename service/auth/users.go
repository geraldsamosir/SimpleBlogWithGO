package auth

import (
	"net/http"
)

type Users struct {
	Username string `form:"name"`
	Email    string `form:"Email"`
}

type UsersService struct{}

func Getallusers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("users"))
}
