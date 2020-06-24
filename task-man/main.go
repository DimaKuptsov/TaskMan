package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", test)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("%s", err.Error())
		os.Exit(1)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test page")
}
