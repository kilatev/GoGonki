package main

import (
	"fmt"
	"github.com/go-martini/martini"
)

type Point struct {
	X int
	Y int
}

type Car struct {
	Color int
	Point Point
}

type State struct {
  X int `json: "x"`
  Y int `json: "y"`
  Crashed bool `json: "crashed"`
  Speed int `json: "speed"`
}

func main() {
	m := martini.Classic()
	m.Get("/", Index)

	m.Get("/coords", Coords)
	m.Get("/move/:x/:y", Move) // ToDo: make this post reqeust
	m.Get("/state", State)
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

func State(params martini.Params) string (result string) { // Todo: change return type to json
  current_state := State()
  serialized_state, err := json.Marshal(current_state)
  if err != nil {
    result = "{'status', 'failed'}"
  }
  else
  {
    result = string(jsonObj)
  }
	return 
}
