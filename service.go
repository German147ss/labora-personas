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

func insertarPersonaEnLaDb(persona Persona) (int, error) {
	// Incrementar el ID de la persona
	var lastInsertID int
	err := DB.QueryRow("INSERT INTO personas (nombre, apellido, edad, country_code) VALUES ($1, $2, $3, $4) RETURNING id;",
		persona.Nombre, persona.Apellido, persona.Edad, persona.CountryCode).Scan(&lastInsertID)

	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// Editar persona
func editarPersona(id int, nombre string, apellido string, edad int, countryCode string) (*Persona, error) {
	var personaAux Persona
	err := DB.QueryRow("SELECT id, nombre, apellido, edad, country_code FROM personas WHERE id = $1", id).Scan(
		&personaAux.ID, &personaAux.Nombre, &personaAux.Apellido, &personaAux.Edad, &personaAux.CountryCode)
	if err != nil {
		return nil, err
	}

	fmt.Println("Persona encontrada", personaAux)
	if nombre != "" { //ger nuevo
		personaAux.Nombre = nombre
	}
	if apellido != "" {
		personaAux.Apellido = apellido
	}
	if edad != 0 {
		personaAux.Edad = edad
	}
	if countryCode != "" {
		personaAux.CountryCode = countryCode
	}

	_, err = DB.Exec("UPDATE personas SET nombre = $1, apellido = $2, edad = $3, country_code = $4 WHERE id = $5",
		personaAux.Nombre, personaAux.Apellido, personaAux.Edad, personaAux.CountryCode, id)
	if err != nil {
		return nil, err
	}
	return &personaAux, nil
}

//OBTENER PERSONA -> obtenerPersonaPorId -> Encuentra la persona y una vez encontrada, va a buscar la información relacionada a su country.

// Obtener persona
// OBTENER PERSONA -> obtenerPersonaPorId -> Encuentra la persona y una vez encontrada, va a buscar la información relacionada a su country.
// Obtener persona
func obtenerPersonaPorId(id int) (PersonaAumentada, error) {
	var persona Persona
	err := DB.QueryRow("SELECT id, nombre, apellido, edad, country_code FROM personas WHERE id = $1", id).Scan(
		&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.CountryCode)
	if err != nil {
		return PersonaAumentada{}, err
	}
	fmt.Println("Persona encontrada", persona)

	countryInfo, err := getCountryInfo(persona.CountryCode)
	if err != nil {
		fmt.Println("Error al obtener información del país %v", err)
		return PersonaAumentada{}, err
	}

	return PersonaAumentada{
		Persona:     persona,
		CountryInfo: countryInfo,
	}, nil
}

func getCountryInfo(countryCode string) (CountryInfo, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud")
		return CountryInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return CountryInfo{}, fmt.Errorf("error en la solicitud: %s", resp.Status)
	}

	var countryResponse CountryResponse
	err = json.NewDecoder(resp.Body).Decode(&countryResponse)
	if err != nil {
		fmt.Println("Error al decodificar la respuesta")
		return CountryInfo{}, err
	}

	// Tomar solo el primer país (suponiendo que solo haya uno en la respuesta)
	country := countryResponse[0]

	//obtener key del mapa

	// Extraer la información necesaria

	countryInfo := CountryInfo{
		Name:     country.Name.Common,
		Timezone: country.Timezones[0], // Tomar solo el primer timezone
		Flag:     country.Flags.PNG,
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
