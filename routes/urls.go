package routes

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gregoflash05/firstapi/views"
	// "github.com/joho/godotenv"
)

func Startsever() {
	App := mux.NewRouter()
	//err := godotenv.Load(".env")

	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default port if not specified
	}
	fmt.Println("Starting Server at", port)
	defer http.ListenAndServe(port, App)
	fmt.Println("Server Started")

	App.HandleFunc("/", views.Helloworld).Methods("GET")
	App.HandleFunc("/contact", views.Contact).Methods("POST")
	App.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))
	App.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
	App.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("static/images/"))))

}
