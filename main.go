package main

import "Stock_broker_application/src/app/authentication/router"

func main() {
	router := router.SetupRouter()
	router.Run(":8080")
}
