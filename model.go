// CA
package main

type Persona struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Edad        int    `json:"edad"`
	CountryCode string `json:"countryCode"`
}

var PersonasDB = []Persona{}

// Validación
func (p *Persona) Validate() bool {
	if p.Nombre == "" || p.Apellido == "" || p.Edad == 0 || p.CountryCode == "" {
		return false
	}
	return true
}

type CountryResponse []struct {
	Name       Name                `json:"name"`
	Timezones  []string            `json:"timezones"`
	Currencies map[string]struct { //symbol
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Flags Flags `json:"flags"`
}

type Flags struct {
	PNG string `json:"png"`
}

type Name struct {
	Common string `json:"common"`
}

type CountryInfo struct {
	Name     string
	Timezone string
	Currency string
	Flag     string
}
