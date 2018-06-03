package server

import(
    "log"
    "os"
    "net/http"
    "github.com/fiscaluno/hyoga/server/routes"
    _ "github.com/fiscaluno/hyoga/database/migrations"

)

func Start() {
    port := os.Getenv("PORT")
    router := routes.GetRouter()

    if err := http.ListenAndServe(":" + port, router); err != nil {
        log.Fatal(err)
    }
}
