package appinit

import (
	"fmt"
	"nosql/datasources"
	"os"
)

var (
	dbReady bool
)

func Init() {
	os.Setenv("TZ", "America/New_York")

	// initialize redis datasource
	dbReady = datasources.SetupRedis()
	if dbReady == true {
		fmt.Println("datasource redis ready")
	} else {
		fmt.Println("datasource redis failed")
	}

	// initialize mongodb datasource
}
