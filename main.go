package main

import (
	"fmt"
	"go-react-todo/router"
	"log"
	"net/http"
);

func main() {
	r := router.Router()
	fmt.Println("Starting the server on PORT 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))

}