package main

import (
	"learningGo/utils"
	"fmt"
	"learningGo/nomuyutils"
)

func main() {
	// Declaración de una variable entera
	var edad int
	edad = 25

	// Declaración y asignación en una línea
	var nombre string = "Juan"

	// Inferencia de tipo
	apellido := "Pérez"

	fmt.Println("Nombre:", nombre)
	fmt.Println("Apellido:", apellido)
	fmt.Println("Edad:", edad)

	utils.Saludar("Luis")
	nomuyutils.Testea()

	//Saludar("Luisk")
}
