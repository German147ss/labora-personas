package main

import (
	"testing"
)

// Test tables
type TestDeSuma struct {
	arg1, arg2, expected int
}

var testsDeSuma = []TestDeSuma{
	TestDeSuma{2, 3, 5},
	TestDeSuma{4, 8, 12},
	TestDeSuma{6, 9, 15},
	TestDeSuma{3, 10, 13},
	TestDeSuma{3, 5, 8},
	TestDeSuma{-1, 1, 0},
}

func TestSuma(t *testing.T) {
	// Prueba con números positivos
	for _, test := range testsDeSuma {
		resultado := Suma(test.arg1, test.arg2)
		if resultado != test.expected {
			t.Errorf("Suma(%d, %d) = %d; se esperaba %d", test.arg1, test.arg2, resultado, test.expected)
		}
	}
}

func BenchmarkSuma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suma(4, 6)
	}
}

// para correr los benchmark
// go test -benchmem -bench=. github.com/vituchon/labora-golang-course/meeting-benchmark
// go test -bench=.  ./ // todos los benchs
// go test -bench=BenchmarkPrint . // los que sigan el patron
// se puede agregar flag -run=none (para que no ejecute tests!) => go test -run=none  -bench .  ./
// https://apuntes.de/golang/pruebas-benchmark/#gsc.tab=0
// https://www.golinuxcloud.com/golang-benchmark/

func BenchmarkPrintWhenIsSaturdaySwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printWhenIsSaturdaySwitch()
	}
}

func BenchmarkPrintWhenIsSaturdayIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printWhenIsSaturdayIf()
	}
}
