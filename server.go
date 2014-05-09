package main

import "github.com/go-martini/martini"
import "GoGonki/app/controllers/app/controllers"

func main() {
  m := martini.Classic()
  m.Get("/", Index())
  
  m.Run("/coords", Coords())
}
