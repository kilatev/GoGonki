package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

/*type Point struct {
  X int
  Y int
}

type Car struct {
  Color int
  Point Point
}*/

type User struct {
	Id       int64  `db: "id"`
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (u User) Save() {
	dbmap := initDb()
	defer dbmap.Db.Close()
	dbmap.Insert(u)
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
