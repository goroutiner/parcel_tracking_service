package database

import (
	"errors"
	"log"
	"parcel/internal"
	"parcel/internal/entities"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	Db *sqlx.DB
}

// NewParcelStore создает таблицу для хранения посылок
func NewParcelStore() (*sqlx.DB, error) {
	var err error
	entities.Db, err = sqlx.Open("pgx", internal.PsqlUrl)
	if err != nil {
		return nil, err
	}

	var exists bool
	err = entities.Db.Get(&exists, `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`, internal.TableName)
	if err != nil {
		return nil, err
	}

	if !exists {
		entities.Db.MustExec(`
		CREATE TABLE parcel (
			number SERIAL PRIMARY KEY,
			client INTEGER NOT NULL DEFAULT 0,
			status TEXT NOT NULL DEFAULT '',
			address TEXT NOT NULL DEFAULT '',
			created_at TEXT NOT NULL DEFAULT ''
		);

		CREATE INDEX parcel_indx ON parcel (client, created_at);
		`)
		log.Println("Creating a database ...")
	}

	return entities.Db, err
}

// Add добавляет посылку в БД
func (s *Store) Add(p *entities.Parcel) (int, error) {
	var id int
	row := s.Db.QueryRow("INSERT INTO parcel (client, status, address, created_at) VALUES ($1, $2, $3, $4) RETURNING number",
		p.Client, p.Status, p.Address, p.CreatedAt)
	if err := row.Err(); err != nil {
		return 0, err
	}

	row.Scan(&id)
	return id, nil
}

// Get получает параметры посылки по ее номеру из БД
func (s *Store) Get(number int) (entities.Parcel, error) {
	var parcel entities.Parcel

	err := s.Db.Get(&parcel, "SELECT number, client, status, address, created_at FROM parcel WHERE number = $1", number)
	return parcel, err
}

// GetByClient получает список всех посылок клиента по его номеру из БД
func (s *Store) GetByClient(client int) ([]entities.Parcel, error) {
	var parcels []entities.Parcel

	err := s.Db.Select(&parcels, "SELECT number, client, status, address, created_at FROM parcel WHERE client = $1", client)
	return parcels, err
}

// GetParcels получает список посылок всех клиентов из БД
func (s *Store) GetParcels() ([]entities.Parcel, error) {
	var parcels []entities.Parcel

	err := s.Db.Select(&parcels, "SELECT number, client, address, status, created_at FROM parcel")
	return parcels, err
}

// SetStatus устанавливает статус посылки в БД
func (s *Store) SetStatus(number int, status string) error {
	if has := s.CheckParcel(number); !has {
		return errors.New("parcel is not found")
	}

	_, err := s.Db.Exec("UPDATE parcel SET status = $1 WHERE number = $2", status, number)
	return err
}

// SetAddress устанавливает адрес посылки в БД
func (s *Store) SetAddress(number int, address string) error {
	if has := s.CheckParcel(number); !has {
		return errors.New("parcel is not found")
	}

	_, err := s.Db.Exec("UPDATE parcel SET address = $1 WHERE number = $2", address, number)
	return err
}

// Delete удаляет посылку из БД
func (s *Store) Delete(number int) error {
	if has := s.CheckParcel(number); !has {
		return errors.New("parcel is not found")
	}

	_, err := s.Db.Exec("DELETE FROM parcel WHERE number = $1", number)
	return err
}

// CheckParcel проверяет наличие посылки с указанным номером в БД
func (s *Store) CheckParcel(number int) bool {
	var exists bool

	s.Db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM parcel WHERE number = $1);", number)
	return exists
}
