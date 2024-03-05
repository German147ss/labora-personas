package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func crearPersona(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud JSON en una variable Persona
	decoder := json.NewDecoder(r.Body)
	var persona Persona
	err := decoder.Decode(&persona)
	if err != nil {
		fmt.Fprintf(w, "ERROR: "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newPersonId := insertarPersonaEnLaDb(persona)
	fmt.Println("PersonasDb", PersonasDB)
	persona.ID = newPersonId
	// Enviar una respuesta JSON con la persona creada
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)

}

func obtenerPersona(w http.ResponseWriter, r *http.Request) {
}

func actualizarPersona(w http.ResponseWriter, r *http.Request) {
}

func eliminarPersona(w http.ResponseWriter, r *http.Request) {

}
