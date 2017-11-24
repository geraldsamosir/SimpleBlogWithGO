package main

import (
	Config "./config"
	AuthService "./service/auth"
	ModelUsers "./service/auth/model"
)

func Migrate() {
	//get connecttion db
	db := Config.Db()

	ModelUsers.DBMigrate(db)
}
