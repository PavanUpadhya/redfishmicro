package main

import(
	"net/http"
	"log"
)

func main(){
	// Init Router
	r := GetAndRegisterRouter()
	//Start the server
	log.Fatal(http.ListenAndServe(":8090", r))
}
