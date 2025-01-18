package services

import (
	"database/sql"
	"log"
	"parcel/internal/database"
	"parcel/internal/entities"
	"time"

	_ "modernc.org/sqlite"
)

type Parcel struct {
	Store database.Store
}

// NewParcelService создает сервис для работы с посылками
func NewParcelService(db *sql.DB) *Parcel {
	return &Parcel{Store: database.Store{Db: db}}
}

// Register регистрирует посылку клиента
func (s *Parcel) Register(client int, address string) (entities.Parcel, error) {
	parcel := entities.Parcel{
		Client:    client,
		Status:    entities.ParcelStatusRegistered,
		Address:   address,
		CreatedAt: time.Now().UTC().Format(time.DateTime),
	}

	id, err := s.Store.Add(&parcel)
	if err != nil {
		return parcel, err
	}

	parcel.Number = id

	log.Printf("Новая посылка № %d на адрес %s от клиента с идентификатором %d зарегистрирована\n",
		parcel.Number, parcel.Address, parcel.Client)

	return parcel, nil
}

// NextStatus изменяет статус посылки на следующий в зависимости от предыдущего
func (s *Parcel) NextStatus(number int) error {
	parcel, err := s.Store.Get(number)
	if err != nil {
		return err
	}

	var nextStatus string
	switch parcel.Status {
	case entities.ParcelStatusRegistered:
		nextStatus = entities.ParcelStatusSent
	case entities.ParcelStatusSent:
		nextStatus = entities.ParcelStatusDelivered
	case entities.ParcelStatusDelivered:
		return nil
	}

	log.Printf("У посылки № %d новый статус: %s\n", number, nextStatus)

	return s.Store.SetStatus(number, nextStatus)
}

// ChangeAddress изменяет адрес посылки 
func (s *Parcel) ChangeAddress(number int, address string) error {
	var err error
	if err = s.Store.SetAddress(number, address); err == nil {
		log.Printf("У посылки № %d новый адрес: %s\n", number, address)
	}

	return err
}

// Delete удаляет посылку  
func (s *Parcel) Delete(number int) error {
	var err error
	if err = s.Store.Delete(number); err == nil {
		log.Printf("Посылка № %d удалена\n", number)
	}

	return err
}
