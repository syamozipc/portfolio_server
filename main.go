package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func rootGet(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("home"))
	if err != nil {
		http.Error(w, "faild writing response", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("home accessed")

	return
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var s = struct {
		Name string `json:"name"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil {
		http.Error(w, "faild parsing request body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("Hello, %s", s.Name),
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "failed marshaling struct", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resJson)
	if err != nil {
		http.Error(w, "faild writing response body", http.StatusInternalServerError)
		return
	}

	fmt.Println("post accessed")

	return
}

func main() {
	http.HandleFunc("/", rootGet)

	http.HandleFunc("/post", post)

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
