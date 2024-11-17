package app

import (
	"DB_CP/internal/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {

	r := mux.NewRouter()
	r.HandleFunc("/", controller.HomeHandler)
	r.HandleFunc("/gyms", controller.GymsHandler)
	r.HandleFunc("/gym_passes", controller.GymPassesHandler)
	r.HandleFunc("/admin", controller.AdminHandler)

	log.Printf("Server started!")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
