package main

import (
	"database/sql"
//	"encoding/json"
//	"fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
//	"log"
)


func CreateUser(user User) User {
	dbmap := initDb()
	defer dbmap.Db.Close()
    user = User{}
	return user
}