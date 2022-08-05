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

// Init -- app initialization tasks. Exposed and called from nosql.go
func Init() {
	// Set Timezone
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
