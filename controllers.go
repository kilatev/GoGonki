package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
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

func LoginForm(w http.ResponseWriter, r render.Render) {
	r.HTML(200, "login", "")
}

func Logout(params martini.Params) (result string) {
	result = "Logout"
	return
}

func Signup(rw http.ResponseWriter, r *http.Request, db *gorp.DbMap) {

	name, email, password := r.FormValue("name"), r.FormValue("email"), r.FormValue("password")
	u := User{name, email, password}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIf(err)

	//_, err = db.Exec("insert into users (name, email, password) values ($1, $2, $3)",
	err = db.Insert(&u)
	//name, email, hashedPassword)

	PanicIf(err)

	http.Redirect(rw, r, "/login", http.StatusFound)
}

func SignupForm(w http.ResponseWriter, r render.Render) {
	r.HTML(200, "signup", "")
}
