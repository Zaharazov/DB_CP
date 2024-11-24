package services

import (
	"DB_CP/internal/database"
	"DB_CP/internal/domain/models"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func GetGymsPages(w http.ResponseWriter, r *http.Request) {
	// Получаем данные о залах
	gyms, err := database.FetchGymsFromDB()
	if err != nil {
		http.Error(w, "Failed to fetch gyms", http.StatusInternalServerError)
		return
	}

	// Парсим HTML-шаблон
	tmpl, err := template.ParseFiles("./frontend/pages/gyms_page.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Рендерим HTML с данными
	if err := tmpl.Execute(w, gyms); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func GetGymsPageById(w http.ResponseWriter, r *http.Request) {
	// Извлечение gym_id из URL
	vars := mux.Vars(r)
	gymID := vars["gym_id"]

	// Получаем данные о конкретном зале
	gym, err := database.FetchGymByID(gymID)
	if err != nil {
		http.Error(w, "Failed to fetch gym details", http.StatusInternalServerError)
		return
	}

	// Получаем данные о тренажерах в этом зале
	// equipment, err := database.FetchEquipmentForGym(gymID)
	// if err != nil {
	// 	http.Error(w, "Failed to fetch gym equipment", http.StatusInternalServerError)
	// 	return
	// }

	// Получаем данные о групповых тренировках
	// groupClasses, err := database.FetchGroupClassesForGym(gymID)
	// if err != nil {
	// 	http.Error(w, "Failed to fetch group classes", http.StatusInternalServerError)
	// 	return
	// }

	// Получаем данные о тренерах
	coaches, err := database.FetchTrainersForGym(gymID)
	if err != nil {
		http.Error(w, "Failed to fetch coaches", http.StatusInternalServerError)
		return
	}

	// Подготавливаем данные для шаблона
	data := struct {
		Gym models.Gym
		// Equipment    []models.Equipment
		// GroupClasses []models.GroupClass
		Coaches []models.Coach
	}{
		Gym: gym,
		// Equipment:    equipment,
		// GroupClasses: groupClasses,
		Coaches: coaches,
	}

	// Загружаем HTML-шаблон
	tmpl, err := template.ParseFiles("./frontend/pages/gym_page.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Рендерим HTML с данными
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
