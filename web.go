package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type Result struct{
	Namespace	string		`json:"namespace"`
	Message		string		`json:"message"`
	Status		string		`json:"status"`
	Timestamp	time.Time	`json:"timestamp"`
}

type Index struct{
	Status		string		`json:"status"`
	Endpoints	[]Endpoint	`json:"endpoints"`
}

type Endpoint struct{
	Uri		string	`json:"endpoint"`
	Method	string	`json:"method"`
}

type Alive struct{
	Status		string		`json:"status"`
}

func main(){
	namespace := mux.NewRouter().StrictSlash(true)
	namespace.HandleFunc("/", isAlive)
	namespace.HandleFunc("/action", action)
	namespace.HandleFunc("/action/reboot", aReboot)
	namespace.HandleFunc("/action/shutdown", aShutdown)
	namespace.HandleFunc("/action/update", aUpdate)

	log.Fatal(http.ListenAndServe(":6660", namespace))
}

//Namespace handlers
func isAlive (w http.ResponseWriter, r *http.Request){
	outputs := Alive{Status: "success"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		panic(err)
	}
}

func action (w http.ResponseWriter, r *http.Request){
	outputs := Index{
		Status: "success",
		Endpoints: []Endpoint{
			{Uri: "reboot", Method: "GET"},
			{Uri: "shutdown", Method: "GET"},
			{Uri: "update", Method: "GET"},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		panic(err)
	}
}

func aReboot (w http.ResponseWriter, r *http.Request) {
	cmdOut := command(string("shutdown"), []string{"-r" , "+1"})

	outputs := new(Result)
	outputs.Namespace = string(r.URL.Path)
	outputs.Message = "initiated reboot"
	outputs.Timestamp = time.Time(time.Now())

	if cmdOut == nil{
		outputs.Status = "success"
	} else {
		outputs.Status = "failed: " + cmdOut.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		panic(err)
	}
}

func aShutdown (w http.ResponseWriter, r *http.Request){
	cmdOut := command(string("shutdown"), []string{"-h" , "+1"})

	outputs := new(Result)
	outputs.Namespace = string(r.URL.Path)
	outputs.Message = "initiated shutdown"
	outputs.Timestamp = time.Time(time.Now())

	if cmdOut == nil{
		outputs.Status = "success"
	} else {
		outputs.Status = "failed: " + cmdOut.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		panic(err)
	}
}

func aUpdate (w http.ResponseWriter, r *http.Request){
	cmdOut := command(string("apt-get"), []string{"upgrade" , "-y"})

	outputs := new(Result)
	outputs.Namespace = string(r.URL.Path)
	outputs.Message = "initiated system update"
	outputs.Timestamp = time.Time(time.Now())

	if cmdOut == nil{
		outputs.Status = "success"
	} else {
		outputs.Status = "failed: " + cmdOut.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		panic(err)
	}
}

//Commands
func command(cmdName string, args []string) error {
	cmd := exec.Command(cmdName, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Start()

	if err != nil {
		return fmt.Errorf("failed starting %s %v: %s: %v", cmdName, args, stderr.String(), err)
	}

	return nil
}