package auth

import (
	Config "../../config"
	Helper "../helper"
	Model "./model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

var db = Config.Db()

func Getallusers(res http.ResponseWriter, req *http.Request) {

	var _users []Model.User
	var _user Model.User
	converqUrl := schema.NewDecoder()

	if err := converqUrl.Decode(&_user, req.URL.Query()); err != nil {
		fmt.Println(err)
		return
	}

	db.Where(&_user).Find(&_users)

	Helper.RespondJSON(res, http.StatusOK, _users)
}

func CreateUsers(res http.ResponseWriter, req *http.Request) {
	data := json.NewDecoder(req.Body)
	var _user Model.User
	if err := data.Decode(&_user); err != nil {
		Helper.RespondError(res, http.StatusBadRequest, err.Error())
		return
	}

	//for access  value exp :
	//_user.Email

	if err := db.Create(&_user).Error; err != nil {
		Helper.RespondError(res, http.StatusInternalServerError, err.Error())
		return
	}

	Helper.RespondJSON(res, http.StatusCreated, _user)
}
