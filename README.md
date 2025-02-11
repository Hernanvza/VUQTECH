# Pudu - Tu Aliado en Tecnología

## Descripción
Pudu es una página web tecnológica que ofrece soluciones en desarrollo web, QA y automatización RPA. La plataforma está construida con tecnologías modernas para garantizar una experiencia fluida y óptima para los usuarios.

---

## Tecnologías Utilizadas

### Frontend:
- **HTML5**: Para la estructura de la página.
- **CSS3**: Estilos y efectos visuales (incluyendo Glassmorphism).
- **Bootstrap 5**: Framework CSS para diseño responsivo y componentes.
- **Spline**: Integración de modelos 3D en la sección Hero.
- **JavaScript**: Interactividad, animaciones y manejo de eventos.

### Backend:
- **Golang (Go)**: Lenguaje principal para la lógica del backend.
- **Gin Framework**: Para el manejo de rutas y API.

### Otras Tecnologías:
- **Bootstrap Icons**: Para los íconos de servicios y redes sociales.
- **WhatsApp API**: Botón de contacto flotante.
- **Google Fonts (Inter)**: Tipografía de la página.

---

## Estructura del Proyecto

```
📂 pudu-project
├── 📂 assets
│   ├── 📂 img  (Imágenes y logos)
│   ├── 📂 styles (CSS personalizado)
│   │   ├── styles.css
│   ├── 📂 js (Scripts adicionales)
│
├── 📂 backend (Código en Go)
│   ├── main.go
│   ├── router.go
│
├── 📂 templates (Archivos HTML base)
│   ├── index.html
│   ├── footer.html
│
├── README.md
└── .gitignore
```

---

## Pasos para Crear el Proyecto

### 1. Configurar el Entorno
- Instalar Go: [Descargar aquí](https://go.dev/dl/)
- Instalar Node.js (opcional para desarrollo frontend avanzado).
- Configurar un servidor local (puedes usar `go run main.go`).

### 2. Crear el Backend en Go
- Inicializar el proyecto con `go mod init pudu-project`.
- Instalar Gin con `go get -u github.com/gin-gonic/gin`.
- Configurar rutas y handlers en `main.go` y `router.go`.

### 3. Construir el Frontend
- Diseñar la estructura HTML con Bootstrap.
- Implementar estilos personalizados en `styles.css`.
- Integrar modelos 3D de Spline en la sección Hero.
- Configurar eventos y animaciones con JavaScript.

### 4. Agregar Funcionalidades Adicionales
- Implementar botón flotante de WhatsApp.
- Crear el popup de "Términos y Condiciones".
- Mejorar interactividad con efectos visuales.

### 5. Pruebas y Despliegue
- Probar la interfaz en dispositivos móviles.
- Revisar compatibilidad en navegadores.
- Desplegar en un servidor o servicio cloud (ej. Vercel, DigitalOcean, AWS).

---

## Contacto
Si tienes preguntas o quieres contribuir, contáctanos en nuestras redes sociales o envíanos un mensaje a través del formulario en la web.

---
🚀 **Pudu - Innovación y Tecnología**

