package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()
	dbMap := initDb()
	m.Map(dbMap)
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
	m.Get("/signup", SignupForm)
	m.Post("/signup", Signup)
	m.Post("/create-room", CreateRoom)
	m.Post("/list-rooms", ListRooms)

	//m.Use(Auth)
	m.Run()
}
