package variables

import "fmt"

func MuestroEnteros() {
	// Variables por declaración
	var intcomun int

	// Variable por asignación
	intde32 := int32(10) // toma el tipo de dato de la variable q le estoy asignando

	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)

}
