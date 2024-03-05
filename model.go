package main

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var PersonasDB = []Persona{} //persona 1: id:1 , someData{}

// buscarPersonaPorID busca una persona en la base de datos(en memoria "PersonasDB") por su ID, y retorna la persona y un booleano que indica si la persona fue encontrada.
func buscarPersonaPorID(id int) (Persona, bool) {
	for _, persona := range PersonasDB {
		if persona.ID == id {
			return persona, true
		}
	}
	return Persona{}, false
}

func buscarIndicePersonaPorID(id int) (int, bool) {
	for i, persona := range PersonasDB {
		if persona.ID == id {
			return i, true
		}
	}
	return -1, false
}

func insertarPersonaEnLaDb(persona Persona) int {
	// Incrementar el ID de la persona
	persona.ID = len(PersonasDB) + 1
	PersonasDB = append(PersonasDB, persona)
	return persona.ID
}
