package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"heroNameGenerator/suffixGenerator"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")

	seed := time.Now().UTC().UnixNano()
	suffixGenerator := suffixGenerator.NewSuffixGenerator(seed)

	nameWithSuffix := suffixGenerator.Generate(name)

	fmt.Fprintf(w, "Your suggested role is: %s\n", nameWithSuffix)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form-submit", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
