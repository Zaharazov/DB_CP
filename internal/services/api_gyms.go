package services

import (
	"DB_CP/internal/database"
	"encoding/json"
	"net/http"
)

func GetAllGyms(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из базы данных
	gyms, err := database.FetchGymsFromDB()
	if err != nil {
		http.Error(w, "Failed to fetch gyms from database", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки
	w.Header().Set("Content-Type", "application/json")

	// Отправляем данные в формате JSON
	err = json.NewEncoder(w).Encode(gyms)
	if err != nil {
		http.Error(w, "Failed to encode gyms data", http.StatusInternalServerError)
		return
	}
}
