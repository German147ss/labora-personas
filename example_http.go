package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Countries []Country

type Country struct {
	Name       Name       `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Timezones  []string   `json:"timezones"`
	Flags      Flags      `json:"flags"`
}

type CountryInfo2 struct {
	Name Name `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Timezone string `json:"timezone"`
}

type Currencies struct {
	Ars Ars `json:"ARS"`
}

type Ars struct {
	Name string `json:"name"`
}

type Flags struct {
	PNG string `json:"png"`
}

type Name struct {
	Common string `json:"common"`
}

func nomain() {
	url := "https://restcountries.com/v3.1/alpha/arg"
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
	fmt.Printf("%+v \n", countries[0].Flags.PNG)
}
