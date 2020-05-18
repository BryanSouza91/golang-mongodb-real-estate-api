package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	"retrck/dataaccessobject"
)

var (
	dao = dataaccessobject.DAO{}
	// Define possible paths/routes; Handle invalid path/route
 	validPath = regexp.MustCompile("^/(nickname|all)/([a-zA-Z0-9]+)$")
)

type Config struct {
	Server   string
	Database string
}

// Find all documents
func findAllHandler(w http.ResponseWriter, r *http.Request) {
	props, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	respondWithJSON(w, http.StatusOK, props)
}

// Find by Nickname
func findByNicknameHandler(w http.ResponseWriter, r *http.Request, nickname string) {
	prop, err := dao.FindOne(nickname)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	respondWithJSON(w, http.StatusOK, prop)
}

// Error Response func
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// Success Response func
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Function to route URL to correct function
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
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("error:", err)
	}
	dao.Server = configuration.Server
	dao.Database = configuration.Database
	dao.Connection()
}

// API routing
func main() {
	http.HandleFunc("/all", findAllHandler)
	http.HandleFunc("/nickname/", makeHandler(findByNicknameHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
