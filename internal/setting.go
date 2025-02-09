package internal

import "os"

const (
	Port    = ":8080"
    TableName = "parcel"
)

var (
	PsqlUrl = os.Getenv("DATABASE_URL")
)
