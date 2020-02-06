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
	io.WriteString(w, "The Force is strong with this one. Vader is commenting on the skills of an X-Wing pilot that is trying to blow up the Death Star, not realizing he’s trying to kill his own son.")
}
