package main

import (
		"fmt"
		"log"
		"net/http"

		bservice "brawlstarsclub.com/m/v2/service/brawlstars"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
		
    http.HandleFunc("/club", bservice.GetClub)

    if err := http.ListenAndServe(":8090", nil); err != nil {
			log.Fatal(err)
		}
}