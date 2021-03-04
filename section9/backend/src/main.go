package main

import (
	"encoding/json"
    "log"
    "net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time time.Time `json:"time"`
	Hostname string `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn {
		Time: time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
	jsonRes, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonRes)
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    log.Fatal(http.ListenAndServe(":9090", nil))
}