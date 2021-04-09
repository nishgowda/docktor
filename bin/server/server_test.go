package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestHcheck(t *testing.T) {
  req, err := http.NewRequest("GET", "/hcheck", nil)
  if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(hcheck)
  handler.ServeHTTP(rr, req)
  status := rr.Code
  if status != http.StatusOK {
    t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }
  expected, err := json.Marshal("Succesfully added health checks to the following container: nginx")
  if err != nil {
    t.Fatal(err)
  }
  if strings.Compare(string(expected), rr.Body.String()) == 0 {
    t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
  }
}

func TestAheal(t *testing.T) {
  req, err := http.NewRequest("GET", "/aheal?containers=ng", nil)
  if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(aheal)
  handler.ServeHTTP(rr, req)
  status := rr.Code
  if status != http.StatusOK {
    t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }
  expected, err := json.Marshal("Succesfully added auto heals to the following container: ")
  if err != nil {
    t.Fatal(err)
  }
  if strings.Contains(string(expected), rr.Body.String()) {
    t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
  }
}

func TestHeal(t *testing.T) {
  req, err := http.NewRequest("GET", "/heal?containers=ng", nil)
  if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(hheal)
  handler.ServeHTTP(rr, req)
  status := rr.Code
  if status != http.StatusOK {
    t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }
  expected, err := json.Marshal("Restarted container: ")
  if err != nil {
    t.Fatal(err)
  }
  if strings.Contains(string(expected), rr.Body.String()) {
    t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
  }
}

func TestScan(t *testing.T) {
  req, err := http.NewRequest("GET", "/scan?image=nginx", nil)
  if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(hscan)
  handler.ServeHTTP(rr, req)
  status := rr.Code
  if status != http.StatusOK {
    t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }
  expected, err := json.Marshal("Successfully wrote vulnerability report to ")
  if err != nil {
    t.Fatal(err)
  }
  if strings.Contains(string(expected), rr.Body.String()) {
    t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
  }
}

func TestSuggest(t *testing.T) {
  req, err := http.NewRequest("GET", "/suggest?file=/Users/nishgowda/Desktop/Code/Projects/docktor/testdata/Dockerfile", nil)
  if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(hsuggest)
  handler.ServeHTTP(rr, req)
  status := rr.Code
  if status != http.StatusOK {
    t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }
  expected, err := json.Marshal("Detected no issues")
  if err != nil {
    t.Fatal(err)
  }
  if strings.Contains(string(expected), rr.Body.String()) {
    t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
  }
}