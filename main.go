package main

import (
	"SilicomAPPv0.3/initialize"
)

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Redis()
	//initialize.Elastic()
	initialize.Router()
}
