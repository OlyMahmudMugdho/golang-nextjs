package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeFile)
	fmt.Println("server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func ServeFile(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSuffix(r.URL.Path, "/")
	path = strings.TrimPrefix(path, "/")

	// fmt.Println("./out/" + path + ".html")
	// fmt.Println(filepath.Join(path) + " here")

	_, err := os.Stat("./out/" + path + ".html")

	if err == nil {
		http.ServeFile(w, r, "./out/"+path+".html")
	} else {
		http.ServeFile(w, r, "./out/index.html")
	}

	// fmt.Println(path)
	//http.ServeFile(w, r, "./out/index.html")
}
