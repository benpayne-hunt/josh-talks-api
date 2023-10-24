package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	// router.HandleFunc("/fetch-services", makeHttpHandleFunc(server.handleRequest))

	// log.Println("JSON API server is running on port: ", server.listenAddr)

	// configs.ConnectDB()

	// routes.Route(router)

	router.Run("localhost:6000")
	// http.ListenAndServe(server.listenAddr, router)
}
