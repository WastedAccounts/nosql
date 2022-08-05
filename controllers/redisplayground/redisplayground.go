package redisplayground

import (
	"net/http"
	"regexp"
)

type redisPlaygroundController struct {
	redisPlaygroundPattern *regexp.Regexp
}

// entry point from front.go
func NewRedisPlaygroundController() *redisPlaygroundController {
	return &redisPlaygroundController{
		redisPlaygroundPattern: regexp.MustCompile(`^/redisplayground.html(\d+)/?`),
	}
}

// ServeHTTP - Entry point from front.go
func (rpgc redisPlaygroundController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/redisplayground.html")
	//pageloadRedisPlayground(w, r)
}

// pageLoadWorkouts - switch to workouts template do work there
func pageloadRedisPlayground(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/redisplayground.html")
}
