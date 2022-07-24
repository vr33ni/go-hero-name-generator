package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"heroNameGenerator/suffixgenerator"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")

	nameWithSuffix := getHeroName(name)

	fmt.Fprintf(w, "Your suggested role is: %s\n", nameWithSuffix)
}

func getHeroName(name string) string {
	seed := time.Now().UTC().UnixNano()
	generator := suffixgenerator.NewHeroNameGenerator(seed)

	nameWithSuffix := generator.Generate(name)
	return nameWithSuffix

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form-submit", handleForm)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
