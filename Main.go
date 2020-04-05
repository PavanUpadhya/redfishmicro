package main

import(
	"net/http"
	"log"
)

func main(){
	// Init Router. 
	r := GetAndRegisterRouter()
	//Start the server on port 8090
	log.Fatal(http.ListenAndServe(":8090", r))
}
