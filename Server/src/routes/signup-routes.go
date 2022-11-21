package routes

import(
	"github.com/gorilla/mux"
	"Server/src/controllers"
)

var RegisterSignUpRoutes = func(router *mux.Router) {
	router.HandleFunc("/signup/", controllers.StoreSignupDetails).Methods("POST")
}