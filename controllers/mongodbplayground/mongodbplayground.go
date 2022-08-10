package mongodbplayground

import (
	"context"
	"fmt"
	"log"
	"nosql/controllers/validateinput"
	"nosql/datasources"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func MongoPlayground() {
	fmt.Println("Nosql: Welcome to MongoDB land")
	mongoDatabase := chooseDatabase()
	prompt := "Nosql: What do you want to do? \n Nosql: 1 - Insert, 2 - Query, 3 - Update, 4 - Delete, back - to go back"
	valid := map[string]bool{"1": true, "2": true, "3": true, "4": true, "back": true}
	decision := validateinput.PromptForInput(prompt, valid)
	if decision != "back" {
		choosePath(decision, mongoDatabase)
	}
}

func choosePath(s string, dbName string) {
	valid := map[string]bool{"1": true, "2": true}
	if valid[s] {
		if s == "1" {
			fmt.Println("Nosql: Insertion \U0001f60f")
			insertToMongo(dbName)
		}
		if s == "2" {
			fmt.Println("Nosql: Read it, dork")
			queryMongo(dbName)
		}
		if s == "3" {
			fmt.Println("Nosql: Update prod, it's totally safe")
		}
		if s == "4" {
			fmt.Println("Nosql: Deleting is fun!")
		}
	}
}

func chooseDatabase() string {
	validMap := map[string]bool{}
	dbMap := map[string]string{}
	promptStr := ""
	conn, err := datasources.MongoConnPool.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nosql: Choose a database to muck around in")
	count := len(conn)
	for i, s := range conn {
		if i+1 == count { //|| i/4 == 2
			ii := strconv.Itoa(i + 1)
			validMap[ii] = true
			validMap["back"] = true
			dbMap[ii] = s
			promptStr = promptStr + fmt.Sprint(i+1, " - ", s, "\n")
			// prompt = prompt + fmt.Sprint(+1, " - ", s)
			// fmt.Print(i+1, " - ", s)
			// fmt.Print("\n")
		} else {
			ii := strconv.Itoa(i + 1)
			validMap[ii] = true
			dbMap[ii] = s
			promptStr = promptStr + fmt.Sprint(i+1, " - ", s, ", ")
		}
	}
	prompt := promptStr
	valid := validMap
	decision := validateinput.PromptForInput(prompt, valid)

	fmt.Println("decision :", decision)
	fmt.Println("decisionMap :", dbMap[decision])
	return dbMap[decision]
}

///// Web stuff if I want to make it a web app

// import (
// 	"net/http"
// 	"regexp"
// )

// type mongodbPlaygroundController struct {
// 	mongodbPlaygroundPattern *regexp.Regexp
// }

// // entry point from front.go
// func NewMongodbPlaygroundController() *mongodbPlaygroundController {
// 	return &mongodbPlaygroundController{
// 		mongodbPlaygroundPattern: regexp.MustCompile(`^/mongodbplayground(\d+)/?`),
// 	}
// }

// // ServeHTTP - Entry point from front.go
// func (rpgc mongodbPlaygroundController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "/mongodbplayground.html")
// 	//pageloadRedisPlayground(w, r)
// }

// // pageLoadWorkouts - switch to workouts template do work there
// func pageloadMongodbPlayground(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "/mongodbplayground.html")
// }
