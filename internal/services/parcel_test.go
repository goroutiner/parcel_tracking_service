package services_test

import (
	"database/sql"
	"math/rand"
	"parcel/internal/database"
	"parcel/internal/entities"
	"parcel/internal/services"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	// randSource — это источник псевдослучайных чисел.
	// Для повышения уникальности в качестве seed
	// используется текущее время в unix-формате в виде числа
	randSource = rand.NewSource(time.Now().UnixNano())
	// randRange использует randSource для генерации случайных чисел
	randRange = rand.New(randSource)
)

// getTestParcel возвращает тестовую посылку
func getTestParcel() entities.Parcel {
	return entities.Parcel{
		Client:    1000,
		Status:    entities.ParcelStatusRegistered,
		Address:   "test",
		CreatedAt: time.Now().UTC().Format(time.DateTime),
	}
}

// TestAddGetDelete проверяет добавление, получение и удаление посылки
func TestAddGetDelete(t *testing.T) {
	// prepare
	db, err := database.NewParcelStore()
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	service := services.NewParcelService(db)
	parcel := getTestParcel()

	// add
	parcel.Number, err = service.Store.Add(&parcel)

	require.NoError(t, err)
	require.NotEmpty(t, parcel.Number)

	// get
	stored, err := service.Store.Get(parcel.Number)

	require.NoError(t, err)
	require.Equal(t, parcel, stored)

	// delete
	err = service.Store.Delete(parcel.Number)
	require.NoError(t, err)

	_, err = service.Store.Get(parcel.Number)
	require.Equal(t, sql.ErrNoRows, err)
}

// TestSetAddress проверяет обновление адреса
func TestSetAddress(t *testing.T) {
	// prepare
	db, err := database.NewParcelStore()
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	service := services.NewParcelService(db)
	parcel := getTestParcel()

	// add
	parcel.Number, err = service.Store.Add(&parcel)

	require.NoError(t, err)
	require.NotEmpty(t, parcel.Number)

	// set address
	newAddress := "new test address"
	err = service.Store.SetAddress(parcel.Number, newAddress)

	require.NoError(t, err)

	// check
	stored, err := service.Store.Get(parcel.Number)

	require.NoError(t, err)
	require.Equal(t, newAddress, stored.Address)
}

// TestSetStatus проверяет обновление статуса
func TestSetStatus(t *testing.T) {
	// prepare
	db, err := database.NewParcelStore()
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	service := services.NewParcelService(db)
	parcel := getTestParcel()

	// add
	parcel.Number, err = service.Store.Add(&parcel)

	require.NoError(t, err)
	require.NotEmpty(t, parcel.Number)

	// set status
	err = service.Store.SetStatus(parcel.Number, entities.ParcelStatusSent)
	
	require.NoError(t, err)

	// check
	stored, err := service.Store.Get(parcel.Number)

	require.NoError(t, err)
	require.Equal(t, entities.ParcelStatusSent, stored.Status)
}

// TestGetByClient проверяет получение посылок по идентификатору клиента
func TestGetByClient(t *testing.T) {
	// prepare
	db, err := database.NewParcelStore()
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	service := services.NewParcelService(db)
	parcels := []entities.Parcel{
		getTestParcel(),
		getTestParcel(),
		getTestParcel(),
	}
	parcelMap := map[int]entities.Parcel{}

	// задаём всем посылкам одного клиента
	client := randRange.Intn(10_000_000)
	parcels[0].Client = client
	parcels[1].Client = client
	parcels[2].Client = client

	// add
	for i := 0; i < len(parcels); i++ {
		id, err := service.Store.Add(&parcels[i])

		require.NoError(t, err)
		require.NotEmpty(t, id)

		parcels[i].Number = id
		parcelMap[id] = parcels[i]
	}

	// get by client
	storedParcels, err := service.Store.GetByClient(client)

	require.NoError(t, err)
	require.Len(t, storedParcels, len(parcels))

	// check
	for _, parcel := range storedParcels {
		expectedParcel, ok := parcelMap[parcel.Number]

		require.True(t, ok)
		require.Equal(t, expectedParcel, parcel)
	}
}
