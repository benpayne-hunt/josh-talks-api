package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.WriteHeader(status)
	writer.Header().Set("Content Type", "application/json")

	return json.NewEncoder(writer).Encode(value)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHttpHandleFunc(function apiFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if error := function(writer, request); error != nil {
			// handle the error
			WriteJSON(writer, http.StatusBadRequest, ApiError{Error: error.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (server *APIServer) Run() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	// router.HandleFunc("/fetch-buttons", makeHttpHandleFunc(server.handleRequest))

	log.Println("JSON API server is running on port: ", server.listenAddr)

	router.Run()
	// http.ListenAndServe(server.listenAddr, router)
}

func (server *APIServer) handleRequest(writer http.ResponseWriter, request *http.Request) error {
	if request.Method == "GET" {
		return server.handleFetchButtons(writer, request)
	}

	return fmt.Errorf("%s method not allowed - only get methods are supported", request.Method)
}

func (server *APIServer) handleFetchButtons(writer http.ResponseWriter, request *http.Request) error {
	button := NewButton("Newest button")

	return WriteJSON(writer, http.StatusOK, button)
}
