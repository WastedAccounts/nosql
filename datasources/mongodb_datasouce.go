package datasources

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoConnPool *mongo.Client
	MongoReady    bool
)

func SetupMongodb() bool {
	initMongoPool()
	pingMongoPool()
	log.Println("MongoDB initialization complete")
	return MongoReady
}

func initMongoPool() *mongo.Client {
	// Get environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
	MONGODB_SVR := os.Getenv("MONGODB_IP")
	MONGODB_PORT := os.Getenv("MONGODB_PORT")
	MONGO_DB_URI := "mongodb://" + MONGODB_SVR + ":" + MONGODB_PORT + "/?maxPoolSize=20&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoConnPool, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_DB_URI))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to mongodb database")
	}

	// pinged := mongoConnPool.Ping(context.TODO(), readpref.Primary())
	// if err != nil {
	// 	log.Println("PINGED .. Maybe")
	// }
	// log.Println(pinged)

	MongoConnPool = mongoConnPool
	return MongoConnPool
}

func pingMongoPool() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := MongoConnPool.Ping(ctx, readpref.Primary())
	if err != nil {
		MongoReady = false
		log.Println("FAILED Mongodb Database PING")
		log.Fatal("ERROR: ", err)
	} else {
		MongoReady = true
		log.Println("Mongodb database ping succesfull")
	}

}
