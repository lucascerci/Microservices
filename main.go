package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//write header allows to specify the http status code that is send back to the caller
			//rw.WriteHeader(http.StatusBadRequest)
			//writing a error message to the user
			//rw.Write([]byte("Ooops"))
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":3000", nil)
}
