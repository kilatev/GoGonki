package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
)

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
