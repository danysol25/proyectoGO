package main

import (
	"html/template"
	"log"
	"net/http"
)

var plantilla = template.Must(template.ParseGlob("Public/*.html"))

func main() {
	css := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", css))
	http.HandleFunc("/", index)

	log.Println("Iniciando server")
	log.Println("Cargando páginas")
	log.Println("Carga finalizada")

	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	plantilla.ExecuteTemplate(w, "index.html", nil)

}
