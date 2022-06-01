package main

import "shopping/common/initialize"

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Router()
}
