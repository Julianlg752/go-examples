package main

import (
    "io"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "net/http"
)

// curl localhost:8000 -d '{"name":"Hello", "email":"a@s.com"}'
type User struct {
	User         string `json:"user"`
	Email        string `json:"email"`
}

func ping(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Hello world!")
}

func post(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		}

		var us User
		err = json.Unmarshal(body, &us)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		output, err := json.Marshal(us)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(output)
	}

	if r.Method == "GET" {
		io.WriteString(w, "Invalid")
	}
}


func main() {
    http.HandleFunc("/ping", ping)
    http.HandleFunc("/", post)
    fmt.Println("Listening: 8001")
    http.ListenAndServe(":8001", nil)
}




