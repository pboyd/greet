package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pboyd/hello/internal/greeting"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, greeting.Message())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
