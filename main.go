package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	os.Setenv("HOST_ADDR", ":8080")
	gi := gin.Default()
	gi.Use(handleCors())
	gi.GET("/", handleRoot)
	addr := os.Getenv("HOST_ADDR")
	panic(http.ListenAndServe(addr, gi))
}

func handleCors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Add("Access-Control-Max-Age", "86400")
		context.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		context.Next()

	}

}

func handleRoot(context *gin.Context) {
	context.Writer.Write([]byte("Helo Gorik"))
}
