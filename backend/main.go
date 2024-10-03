package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", ServeWebFile)
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

func ServeWebFile(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	path = strings.TrimSuffix(path, "/")

	file := "../out/" + path + ".html"

	_, err := os.Stat(file)

	if err != nil {
		http.ServeFile(w, r, "../out/index.html")
	} else {
		http.ServeFile(w, r, file)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ok")
}
