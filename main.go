package main

import (
	"fmt"
	"log"
	"movie-reservation-system/internal/config"
	"movie-reservation-system/internal/db"
	"movie-reservation-system/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start project")

	r := gin.Default()

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error config: %v", err)
	}
	log.Printf("%+v", config)
	db.InitDB(config)

	routes.AuthRoutes(r)

	r.Run(":8080")
}
