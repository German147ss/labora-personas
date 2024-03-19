// CAPA DE SERVICIO
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var DB *sql.DB

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
func obtenerPersonaPorId(id int) PersonaAumentada {
	for i := 0; i < len(PersonasDB); i++ {
		if PersonasDB[i].ID == id {
			aux := PersonasDB[i]
			countryInfo, err := getCountryInfo(aux.CountryCode)
			if err != nil {
				fmt.Println("Error al obtener información del país")
			}
			return PersonaAumentada{
				Persona:     aux,
				CountryInfo: countryInfo,
			}
		}
	}
	return PersonaAumentada{}
}

func getCountryInfo(countryCode string) (CountryInfo, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil {
		return CountryInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CountryInfo{}, fmt.Errorf("error en la solicitud: %s", resp.Status)
	}

	var countryResponse CountryResponse
	err = json.NewDecoder(resp.Body).Decode(&countryResponse)
	if err != nil {
		return CountryInfo{}, err
	}

	// Tomar solo el primer país (suponiendo que solo haya uno en la respuesta)
	country := countryResponse[0]

	//obtener key del mapa

	// Extraer la información necesaria

	countryInfo := CountryInfo{
		Name:     country.Name.Common,
		Timezone: country.Timezones[0], // Tomar solo el primer timezone
	}

	for key, _ := range country.Currencies {
		countryInfo.Currency = key //obtenemos el simbolo de la moneda
	}

	return countryInfo, nil
}

func eliminarPersonaPorId(id int) {
	_, err := DB.Exec("DELETE FROM personas WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error al eliminar la persona")
	}
	fmt.Println("Persona eliminada")
}

func initDB() {
	var err error
	DB, err = sql.Open("postgres", "user=alfred dbname=labora host=localhost sslmode=disable password=4lfr3d port=5431")
	if err != nil {
		fmt.Println("Error en la conexión a la base de datos")
		panic(err)
	}
}
