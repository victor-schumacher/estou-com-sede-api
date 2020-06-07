package main

import (
	"github.com/victor-schumacher/estou-com-sede-api/api/handler"
	"log"
	"net/http"

)

func main(){

	plantManager := handler.NewPlantManager()

	http.Handle("/humidity" , plantManager.Handle())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
