package controllers

import (
	"net/http"
	"nosql/controllers/mongodbplayground"
	"nosql/controllers/redisplayground"
	"regexp"
)

func RegisterLocalhostControllers() {
	// // register multiple registrations
	http.HandleFunc("/", index)
	mux := http.NewServeMux()
	// //mux.Handler("/", requesthandler)

	// // localhost controller registration
	// //lhc := newLocalhostControllerController()
	// mux.Handler("/", pageloadLocalhost(w, r))
	// //http.Handle("/", *lhc)
	// redisplayground controller registration
	rpgc := redisplayground.NewRedisPlaygroundController()
	mux.Handle("/redisplayground", *rpgc)

	// redisplayground controller registration
	mpgc := mongodbplayground.NewMongodbPlaygroundController()
	mux.Handle("/mongodbplayground", *rpgc)
	// http.Handle("/localhost/redisplayground.html", *rpgc)
	//http.ServeFile(w, r, "localhost/localhost.html")
}

// server localhost
type localhostController struct {
	localhostPattern *regexp.Regexp
}

// entry point from front.go
func newLocalhostControllerController() *localhostController {
	return &localhostController{
		localhostPattern: regexp.MustCompile(`^/(\d+)/?`),
	}
}

// // ServeHTTP - Entry point from front.go
// func (lhc localhostController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	pageloadLocalhost(w, r)
// }

// pageLoadWorkouts - switch to workouts template do work there
func PageloadLocalhost(w http.ResponseWriter, r *http.Request) {
	//RegisterLocalhostControllers()
	http.ServeFile(w, r, "localhost/localhost.html")
}
