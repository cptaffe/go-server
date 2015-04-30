package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Test struct {
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	test := Test{
		fmt.Sprintf("Hi there, I love %s!", r.URL.Path[1:]),
	}

	b, err := json.Marshal(test)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", b)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
