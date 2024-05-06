package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaliebenberg/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	filter := db.Map{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return ErrResourceNotFound("hotel")
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotels, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("hotels")
	}
	return c.JSON(hotels)
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var pagination db.Pagination
	if err := c.QueryParser(&pagination); err != nil {
		return ErrBadRequest()
	}

	// __MongoDB specific implementation__
	// opts := options.FindOptions{}
	// opts.SetSkip(int64((page - 1) * limit))
	// opts.SetLimit(int64(limit))

	hotels, err := h.store.Hotel.GetHotels(c.Context(), nil, &pagination)
	if err != nil {
		return ErrResourceNotFound("hotels")
	}
	resp := ResourceResp{
		Data:    hotels,
		Results: len(hotels),
		Page:    int(pagination.Page),
	}
	return c.JSON(resp)
}
