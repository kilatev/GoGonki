package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
)

type Point struct {
	X int
	Y int
}

type State struct {
	X       int  `json: "x"`
	Y       int  `json: "y"`
	Crashed bool `json: "crashed"`
	Speed   int  `json: "speed"`
}

type Car struct {
	Color int
	Point Point
	State State
}

func (c Car) MoveTo(point Point) Point {
	c.Point = point
	// ToDo: implement correct alghoritm with use of intesection
	return c.Point
}

func (c Car) Status() State {
	// ToDo: this should return summarized car's status at current time.
	// probably it should contain coords, state, leadership
	return c.State
}

type User struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func main() {
	m := martini.Classic()
	m.Get("/", Index)

	m.Get("/coords", Coords)
	m.Get("/move/:x/:y", Move) // ToDo: make this post reqeust
	m.Get("/state", GetState)
	m.Get("/account/profile", Profile)
	m.Get("/account/login", Login)
	m.Get("/account/logout", Logout)
	m.Run()
}

func Index() string {
	return "This is index"
}

func Coords() string {
	return "coords"
}

func Move(params martini.Params) string { // this is fake function for now. ToDo: change it to use actual Point
	fmt.Println("Scuka kruta")
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
