package main

import (
	"fmt"
	"github.com/victor-schumacher/estou-com-sede-api/api/handler"
	"log"
	"net/http"
	"os"
)

func main(){

	plantManager := handler.NewPlantManager()

	http.Handle("/humidity" , plantManager.Handle())
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, nil))
}
