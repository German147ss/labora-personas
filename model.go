package main

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var PersonasDB = []Persona{} //persona 1: id:1 , someData{}

//Validación
