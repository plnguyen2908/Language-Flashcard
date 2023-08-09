package main

import (
	"service/handler"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	r := gin.Default()
	// Enable CORS middleware
	r.Use(corsMiddleware())
	r.GET("/flash", handler.DictGet("mongodb://mongo:27017", "flashcard"))
	r.GET("/flash/:id", handler.DictGetOne("mongodb://mongo:27017", "flashcard"))
	r.POST("/flash/:id", handler.DictPost("mongodb://mongo:27017", "flashcard"))
	r.PUT("/flash/:id", handler.DictPut("mongodb://mongo:27017", "flashcard"))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// CORS middleware functions
func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS"},
		ExposedHeaders:   []string{"Content-Length"},
		MaxAge:           86400,
		AllowCredentials: true,
	})
}
