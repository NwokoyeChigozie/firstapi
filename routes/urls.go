package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gregoflash05/firstapi/views"
)

func Startsever() {
	App := mux.NewRouter()
	fmt.Println("Starting Server")
	defer http.ListenAndServe(":1234", App)
	fmt.Println("Server Started")

	App.HandleFunc("/", views.Helloworld).Methods("GET")
	App.HandleFunc("/contact", views.Contact).Methods("POST")
	App.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))
	App.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
	App.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("static/images/"))))

}
