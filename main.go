package main

import (
	"Stock_broker_application/docs"
	"Stock_broker_application/src/app/authentication/router"
	"Stock_broker_application/src/app/authentication/utils"
	"Stock_broker_application/src/app/authentication/utils/db"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your API Title
// @version 1.0
// @description This is a sample API for demonstration purposes.
// @host localhost:8080
// @BasePath /

func main() {
	//Initialize database
	db, err := db.SetupGorm()
	if err != nil {
		log.Fatalf("Error initializing database:%v", err)
	}
	logFilePath := "logs.log"
	utils.InitLogger(logFilePath)
	router := router.SetupRouter(db)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
