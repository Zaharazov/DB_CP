package app

import (
	"DB_CP/internal/database"
	"DB_CP/internal/presentation/routers"
	"DB_CP/internal/server/configs"
	"fmt"
	"log"
	"net/http"
)

func Run() {

	r := routers.NewRouter()

	log.Printf("Server started!")

	database.ConnectToPostrges()

	if err := http.ListenAndServe(configs.Port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
