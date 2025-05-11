package main

import (
	"Pudu/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Definir las rutas
	http.HandleFunc("/", handlers.HandleLandingPage)        // Página principal
	http.HandleFunc("/contact", handlers.HandleContactForm) // Manejador del formulario de contacto

	// Iniciar el servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
