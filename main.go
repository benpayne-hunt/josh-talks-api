package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CNX = Connection()

func Connection() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Print(err)
		}
	}()

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

var buttonsCollection = CNX.Database("joshTalks").Collection("buttons")

func getButtons(c *gin.Context) {
	cursor, err := buttonsCollection.Find(, bson.M{})
	if err != nil {
		fmt.Print(err)
	}

	var buttons []bson.M
	if err = cursor.All(c, &buttons); err != nil {
		fmt.Print(err)
	}

	c.IndentedJSON(http.StatusOK, buttons)
}

func buttonById(c *gin.Context) {
	button, err := getButtonById(c)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Button not found."})
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, button)
}

func getButtonById(c *gin.Context) (*bson.M, error) {
	var button bson.M
	err := buttonsCollection.FindOne(c, bson.M{}).Decode(&button)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// do nothing
		}
		panic(err)
	}

	return &button, nil
}

// func createButton(c *gin.Context) {
// 	var newButton models.Button

// 	if err := c.BindJSON(&newButton); err != nil {
// 		return
// 	}

// 	buttons = append(buttons, newButton)
// 	c.JSON(http.StatusCreated, newButton)
// }

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.GET("/buttons", getButtons)

	router.GET("/buttons/:id", buttonById)

	// router.POST("/new-button", createButton)

	router.Run("localhost:6000")
}
