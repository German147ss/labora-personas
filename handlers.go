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

	newPersonId := insertarPersonaEnLaDb(persona)
	fmt.Println("PersonasDb", PersonasDB)
	persona.ID = newPersonId
	// Enviar una respuesta JSON con la persona creada
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)

}

func obtenerPersona(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la persona de la ruta
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Buscar la persona en el slice por su ID
	persona, encontrado := buscarPersonaPorID(id)
	if !encontrado {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Enviar una respuesta JSON con la persona
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)

}

/*
func actualizarPersona(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {

		// Obtener el ID de la persona de la ruta
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Decodificar el cuerpo de la solicitud JSON en una variable Persona
		decoder := json.NewDecoder(r.Body)
		var persona Persona
		err = decoder.Decode(&persona)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Buscar la persona en el slice por su ID
		personaEncontrada, encontrado := buscarPersonaPorID(id)
		if !encontrado {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Actualizar la información de la persona en el slice
		personaEncontrada.Nombre = persona.Nombre
		personaEncontrada.Apellido = persona.Apellido
		personaEncontrada.Edad = persona.Edad

		// Enviar una respuesta JSON con la persona actualizada
		encoder := json.NewEncoder(w)
		encoder.Encode(personaEncontrada)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func eliminarPersona(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodDelete {
		// Obtener el ID de la persona de la ruta
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Buscar la posición de la persona en el slice por su ID
		indice, encontrado := buscarIndicePersonaPorID(id)
		if !encontrado {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Eliminar la persona del slice
		PersonasDB = append(PersonasDB[:indice], PersonasDB[indice+1:]...)

		// Enviar una respuesta JSON con un mensaje de éxito
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string]string{"mensaje": "Persona eliminada con éxito"})
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
*/
