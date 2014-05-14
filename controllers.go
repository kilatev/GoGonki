package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"log"
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

func Login(params martini.Params) (result string) {
	result = "Login"
	return
}

func Logout(params martini.Params) (result string) {
	result = "Logout"
	return
}
