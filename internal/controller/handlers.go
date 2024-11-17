package controller

import (
	"fmt"
	"net/http"
	"os"
)

func readHTML(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func readCSS(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := readHTML("./frontend/pages/home_page/home_page.html")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	cssContent, err := readCSS("./frontend/pages/home_page/styles.css")
	if err != nil {
		http.Error(w, "Ошибка при чтении CSS файла", http.StatusInternalServerError)
		return
	}
	htmlContent = fmt.Sprintf("<style>%s</style>%s", cssContent, htmlContent)
	fmt.Fprint(w, htmlContent)
}

func GymsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Gyms Page</h1>")
}

func GymPassesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Gym Passes Page</h1>")
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Admin Page</h1>")
}
