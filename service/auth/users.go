package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `form:"name"`
	Email    string `form:"Email"`
}

var Users = []User{
	User{"gerald", "gerald@mail.com"},
	User{"dimas", "dimas@mail.com"},
}

type UsersService struct{}

func Getallusers(res http.ResponseWriter, req *http.Request) {
	var result, err = json.Marshal(Users)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(result)
}

func CreateUsers(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)

	var _user User
	err := decoder.Decode(&_user)
	if err != nil {
		panic(err)
	}

	fmt.Println(_user.Username)

}
