package main

import (
	"net/http"
)

func main() {
	/*
		En Go, un multiplexor es una herramienta que te permite enrutar solicitudes HTTP a diferentes manejadores, similar a como funciona un "switch".

			Este multiplexor te permite definir rutas y asociarlas a funciones específicas (manejadores) que se ejecutan cuando llega una solicitud a esa ruta.
			Es una opción simple y eficiente para aplicaciones web pequeñas
	*/
	router := http.NewServeMux()

	// Crear
	router.HandleFunc("POST /personas", crearPersona)

	// Iniciar el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", router)
}
