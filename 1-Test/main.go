package main

// Con package es la palabra que empiezan todos los archivos de go
// Como el archivo se llama main y esta en la carpeta principal, será package main
// Para otros casos, como en el carpeta de utils, el package será package utils, ya
// ya que se encuentra en la carpeta de utils

import (
	"fmt"

	"github.com/cabritavazquez/learningGo/1-Test/nomuyutils"
	"github.com/cabritavazquez/learningGo/1-Test/utils"
)

func main() {
	// Esta es la funcion principal de Go, tiene que empezar con el nombre: main

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
