package main

import (
	"log"
	"net/http"
	"parcel/internal"
	"parcel/internal/database"
	"parcel/internal/handlers"
	"parcel/internal/services"
)

func main() {
	db, err := database.NewParcelStore()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	service := services.NewParcelService(db)
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("ui")))

	mux.HandleFunc("POST /parcels", handlers.RegisterParcel(service))
	mux.HandleFunc("GET /parcels", handlers.GetParcels(service))
	mux.HandleFunc("PUT /parcels/{number}/next-status", handlers.UpdateStatus(service))
	mux.HandleFunc("PUT /parcels/{number}/change-address", handlers.UpdateAddress(service))
	mux.HandleFunc("DELETE /parcels/{number}", handlers.DeleteParcel(service))

	log.Println("Service is running ...")
	if err := http.ListenAndServe(internal.Port, mux); err != nil {
		log.Fatal(err.Error())
	}
}
