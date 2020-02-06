package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hellavader", nowTime)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("An error ocurred: %s\n", err)
	}
}

func nowTime(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Be Careful Not To Choke On Your Convictions")
}
