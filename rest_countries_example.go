// Esto es solo un ejemplo de como implementar una llamada a un servicio externo.
package main

import (
	"fmt"
	"net/http"
)

type CountryResponse []struct {
	Name       Name     `json:"name"`
	Timezones  []string `json:"timezones"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
}

type NameO struct {
	Common string `json:"common"`
}

type CountryInfo struct {
	Name     string
	Timezone string
	Currency string
}

func getCountryInfo(countryCode string) (CountryInfo, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil {
		return CountryInfo{}, err
	}
	defer resp.Body.Close()
	/*
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
	*/
	/* // Extraer la información necesaria
	countryInfo := CountryInfo{
		Name:     country.Name.Common,
		Timezone: country.Timezones[0],           // Tomar solo el primer timezone
		Currency: country.Currencies["COP"].Name, // Tomar la moneda COP
	} */

	return CountryInfo{}, nil
}
