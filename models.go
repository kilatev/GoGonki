package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"time"
)

var channels map[Room]chan int = make(map[Room]chan int)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=gogonki sslmode=disable")
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id       int64  `db: "id"`
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

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

type Room struct {
	room_id      int32
	room_channel string
}

func (r Room) NewRoom() (room Room) {
	room = Room{}
	room.room_id = room.GenerateId()
	return
}

func SendToRoom(room_id int64, car Car) {
	ch := make(chan<- Car) // ToDo move channels to the channel dict or something
	ch <- car
}

func (r Room) GenerateId() int32 {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int31()
}

func (u User) Save() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.Insert(u)
	checkErr(err, "User saving failed")
}

func (u User) Get(id int) User {
	dbmap := initDb()
	defer dbmap.Db.Close()
	user := User{}
	err := dbmap.SelectOne(&user, "Select * from users where id=?", id)
	checkErr(err, "User is unreachable")
	return user
}

func (u User) GetByEmail(email string) User {
	dbmap := initDb()
	defer dbmap.Db.Close()
	user := User{}
	err := dbmap.SelectOne(&user, "Select * from users where email=?", email)
	checkErr(err, "User is unreachable")
	return user
}

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

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./tmp/post_db.bin")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(User{}, "posts").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
