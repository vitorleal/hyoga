package server

import(
    "log"
    "os"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/fiscaluno/hyoga/server/routes"
    _ "github.com/fiscaluno/hyoga/database/migrations"

)

func Start() {
    port := os.Getenv("PORT")
    router := routes.GetRouter()

    origins := handlers.AllowedOrigins([]string{"*"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

    if err := http.ListenAndServe(":" + port, handlers.CORS(origins, methods)(router));
    err != nil {
        log.Fatal(err)
    }
}
