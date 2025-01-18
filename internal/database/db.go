package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"parcel/internal/entities"
)

type Store struct {
	Db *sql.DB
}

// NewParcelStore создает таблицу для хранения посылок
func NewParcelStore(fileName string) (*sql.DB, error) {
	var (
		err error
		has bool
	)

	if _, err = os.Stat(fileName); err == nil {
		has = true
	}

	if !has {
		log.Println("Creating a database ...")
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		entities.Db, err = sql.Open("sqlite", fileName)
		if err != nil {
			return nil, err
		}

		_, err = entities.Db.Exec(`CREATE TABLE parcel (
        number INTEGER PRIMARY KEY AUTOINCREMENT,
        client INTEGER NOT NULL DEFAULT 0,
        status TEXT NOT NULL DEFAULT "",
        address TEXT NOT NULL DEFAULT "",
        created_at TEXT NOT NULL DEFAULT "");
		CREATE INDEX parcel_indx ON parcel (client, created_at);
		`)
		if err != nil {
			return nil, err
		}
	}

	entities.Db, err = sql.Open("sqlite", fileName)

	return entities.Db, err
}

// Add добавляет посылку в БД  
func (s *Store) Add(p *entities.Parcel) (int, error) {
	res, err := s.Db.Exec("INSERT INTO parcel (client, status, address, created_at) VALUES (:client, :status, :address, :created_at)",
		sql.Named("client", p.Client),
		sql.Named("status", p.Status),
		sql.Named("address", p.Address),
		sql.Named("created_at", p.CreatedAt))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get получает параметры посылки по ее номеру из БД
func (s *Store) Get(number int) (entities.Parcel, error) {
	var p entities.Parcel

	row := s.Db.QueryRow("SELECT number, client, status, address, created_at FROM parcel WHERE number = :number",
		sql.Named("number", number))
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)

	return p, err
}

// GetByClient получает список всех посылок клиента по его номеру из БД
func (s *Store) GetByClient(client int) ([]entities.Parcel, error) {
	rows, err := s.Db.Query("SELECT number, client, status, address, created_at FROM parcel WHERE client = :client",
		sql.Named("client", client))
	if err != nil {
		return []entities.Parcel{}, err
	}
	defer rows.Close()

	var res []entities.Parcel
	for rows.Next() {
		p := entities.Parcel{}

		err := rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
		if err != nil {
			return []entities.Parcel{}, err
		}

		res = append(res, p)
	}

	if err := rows.Err(); err != nil {
		return []entities.Parcel{}, err
	}

	return res, nil
}

// GetParcels получает список посылок всех клиентов из БД
func (s *Store) GetParcels() ([]entities.Parcel, error) {
	var parcels []entities.Parcel

	rows, err := s.Db.Query("SELECT number, client, address, status, created_at FROM parcel")
	if err != nil {
		return []entities.Parcel{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var parcel entities.Parcel
		if err := rows.Scan(&parcel.Number, &parcel.Client, &parcel.Address, &parcel.Status, &parcel.CreatedAt); err != nil {
			return []entities.Parcel{}, err
		}
		parcels = append(parcels, parcel)
	}
	if err = rows.Err(); err != nil {
		return []entities.Parcel{}, err
	}

	return parcels, nil
}

// SetStatus устанавливает статус посылки в БД
func (s *Store) SetStatus(number int, status string) error {
	if has := s.CheckParcel(number); !has {
		return errors.New("parcel is not found")
	}

	_, err := s.Db.Exec("UPDATE parcel SET status = :status WHERE number = :number",
		sql.Named("status", status),
		sql.Named("number", number))

	return err
}

// SetAddress устанавливает адрес посылки в БД
func (s *Store) SetAddress(number int, address string) error {
	if has := s.CheckParcel(number); !has {
		return errors.New("parcel is not found")
	}

	_, err := s.Db.Exec("UPDATE parcel SET address = :address WHERE number = :number",
		sql.Named("address", address),
		sql.Named("number", number),
	)
	return err
}

// Delete удаляет посылку из БД
func (s *Store) Delete(number int) error {
	if has := s.CheckParcel(number); !has {
		return errors.New("parcel is not found")
	}

	_, err := s.Db.Exec("DELETE FROM parcel WHERE number = :number", sql.Named("number", number))
	return err
}

// CheckParcel проверяет наличие посылки с указанным номером в БД
func (s *Store) CheckParcel(number int) bool {
	has := true
	row := s.Db.QueryRow("SELECT client FROM parcel WHERE number = :number", sql.Named("number", number))
	if err := row.Scan(&has); err == sql.ErrNoRows {
		has = false
	}
	return has
}
