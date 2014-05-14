package main

import (
	"encoding/json"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"log"
)

func (c Car) MoveTo(point Point) Point {
	c.Point = point
	// ToDo: implement correct algorithm with use of intersection
	// Assume collision detection done on client side for now
	return c.Point
}

func (c Car) Status() State {
	// ToDo: this should return summarized car's status at current time.
	// probably it should contain coords, state, leadership
	return c.State
}

func (u User) GetRoom() string {
	// ToDo this should return users room data, or room itself.
	return "room data"
}

func (u User) MoveToRoom() {
	// ToDo: this will place user to new room
}

func main() {
	m := martini.Classic()
	m.Get("/", Index)

	m.Get("/coords", Coords)
	m.Get("/move/:x/:y", Move) // ToDo: make this post request
	m.Get("/state", GetState)
	m.Get("/account/profile", Profile)
	m.Get("/account/login", Login)
	m.Get("/account/logout", Logout)
	m.Post("/account/create", binding.Bind(User{}), CreateUser)
	m.Run()
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
