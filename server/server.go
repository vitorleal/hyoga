package server

import(
    "log"
    "net/http"
    "github.com/fiscaluno/hyoga/config"
    "github.com/fiscaluno/hyoga/server/routes"
    _ "github.com/fiscaluno/hyoga/database/migrations"
    
)

func Start() {
    port := config.DEV_SERVER_PORT
    router := routes.GetRouter()

    if err := http.ListenAndServe(":" + port, router); err != nil {
        log.Fatal(err)
    }
}
