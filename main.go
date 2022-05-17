package main

import (
    "fmt"
    "log"
	"os"
	"net"
    "net/http"
	"strconv"
	"encoding/json"

)

type Pod struct {
	Name		string `json:"name"`
	PID			string `json:"os_pid"`
    Ip 			string `json:"ip"`
    Port		string `json:"port"`
}

var my_pod Pod
var port = "8900" 

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage Test go 1.0 !")
}

func health(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: health")
	json.NewEncoder(w).Encode(my_pod)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
	http.HandleFunc("/health", health)
    log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	fmt.Println("Inicio server http port ", port)

	// Buscando IP e PID
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Erro Fatal")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				my_pod.Ip = ipnet.IP.String()
			}
		}
	}

	my_pod.Port = envVariable("PORT")
	if my_pod.Port == ""{
		my_pod.Port = "8900"
	}
	my_pod.Name = envVariable("NAME_POD")
	if my_pod.Name == ""{
		my_pod.Name = "no-name"
	}

	my_pod.PID = strconv.Itoa(os.Getpid())

    handleRequests()
}
