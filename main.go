package main

import (
	"fmt"
	"log"
	"movie-reservation-system/internal/config"
	"movie-reservation-system/internal/db"
	"movie-reservation-system/internal/routes"
	"movie-reservation-system/internal/tmdb"

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

	movie, err := tmdb.GetPopularMovies()
	if err != nil {
		log.Fatalf("error getting movies: %v", err)
	}
	for i := 0; i < len(movie); i++ {
		fmt.Printf("%v: %d\n", movie[i].Title, movie[i].ID)
	}
	//fmt.Printf("movies: %+v", movie)

	movieInfo, err := tmdb.GetMovieInfo(1297763)
	fmt.Printf("%+v", movieInfo)
	r.Run(":8080")

}
