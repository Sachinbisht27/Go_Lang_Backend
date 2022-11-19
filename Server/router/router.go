package router

import(
	"github.com/gorilla/mux"
	"Server/middleware"
	// "net/http"
)

func Router() *mux.Router{

	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.StoreSignupDetails).Methods("POST", "OPTIONS")
	return router
}