package auth

import (
	Config "../../config"
	Helper "../helper"
	Model "./model"
	"net/http"
)

func Getallusers(res http.ResponseWriter, req *http.Request) {
	var _user Model.User
	db := Config.Db()
	db.Find(&_user)
	Helper.RespondJSON(res, http.StatusOK, _user)
}

func CreateUsers(res http.ResponseWriter, req *http.Request) {

}
