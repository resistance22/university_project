package entity

import (
	"time"

	"github.com/google/uuid"
)

type Consumable struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	UOM       string    `json:"uom"`
	Remaining float64   `json:"remaining"`
}

func NewConsumable(
	title string,
	uom string,
	remaining float64,
) (*Consumable, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Consumable{
		ID:        uuid,
		CreatedAt: time.Now(),
		Title:     title,
		UOM:       uom,
		Remaining: remaining,
	}, nil

}
