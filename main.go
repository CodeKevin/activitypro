package main

import (
	db "activitypro/database"
)

func main() {

	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8080")
}
