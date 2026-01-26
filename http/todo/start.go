package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var notes []note

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", add)
	mux.HandleFunc("/get", get)
	mux.HandleFunc("/delete", Delete)
	mux.HandleFunc("/update", update)

	http.ListenAndServe(":8080", mux)
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var newNote note

	err := json.NewDecoder(r.Body).Decode(&newNote)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("Note added"))

}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	notes = append(notes[:id], notes[id:]...)
	w.Write([]byte("Note deleted"))
}
func update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	notes[id].Status = "completed"
	w.Write([]byte("Note updated"))
}
