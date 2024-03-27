// CAPA DE INFRAESTRUCTURA
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	// Validar la persona
	if !persona.Validate() {
		fmt.Fprintf(w, "ERROR: La persona no es válida")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newPersonId, err := insertarPersonaEnLaDb(persona)
	if err != nil {
		fmt.Fprintf(w, "ERROR: "+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("PersonasDb", PersonasDB)
	persona.ID = newPersonId
	// Enviar una respuesta JSON con la persona creada
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)

}
func obtenerPersona(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idAsInt, _ := strconv.Atoi(idString)

	persona, err := obtenerPersonaPorId(idAsInt)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(persona)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func actualizarPersona(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var persona Persona
	err := decoder.Decode(&persona)
	if err != nil {
		fmt.Fprintf(w, "ERROR: "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	personaPtr, err := editarPersona(persona.ID, persona.Nombre, persona.Apellido, persona.Edad, persona.CountryCode)
	if err != nil {
		fmt.Fprintf(w, "ERROR: "+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(*personaPtr)
}

func eliminarPersona(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idAsInt, _ := strconv.Atoi(idString)
	eliminarPersonaPorId(idAsInt)
	w.WriteHeader(http.StatusOK)
}
