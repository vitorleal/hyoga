package mu

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fiscaluno/pandorabox"
)

// AuthMiddleware auth all requests
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authURL := pandorabox.GetOSEnvironment("AUTH_URL", "http://localhost:5001/users/token/")

		if r.Header.Get("X-Client-ID") == "" {
			log.Println("X-Client-ID is NULL")
			msg := pandorabox.Message{
				Content: "Security",
				Status:  http.StatusText(http.StatusUnauthorized),
				Body:    nil,
			}
			pandorabox.RespondWithJSON(w, http.StatusUnauthorized, msg)
			return
		}

		if r.Header.Get("X-User-Token") == "" {
			log.Println("X-User-Token is NULL")
			msg := pandorabox.Message{
				Content: "Security",
				Status:  http.StatusText(http.StatusForbidden),
				Body:    nil,
			}
			pandorabox.RespondWithJSON(w, http.StatusForbidden, msg)
			return
		}

		type resp struct {
			Valid  bool   `json:"valid"`
			Active bool   `json:"active"`
			Note   string `json:"note"`
		}

		var response resp

		token := r.Header.Get("X-User-Token")

		client := &http.Client{}

		req, err := http.NewRequest("GET", authURL+token, nil)

		req.Header.Add("X-Client-ID", r.Header.Get("X-Client-ID"))

		res, err := client.Do(req)
		if err != nil {
			log.Println("Bad Requet, ERROR: " + err.Error())
			msg := pandorabox.Message{
				Content: "ERROR",
				Status:  http.StatusText(http.StatusInternalServerError),
				Body:    nil,
			}
			pandorabox.RespondWithJSON(w, http.StatusInternalServerError, msg)
			return
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			msg := pandorabox.Message{
				Content: "ERROR",
				Status:  http.StatusText(res.StatusCode),
				Body:    nil,
			}
			pandorabox.RespondWithJSON(w, res.StatusCode, msg)
			return
		}
		json.NewDecoder(res.Body).Decode(&response)

		if !response.Active || !response.Valid {
			msg := pandorabox.Message{
				Content: "ERROR",
				Status:  http.StatusText(http.StatusForbidden),
				Body:    nil,
			}
			pandorabox.RespondWithJSON(w, http.StatusForbidden, msg)
			return
		}

		log.Println("X-Client-ID is " + r.Header.Get("X-Client-ID"))

		next.ServeHTTP(w, r)
	})
}
