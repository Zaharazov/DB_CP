package database

import (
	"DB_CP/internal/domain/models"
	"database/sql"
	"fmt"
)

// FetchGymsFromDB - функция для получения данных о спортзалах из базы данных
func FetchGymsFromDB() ([]models.Gym, error) {
	connStr := "user=postgres password=dominator123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения данных
	query := "SELECT gym_id, name, address, description FROM gyms"

	// Выполняем запрос
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читаем данные
	var gyms []models.Gym
	for rows.Next() {
		var gym models.Gym
		if err := rows.Scan(&gym.ID, &gym.Name, &gym.Address, &gym.Description); err != nil {
			return nil, err
		}
		gyms = append(gyms, gym)
	}

	return gyms, nil
}

func FetchGymByID(id string) (models.Gym, error) {
	connStr := "user=postgres password=dominator123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return models.Gym{}, err
	}
	defer db.Close()

	// SQL-запрос для получения данных о зале по ID
	query := "SELECT gym_id, name, address, description FROM gyms WHERE gym_id = $1"

	// Используем QueryRow, так как ожидаем только одну строку результата
	row := db.QueryRow(query, id)

	// Читаем данные из результата
	var gym models.Gym
	err = row.Scan(&gym.ID, &gym.Name, &gym.Address, &gym.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если зал с указанным ID не найден
			return models.Gym{}, fmt.Errorf("gym with ID %s not found", id)
		}
		// Другая ошибка
		return models.Gym{}, err
	}

	// Возвращаем найденный зал
	return gym, nil
}

// func FetchEquipmentForGym(id string) ([]models.Equipment, error) {
// 	// Логика получения тренажеров для зала по gymID
// 	// Верните срез Equipment и ошибку
// }

// func FetchGroupClassesForGym(id string) ([]models.GroupClass, error) {
// 	// Логика получения групповых тренировок для зала
// 	// Верните срез GroupClass и ошибку
// }

func FetchTrainersForGym(id string) ([]models.Coach, error) {
	connStr := "user=postgres password=dominator123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения данных
	query := "SELECT name, description FROM coaches WHERE gym_id = $1"

	// Выполняем запрос
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читаем данные
	var coaches []models.Coach
	for rows.Next() {
		var coach models.Coach
		if err := rows.Scan(&coach.Name, &coach.Description); err != nil {
			return nil, err
		}
		coaches = append(coaches, coach)
	}

	return coaches, nil
}
