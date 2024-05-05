package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lucaliebenberg/hotel-reservation/api"
	"github.com/lucaliebenberg/hotel-reservation/db"
	"github.com/lucaliebenberg/hotel-reservation/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "foo", "baz", false)
	fmt.Println("user -> ", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "baz", "foo", true)
	fmt.Println("admin -> ", api.CreateTokenFromUser(admin))
	hotel := fixtures.AddHotel(store, "some hotel", "bermuda", 5, nil)
	fmt.Println("hotel -> ", hotel)
	room := fixtures.AddRoom(store, "large", true, 99.99, hotel.ID)
	fmt.Println("room -> ", room)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fmt.Println("booking -> ", booking)
}
