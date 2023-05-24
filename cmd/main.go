package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {
	// asciiArtFs.CheckHash("")
	log.Print("http://localhost:8080")
	http.HandleFunc("/", handlers.MainPage)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
