package appinit

import (
	"log"
	"nosql/datasources"
	"os"
)

var (
	redisReady   bool
	mongodbReady bool
)

func Init() {
	os.Setenv("TZ", "America/New_York")

	// initialize redis datasource
	redisReady = datasources.SetupRedis()
	if redisReady == true {
		log.Println("Redis datasource is ready")
	} else {
		log.Println("datasource redis failed")
	}

	// initialize mongodb datasource
	mongodbReady = datasources.SetupMongodb()
	if mongodbReady == true {
		log.Println("Mongodb datasource is ready")
	} else {
		log.Println("datasource mongo failed")
	}

}
