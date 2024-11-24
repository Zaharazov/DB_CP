package app

import (
	"DB_CP/internal/controller"
	"DB_CP/internal/database"
	"DB_CP/internal/presentation/routers"
	"DB_CP/internal/server/configs"
	"fmt"
	"log"
	"net/http"
)

func Run() {

	r := routers.NewRouter()
	r.HandleFunc("/", controller.HomeHandler)
	r.HandleFunc("/gyms", controller.GymsHandler)
	r.HandleFunc("/gym_passes", controller.GymPassesHandler)
	r.HandleFunc("/admin", controller.AdminHandler)

	log.Printf("Server started!")

	database.ConnectToPostrges()

	if err := http.ListenAndServe(configs.Port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
