package handlers

import (
	"Pudu/internal/utils"
	"fmt"
	"log"
	"net/http"
)

// HandleLandingPage maneja la página principal
func HandleLandingPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

// HandleContactForm maneja los envíos del formulario de contacto
func HandleContactForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// Construir el cuerpo del correo
		subject := "Nuevo mensaje de contacto"
		body := fmt.Sprintf("Nombre: %s\nEmail: %s\nMensaje: %s", name, email, message)

		// Enviar el correo
		err := utils.SendEmail("tuemail@gmail.com", subject, body)
		if err != nil {
			log.Printf("Error al enviar el correo: %v", err)
			http.Error(w, "No se pudo enviar el mensaje, inténtalo de nuevo más tarde.", http.StatusInternalServerError)
			return
		}

		// Respuesta al usuario
		fmt.Fprintf(w, "Gracias, %s. Hemos recibido tu mensaje.", name)
		return
	}
	http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
}
