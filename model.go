// CA
package main

type Persona struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Edad        int    `json:"edad"`
	CountryCode string `json:"countryCode"`
}

var PersonasDB = []Persona{} //persona 1: id:1 , someData{}

//Validación

func (p *Persona) Validate() bool {
	if p.Nombre == "" || p.Apellido == "" || p.Edad == 0 || p.CountryCode == "" {
		return false
	}
	return true
}

type CountryResponse []struct {
	Name       Name     `json:"name"`
	Timezones  []string `json:"timezones"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
}

type CountryInfo struct {
	Name     string
	Timezone string
	Currency string
}
