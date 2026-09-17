package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const (
	RUTA_ARCHIVO_1 = "archivo1.in"
	RUTA_ARCHIVO_2 = "archivo2.in"
)

func lectura(ruta string) []int {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	vector := []int{}
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		contenido, _ := strconv.Atoi(lectura.Text())
		vector = append(vector, contenido)
	}
	return vector
}

func imprimirVectorOrdenado(vector []int) {
	ejercicios.Seleccion(vector)
	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}
}

func main() {
	vector1 := lectura(RUTA_ARCHIVO_1)
	vector2 := lectura(RUTA_ARCHIVO_2)

	igualdad := ejercicios.Comparar(vector1, vector2)

	if igualdad >= 0 {
		imprimirVectorOrdenado(vector1)
	} else {
		imprimirVectorOrdenado(vector2)
	}
}
