package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nishgowda/docktor/lib/autoheal"
	"github.com/nishgowda/docktor/lib/heal"
	"github.com/nishgowda/docktor/lib/healthcheck"
	"github.com/nishgowda/docktor/lib/scan"
	"github.com/nishgowda/docktor/lib/suggestions"
)

func hcheck(w http.ResponseWriter, req *http.Request) {
	container, _ := req.URL.Query()["container"]
	msg, err := healthcheck.PerformHealthCheck(container)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func aheal(w http.ResponseWriter, req *http.Request) {
	container, _ := req.URL.Query()["container"]
	msg, err := autoheal.AutoHeal(container)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func hheal(w http.ResponseWriter, req *http.Request) {
	container, _ := req.URL.Query()["container"]
	msg, err := heal.ContainerHeal(container)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func hscan(w http.ResponseWriter, req *http.Request) {
	image, ok := req.URL.Query()["image"]
	if !ok || len(image[0]) < 1 {
		fmt.Fprint(w, "No file specified")
	}
	msg, err := scan.Vulnerabilities(image[0])
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func hsuggest(w http.ResponseWriter, req *http.Request) {
	file, ok := req.URL.Query()["file"]
	if !ok || len(file[0]) < 1 {
		fmt.Fprint(w, "No file specified")
	}
	msg, err := suggestions.ReadImage(file[0])
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Start the http server given a port
func Start(port string) {
	http.HandleFunc("/aheal", aheal)
	http.HandleFunc("/hcheck", hcheck)
	http.HandleFunc("/heal", hheal)
	http.HandleFunc("/scan", hscan)
	http.HandleFunc("/suggest", hsuggest)
	http.ListenAndServe(":"+port, nil)
}
