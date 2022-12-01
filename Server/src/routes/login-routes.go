package routes

import(
	"github.com/gorilla/mux"
	controller "Server/src/controllers"
	// "log"
	// "net/http"
)

var RegisterLogInDetailsRoutes = func(router *mux.Router){
	router.HandleFunc("/login",controller.StoreLogInDetails).Methods("POST")
}