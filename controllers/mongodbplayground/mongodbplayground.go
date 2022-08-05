package mongdbplayground

import (
	"net/http"
	"regexp"
)

type mongodbPlaygroundController struct {
	mongodbPlaygroundPattern *regexp.Regexp
}

// entry point from front.go
func NewMongodbPlaygroundController() *mongodbPlaygroundController {
	return &mongodbPlaygroundController{
		mongodbPlaygroundPattern: regexp.MustCompile(`^/mongodbplayground(\d+)/?`),
	}
}

// ServeHTTP - Entry point from front.go
func (rpgc mongodbPlaygroundController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mongodbplayground.html")
	//pageloadRedisPlayground(w, r)
}

// pageLoadWorkouts - switch to workouts template do work there
func pageloadMongodbPlayground(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mongodbplayground.html")
}
