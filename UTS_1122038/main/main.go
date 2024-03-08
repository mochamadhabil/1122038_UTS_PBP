package main

import (
	"UTS_1122038/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/allRooms", controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/detailAllRooms", controller.GetDetailRooms).Methods("GET")
	router.HandleFunc("/v1/insertProduct", controller.InsertNewProducts).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to Port 8888")
	log.Println("Connected to Port 8888")
	log.Fatal(http.ListenAndServe(":8080", router))
}
