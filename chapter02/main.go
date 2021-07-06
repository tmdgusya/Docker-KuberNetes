package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func (w http.ResponseWriter, r * http.Request)  {
		log.println("received request")
		fmt.println(w, "Hello Docker!")
	})

	log.println("start server")
	server: = &http.Server {
		Addr: ":8080"
	}
	if err: = server.ListenAndServe(); err != nil {
		log.println(err)
	}
}