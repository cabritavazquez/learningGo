package utils

import "fmt"

// Para que la funcion Saludar este exportada tiene que comenzar con Mayuscula
func Saludar(nombre string) {
    fmt.Println("Hola, " + nombre + "! Bienvenido a GoLand!")
}

func Test(){
	fmt.Println("Test function")
}
