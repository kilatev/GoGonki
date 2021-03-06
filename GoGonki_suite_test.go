package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	//"github.com/codegangsta/martini-contrib/binding"
)

var (
	response *httptest.ResponseRecorder
)

func TestGoGonki(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGonki Suite")
}


func Request(method string, route string, handler martini.Handler) {
	m := martini.Classic()
	m.Get(route, handler)
	m.Use(render.Renderer())
	request, _ := http.NewRequest(method, route, nil)
	response = httptest.NewRecorder()
	m.ServeHTTP(response, request)
}