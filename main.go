package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lucaliebenberg/hotel-reservation/api"
	"github.com/lucaliebenberg/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userCol = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	// handler initilization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	listenAddr := flag.String("listenAddr", ":4000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	if err := app.Listen(*listenAddr); err != nil {
		panic(err)
	}

}
