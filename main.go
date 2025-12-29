package main

import "fmt"

func main() {

	// 1. Declarar variable nombre inferida
	nombre := "Arroz x 1gk" // De esta manera, Go infiere que el tipo de la variable es string

	// 2. Declarar variable stock definida
	var stock int = 10 // De esta manera, declaramos el tipo de la variable y esta ya no puede cambiar de tipo

	// 3. Declarar variable precio inferida
	precio := 1200.50 // De esta manera, Go infiere que el tipo de la variable es float64

	// 3.a Declarar variable precio definida
	// var precio float64 = 1200.50 // De esta otra manera, definimos el tipo de la variable

	// 4. Impresión por consola de los datos
	fmt.Printf("El producto %s tiene un precio de $%.2f y actualmente quedan %d en stock \n", nombre, precio, stock)

}