package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sarthaksanjay/netflix-go/router"
)

func main() {
	fmt.Println("Netflix Apis")
	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server listing on port 4000")
}
