package router

import(
	"github.com/gorilla/mux"
)

func Router() *mux.Router{

	router := mux.NewRouter()
	// router.handleFunc("/api/task", middle)
	return router
}