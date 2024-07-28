package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gi := gin.Default()
	gi.Use(handleCors())
	gi.GET("/", handleRoot)
	//addr := os.Getenv("HOST_ADDR")
	//addr = "0.0.0.0:9000"
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
