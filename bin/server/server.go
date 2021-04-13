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
	containers := req.URL.Query().Get("containers")
	c := []string{containers}
	msg, err := healthcheck.PerformHealthCheck(c)
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
	containers := req.URL.Query().Get("containers")
	c := []string{containers}
	msg, err := autoheal.AutoHeal(c)
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
	containers := req.URL.Query().Get("containers")
	c := []string{containers}
	msg, err := heal.ContainerHeal(c)
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
	image := req.URL.Query().Get("image")
	file := req.URL.Query().Get("file")
	fmt.Println("PARAMS", image, file)
	if len(image) < 1 {
		fmt.Fprint(w, "No file specified")
	}
	out, err := scan.Vulnerabilities(image)
	if err != nil {
		log.Fatal(err)
	}
	if len(file) > 1 {
		scan.WriteFile(out, file)
		out = "Succesfully wrote scan to " + file
	}
	data, err := json.Marshal(out)
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
