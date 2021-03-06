package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func SetupServer() {
	router := mux.NewRouter()
	routerConfig(router)
	router.Use(JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
