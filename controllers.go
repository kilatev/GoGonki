package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	//"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/coopernurse/gorp"
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

func Login(params martini.Params, db *gorp.DbMap, s sessions.Session) (int, string) {
	user := User{}
	email := params["email"]
	password := params["password"]
	err := db.SelectOne(&user, "Select * from users where Email=? ", email)
	log.Println(err)
	//err := db.QueryRow("select id, password from users where email=$1", email).Scan(&userId, &dbPasswd)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return 401, "Unauthorized"
	}
	s.Set("userId", user.Id)
	return 200, "User id is " + string(user.Id)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIf(err)

	u := User{
		Name:     name,
		Email:    email,
		Password: hashedPassword}

	//_, err = db.Exec("insert into users (name, email, password) values ($1, $2, $3)",
	err = db.Insert(&u)
	//name, email, hashedPassword)

	PanicIf(err)

	http.Redirect(rw, r, "/login", http.StatusFound)
}

func SignupForm(w http.ResponseWriter, r render.Render) {
	r.HTML(200, "signup", "")
}
