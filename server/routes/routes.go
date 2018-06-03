package routes

import(
    "github.com/gorilla/mux"
    "github.com/fiscaluno/hyoga/controllers"
)

func GetRouter() (router *mux.Router) {
    router = mux.NewRouter()
    loadRoutes(router)
    return
}

func loadRoutes(router *mux.Router) {
    routes := router.PathPrefix("/").Subrouter()
    routes.HandleFunc("/", controllers.All).Methods("GET")
    routes.HandleFunc("/{id:[0-9]+}", controllers.ById).Methods("GET")
    routes.HandleFunc("/", controllers.New).Methods("POST")
}
