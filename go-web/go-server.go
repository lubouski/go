package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("date", "+%F").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Whoa Go is neat! today %s\n", out)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

