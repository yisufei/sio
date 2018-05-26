
package main

import (
	"encoding/json"
	"time"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name		string		`json:"name"`
	Completed	bool		`json:"completed"`
	Due			time.Time	`json:"due"`
}

type Todos	[]Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write pressentation"},
		Todo{Name: "Host meetup"},
	}
	
	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoId)
}
