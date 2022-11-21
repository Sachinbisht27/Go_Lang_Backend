package routes

import(
	"github.com/gorilla/mux"
	controller "Server/src/controllers"
	"log"
	"net/http"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.StoreSignupDetails).Methods("POST")


	log.Fatal(http.ListenAndServe(":9000",r))
	
}