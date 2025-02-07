package main

import (
	"fmt"
	"os"
	"time"
)

// Función para verificar si un número es primo
func esPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Función para calcular la suma de los primeros 'n' números primos
func sumaPrimerosPrimos(n int) int {
	suma := 0
	contador := 0
	numero := 2

	for contador < n {
		if esPrimo(numero) {
			suma += numero
			contador++
		}
		numero++
	}

	return suma
}

// Función para medir el tiempo de ejecución
func medirTiempoEjecucion(f func(int) int, n int) float64 {
	// Registra el inicio del tiempo
	inicio := time.Now()

	// Llama a la función que calcula la suma de los primeros 'n' primos
	f(n)

	// Calcula el tiempo que pasó desde el inicio
	fin := time.Since(inicio).Milliseconds() // Milisegundos

	// Imprime el tiempo de ejecución
	return float64(fin)
}

func main() {
	// Número de primos a sumar
	n := 10000

	// Medir el tiempo que tarda en calcular los primeros 'n' primos
	tiempo := medirTiempoEjecucion(sumaPrimerosPrimos, n)

	// Crear o abrir el archivo output.txt para escribir
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir el tiempo de ejecución en el archivo
	_, err = fmt.Fprintf(file, "%.2f \n", tiempo)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	// También imprimir en la consola
	fmt.Printf("%.2f \n", tiempo)
}
