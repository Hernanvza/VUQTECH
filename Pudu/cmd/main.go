package main

import (
	"log"
	"net/http"
)

func main() {
	// Ruta para servir archivos estáticos (CSS, JS, imágenes)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Ruta para la página principal
	http.HandleFunc("/", handlers.HandleLandingPage)

	// Ruta para el formulario de contacto
	http.HandleFunc("/contact", handlers.HandleContactForm)

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
