package server

import(
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/fiscaluno/hyoga/config"
    "github.com/fiscaluno/hyoga/server/routes"
    _ "github.com/fiscaluno/hyoga/database/migrations"
)

func Start() {
    router := routes.GetRouter()

    origins := handlers.AllowedOrigins([]string{"*"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

    if err := http.ListenAndServe(":" + config.DEV_SERVER_PORT, handlers.CORS(origins, methods)(router));

    err != nil {
        log.Fatal(err)
    }
}
