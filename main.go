package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Servidor go rodando")
	HandleRequest()
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
