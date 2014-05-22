package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"log"
  "code.google.com/p/go.crypto/bcrypt"
	_ "github.com/lib/pq"
	"database/sql"
  "github.com/martini-contrib/sessions"
	//"text/template"
)

func CreateUser(user User) (result string) {
	user.Save()
	return
}

func Index() string {
	log.Println("This is index page")
	return "This is index"
}

func Coords() string {
	return "coords"
}

func Move(params martini.Params) string { // this is fake function for now. ToDo: change it to use actual Point
	return "Hello " + params["x"] + params["y"]
}

func GetState(params martini.Params) (result string) { // Todo: change return type to json
	current_state := State{}
	serialized_state, err := json.Marshal(current_state)
	if err != nil {
		result = "{'status', 'failed'}"
	} else {
		result = string(serialized_state)
	}
	return
}

func Profile(params martini.Params) (result string) {
	result = "Profile"
	return
}

func Login(params martini.Params, db *sql.DB, s sessions.Session) (int, string) {
  var userId int64
  var dbPasswd string
  email := params["email"]
  password := params["password"]
  err := db.QueryRow("select id, password from users where email=$1", email).Scan(&userId, &dbPasswd)
  if err != nil || bcrypt.CompareHashAndPassword([]byte(dbPasswd), []byte(password)) != nil {
		return 401, "Unauthorized"
	}
	s.Set("userId", userId)
  return 200, "User id is " + string(userId)
}

func Logout(params martini.Params) (result string) {
	result = "Logout"
	return
}
