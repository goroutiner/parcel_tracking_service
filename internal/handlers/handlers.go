package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"parcel/internal/entities"
	"parcel/internal/services"
)

// RegisterParcel - обработчик для регистрации  посылок
func RegisterParcel(service *services.Parcel) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var parcel entities.Parcel
		if err := json.NewDecoder(r.Body).Decode(&parcel); err != nil {
			http.Error(w, fmt.Sprintf("Неверные данные: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		parcel, err := service.Register(parcel.Client, parcel.Address)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при регистрации посылки: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// GetParcels - обработчик для получения списка всех посылок
func GetParcels(service *services.Parcel) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var parcels []entities.Parcel

		parcels, err := service.Store.GetParcels()
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при получении данных: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(parcels)
	}
}

// UpdateStatus - обработчик для обновления статуса посылки на новый в зависимости от предыдущего
func UpdateStatus(service *services.Parcel) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var number int
		if _, err := fmt.Sscanf(r.PathValue("number"), "%d", &number); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обновлении статуса: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		err := service.NextStatus(number)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обновлении статуса: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// UpdateAddress - обработчик для обновления адреса посылки
func UpdateAddress(service *services.Parcel) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var number int
		if _, err := fmt.Sscanf(r.PathValue("number"), "%d", &number); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обновлении статуса: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		var input map[string]string
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, fmt.Sprintf("Неверные данные: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		if err := service.ChangeAddress(number, input["address"]); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обновлении адреса: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteParcel - обработчик для удаления посылки
func DeleteParcel(service *services.Parcel) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var number int
		if _, err := fmt.Sscanf(r.PathValue("number"), "%d", &number); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обновлении статуса: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		if err := service.Delete(number); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при удалении посылки: %v", err), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
