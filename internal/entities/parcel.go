package entities

import (
	"github.com/jmoiron/sqlx"
)

type Parcel struct {
	Number    int    `json:"number,omitempty" db:"number"`
	Client    int    `json:"client,omitempty" db:"client"`
	Status    string `json:"status,omitempty" db:"status"`
	Address   string `json:"address,omitempty" db:"address"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
}

const (
	ParcelStatusRegistered = "registered"
	ParcelStatusSent       = "sent"
	ParcelStatusDelivered  = "delivered"
)

var (
	Db *sqlx.DB
)
