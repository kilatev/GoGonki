package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
)

func Auth(params martini.Params) {
	auth.Basic("username", "secretpassword")

}
