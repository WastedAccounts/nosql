package mongodbplayground

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"nosql/controllers/validateinput"
	"nosql/datasources"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func queryMongo(dbName string) {

	fmt.Println("Nosql: What are we looking for?")
	fmt.Println("Nosql: Doesn't matter, hit enter and I will get everything")
	// fmt.Println("Nosql: Format - {\"descriptionwords\":\"datawords\"}")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			fmt.Println("Nosql: Later nerd")
			os.Exit(12)
			// }
			// marshalBson(scanner.Text())
			// break
		}
		validateinput.MarshalBson("{}")
		queryAll(dbName, &validateinput.Bdoc)
	}
}

func queryAll(dbName string, value interface{}) {
	conn := datasources.MongoConnPool.Database(dbName).Collection(dbName)
	findResult, err := conn.Find(context.TODO(), value, options.Find()) //InsertOne(context.TODO(), bsondoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(findResult)
	// printing the result of query.
	fmt.Println("Query Result")
	for findResult.Next(context.Background()) {
		fmt.Println(findResult.Current.String())
	}
}
