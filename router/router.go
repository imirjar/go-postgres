package router

import (
	"github.com/imirjar/go-postgres/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router = mux.Router()

	router.HandleFunc("/api/stock/{id}", 		middleware.GetStock   ).Method("GET", 		"OPTIONS")
	router.HandleFunc("/api/stock",				middleware.GetAllStock).Method("GET", 		"OPTIONS")
	router.HandleFunc("/api/newstock",	 		middleware.CreateStock).Method("POST",		"OPTIONS")
	router.HandleFunc("/api/stock/{id}", 		middleware.UpdateStock).Method("PUT", 		"OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", 	middleware.DeleteStock).Method("DELETE", 	"OPTIONS")
}