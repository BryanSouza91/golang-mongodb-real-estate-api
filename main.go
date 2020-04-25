package main

import (
	"encoding/json"
	"log"
	"regexp"
	"net/http"

	"retrck/config"
	"retrck/dataaccessobject"
)

var conf = config.Config{}
var dao = dataaccessobject.DAO{}
var apn string

func findAllHandler(w http.ResponseWriter, r *http.Request) {
	props, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	respondWithJSON(w, http.StatusOK, props)
}

func findByNicknameHandler(w http.ResponseWriter, r *http.Request, nickname string) {
	prop, err := dao.FindOne(nickname)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	respondWithJSON(w, http.StatusOK, prop)
}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

var validPath = regexp.MustCompile("^/(nickname|all)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.EscapedPath())
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	conf.Read()

	dao.Server = conf.Server
	dao.Database = conf.Database
	dao.Connection()
}

func main() {
	http.HandleFunc("/all", findAllHandler)
	http.HandleFunc("/nickname/", makeHandler(findByNicknameHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
 