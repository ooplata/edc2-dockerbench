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
	suma := 2
	contador := 1
	numero := 3

	for contador < n {
		if esPrimo(numero) {
			suma += numero
			contador++
		}
		numero += 2
	}

	return suma
}

func escribirPrimerosPrimos(n int) {
	primos := sumaPrimerosPrimos(n)

	// Crear o abrir el archivo out.txt para escribir
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir la suma en un archivo
	_, err = fmt.Fprintf(file, "%d\n", primos)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

// Función para medir el tiempo de ejecución
func medirTiempoEjecucion(f func(int), n int) float64 {
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
	tiempo := medirTiempoEjecucion(escribirPrimerosPrimos, n)

	// También imprimir en la consola
	fmt.Printf("%.2fms\n", tiempo)
}
