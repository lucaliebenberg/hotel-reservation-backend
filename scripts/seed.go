package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lucaliebenberg/hotel-reservation/api"
	"github.com/lucaliebenberg/hotel-reservation/db"
	"github.com/lucaliebenberg/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	roomStore    db.RoomStore
	hotelStore   db.HotelStore
	userStore    db.UserStore
	bookingStore db.BookingStore
	ctx          = context.Background()
)

func seedUser(isAdmin bool, fName, lName, email, psswd string) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     email,
		FirstName: fName,
		LastName:  lName,
		Password:  psswd,
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = isAdmin
	insertedUser, err := userStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))
	return insertedUser
}

func seedBooking(userID, roomID primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:   userID,
		RoomID:   roomID,
		FromDate: from,
		TillDate: till,
	}
	resp, err := bookingStore.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("booking: ", resp.ID)
	return nil
}

func seedHotel(name string, location string, rating int) *types.Hotel {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func seedRoom(size string, seaside bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		SeaSide: seaside,
		Price:   price,
		HotelID: hotelID,
	}
	insertedRoom, err := roomStore.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func main() {
	foo := seedUser(false, "foo", "baz", "foobaz@twitter.com", "notsecurepsswd")
	seedUser(true, "baz", "foo", "bazfoo@twitter.com", "notsecurepsswd")
	seedHotel("Twelve Apostles", "South Africa", 3)
	seedHotel("Silo Hotel", "South Africa", 4)
	hotel := seedHotel("Weed Hotel", "Netherlands", 5)
	seedRoom("small", true, 89.99, hotel.ID)
	seedRoom("medium", true, 189.99, hotel.ID)
	room := seedRoom("large", false, 289.99, hotel.ID)
	seedBooking(foo.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2))

}

func init() {
	var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	bookingStore = db.NewMongoBookingStore(client)
}
