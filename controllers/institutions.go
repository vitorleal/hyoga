package controllers

import(
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/fiscaluno/hyoga/models/institution"
    "github.com/gorilla/mux"
)

func All(w http.ResponseWriter, r *http.Request) {
    institutions := institution.All()

    response, _ := json.Marshal(institutions)
    w.Write([]byte(response))
}

func ById(w http.ResponseWriter, r *http.Request)  {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    response, _ := json.Marshal(institution.Find(id))
    w.Write([]byte(response))
}

func New(w http.ResponseWriter, r *http.Request) {
    new_institution := map[string] interface{} {
        "Name": r.FormValue("name"),
        "Address": r.FormValue("address"),
        "Cnpj": r.FormValue("cnpj"),
    	"Email": r.FormValue("email"),
    	"Website": r.FormValue("website"),
    	"Phone": r.FormValue("phone"),
    	"ImageUri": r.FormValue("image"),
    }
    institution := institution.Create(new_institution)
    institution.Save()

    w.Write([]byte("Success"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    institution.Delete(id)
}
