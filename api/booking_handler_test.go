package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/lucaliebenberg/hotel-reservation/db/fixtures"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	user := fixtures.AddUser(db.Store, "foo", "baz", false)
	hotel := fixtures.AddHotel(db.Store, "foo hotel", "å", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 4, hotel.ID)

	from := time.Now()
	till := from.AddDate(0, 0, 5)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
	fmt.Println(booking)
}
