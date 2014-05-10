package main

import (
    "github.com/go-martini/martini"
)


func main() {
  m := martini.Classic()
  m.Get("/", Index)
  
  m.Get("/coords", Coords)
  m.Run()
}


func Index() string {
    return "This is index"
}

func Coords() string {
    return "coords"
}
