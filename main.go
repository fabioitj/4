package main;

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"math/rand"
	"strconv"
);

type Client struct {
	ID string `json: "id"`
	Name string `json: "name"`
	Age int `json: "age"`
	Height float64 `json: "height"`
}

var clients []Client;

func getClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(clients);
}

func getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	for _, item := range clients {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item);
			return;
		}
	}

	json.NewEncoder(w).Encode(&Client{});
}

func createClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	var client Client;
	_ = json.NewDecoder(r.Body).Decode(&client);
	client.ID = strconv.Itoa(rand.Intn(1000000));
	clients = append(clients, client);

	json.NewEncoder(w).Encode(client);
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	for index, item := range clients {
		if item.ID == params["id"] {
			clients = append(clients[:index], clients[index+1:]...);
			
			var client Client;
			_ = json.NewDecoder(r.Body).Decode(&client);
			client.ID = item.ID;
			clients = append(clients, client);

			json.NewEncoder(w).Encode(client);
			return;
		}
	}

	json.NewEncoder(w).Encode(&Client{});
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);

	for index, item := range clients {
		if item.ID == params["id"] {
			clients = append(clients[:index], clients[index+1:]...);
			break;
		}
	}

	json.NewEncoder(w).Encode(clients);
}


func main() {
	r := mux.NewRouter();

	clients = append(clients, Client{ID: "1", Name: "Fabio", Age: 20, Height: 1.76});
	clients = append(clients, Client{ID: "2", Name: "Isabel", Age: 20, Height: 1.59});
	clients = append(clients, Client{ID: "3", Name: "Vanessa", Age: 40, Height: 1.60});

	r.HandleFunc("/api/client", getClients).Methods("GET");
	r.HandleFunc("/api/client/{id}", getClient).Methods("GET");
	r.HandleFunc("/api/client", createClient).Methods("POST");
	r.HandleFunc("/api/client/{id}", updateClient).Methods("PUT");
	r.HandleFunc("/api/client/{id}", deleteClient).Methods("DELETE");

	log.Fatal(http.ListenAndServe(":8090", r));
}