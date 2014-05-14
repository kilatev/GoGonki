package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
	"log"
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

type Road struct {
	InnerRoad []Point
	OuterRoad []Point
}

type Car struct {
	Color int
	Point Point
	State State
}

func (c Car) MoveTo(point Point) Point {
	c.Point = point
	// ToDo: implement correct algorithm with use of intersection
	return c.Point
}

func (c Car) Status() State {
	// ToDo: this should return summarized car's status at current time.
	// probably it should contain coords, state, leadership
	return c.State
}

type User struct {
	Id       int64  `db: "id"`
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
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

func CreateUser(user User) User {
	dbmap := initDb()
	defer dbmap.Db.Close()

	return
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./tmp/post_db.bin")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}
