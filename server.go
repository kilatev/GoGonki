package main

import (
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

func main() {
  m := martini.Classic()
  m.Get("/", Index)
  
  m.Get("/coords", Coords)
  m.Get("/move/:x/:y", Move) // ToDo: make this post reqeust
  m.Run()
}


func Index() string {
    return "This is index"
}

func Coords() string {
    return "coords"
}

func Move(params martini.Params) (string) { // this is fake function for now. ToDo: change it to use actual Point
    return "Hello " + params["x"] + params["y"]
}
