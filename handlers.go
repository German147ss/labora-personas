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

	newPersonId := insertarPersonaEnLaDb(persona)
	fmt.Println("PersonasDb", PersonasDB)
	persona.ID = newPersonId
	// Enviar una respuesta JSON con la persona creada
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)

}

func obtenerPersona(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	//convertir id a int
	idAsInt, _ := strconv.Atoi(id)
	//buscar persona por id
	persona := obtenerPersonaPorId(idAsInt)
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)

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
	editarPersona(persona.ID, persona.Nombre, persona.Apellido, persona.Edad, persona.CountryCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)
}

func eliminarPersona(w http.ResponseWriter, r *http.Request) {

}
