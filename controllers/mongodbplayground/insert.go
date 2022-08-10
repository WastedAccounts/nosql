package mongodbplayground

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"nosql/controllers/validateinput"
	"nosql/datasources"
	"os"
)

func insertToMongo(dbName string) {
	fmt.Println("Nosql: Enter data to insert")
	fmt.Println("Nosql: Format - {\"descriptionwords\":\"datawords\"}")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			fmt.Println("Nosql: Later nerd")
			os.Exit(12)
		}
		validateinput.MarshalBson(scanner.Text())
		break
	}
	insertOne(dbName, &validateinput.Bdoc)

}

// var bdoc interface{}

// func marshalBson(value string) {
// 	// `{"description":"matt"}`
// 	err := bson.UnmarshalExtJSON([]byte(value), true, &bdoc)
// 	if err != nil {
// 		fmt.Println("Nosql: ERROR: badly formatted json data")
// 	}
// }

func insertOne(dbName string, bsondoc interface{}) {
	fmt.Println("starting insert")

	conn := datasources.MongoConnPool.Database(dbName).Collection(dbName)
	insertResult, err := conn.InsertOne(context.TODO(), bsondoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult)
}

func insertMany() {

}
