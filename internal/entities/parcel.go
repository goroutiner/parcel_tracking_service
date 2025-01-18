package entities

import "database/sql"

type Parcel struct {
	Number    int    `json:"number,omitempty"`
	Client    int    `json:"client,omitempty"`
	Status    string `json:"status,omitempty"`
	Address   string `json:"address,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

const (
	ParcelStatusRegistered = "registered"
	ParcelStatusSent       = "sent"
	ParcelStatusDelivered  = "delivered"
)

var (
	Db *sql.DB
)
