package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"retrck/dataaccessobject"
)

var (
	dao = dataaccessobject.DAO{}
)

type Config struct {
	Server   string
	Database string
}

func viewAllHandler(w http.ResponseWriter, r *http.Request) {
	props, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	respondWithJSON(w, http.StatusOK, props)
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

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	dao.Server = configuration.Server
	dao.Database = configuration.Database
	dao.Connection()
}

func main() {
	http.HandleFunc("/", viewAllHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
