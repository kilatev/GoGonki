package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()
	m.Map(SetupDB())
	store := sessions.NewCookieStore([]byte("nonotestetstsst"))
	m.Use(sessions.Sessions("gogonki", store))
	m.Use(render.Renderer())
	m.Get("/", Index)

	m.Get("/coords", Coords)
	m.Get("/move/:x/:y", Move) // ToDo: make this post request
	m.Get("/state", GetState)
	m.Get("/profile", Profile)
	m.Get("/login", LoginForm)
	m.Post("/login", Login)
	m.Get("/logout", Logout)
	m.Post("/create", binding.Bind(User{}), CreateUser)
	m.Post("signup", Signup)

	//m.Use(Auth)
	m.Run()
}
