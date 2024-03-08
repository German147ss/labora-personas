// CAPA DE SERVICIO
package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

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
func obtenerPersonaPorId(id int) PersonaAndCountry {
	for i := 0; i < len(PersonasDB); i++ {
		p := PersonasDB[i]
		if p.ID == id {
			countryInfo, _ := getCountryInfo2(p.CountryCode)
			return PersonaAndCountry{p, countryInfo}
		}
	}
	return PersonaAndCountry{}
}

func getCountryInfo2(countryCode string) (cInfo CountryInfo2, e error){
	url := fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("respuesta:", resp.Body)

	decoder := json.NewDecoder(resp.Body)

	var countries Countries
	err = decoder.Decode(&countries)
	if err != nil {
		fmt.Println(err)
	}
	country := countries[0]
	return CountryInfo2{
			Name:country.Name,
			Currencies:country.Currencies,
			Timezone:country.Timezones[0],
		}, 
		nil
}
