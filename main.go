package main

import (
	"net/http"

	"github.com/benpayne-hunt/josh-talks-api/configs"
	"github.com/benpayne-hunt/josh-talks-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	// router.HandleFunc("/fetch-buttons", makeHttpHandleFunc(server.handleRequest))

	// log.Println("JSON API server is running on port: ", server.listenAddr)

	// configs.ConnectDB()

	// routes.Route(router)

	router.Run("localhost:6000")
	// http.ListenAndServe(server.listenAddr, router)
}
