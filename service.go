// CAPA DE SERVICIO
package main

type PersonaAumentada struct {
	Persona
	CountryInfo
}

func insertarPersonaEnLaDb(persona Persona) int {
	// Incrementar el ID de la persona
	persona.ID = len(PersonasDB) + 1
	PersonasDB = append(PersonasDB, persona)
	return persona.ID
}

// Editar persona
func editarPersona(id int, nombre string, apellido string, edad int, countryCode string) {

	//Encontrar persona segun su id
	for i := 0; i < len(PersonasDB); i++ {
		if PersonasDB[i].ID == id {
			if nombre != "" {
				PersonasDB[i].Nombre = nombre
			}
			if apellido != "" {
				PersonasDB[i].Apellido = apellido
			}
			if edad != 0 {
				PersonasDB[i].Edad = edad
			}
			if countryCode != "" {
				PersonasDB[i].CountryCode = countryCode
			}
		}
	}
}

//OBTENER PERSONA -> obtenerPersonaPorId -> Encuentra la persona y una vez encontrada, va a buscar la información relacionada a su country.

// Obtener persona
func obtenerPersonaPorId(id int) Persona {
	for i := 0; i < len(PersonasDB); i++ {
		if PersonasDB[i].ID == id {
			return PersonasDB[i]
		}
	}
	return Persona{}
}
