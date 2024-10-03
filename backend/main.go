package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("../out/_next/"))
	router.Handle("/_next/", http.StripPrefix("/_next/", fs))

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
		_, err := os.Stat("../out/static")
		if err == nil {
			fmt.Println("exists")
		}
		http.ServeFile(w, r, "../out/index.html")
	} else {
		http.ServeFile(w, r, file)
	}
}
