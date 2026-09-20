package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola"
)

const (
	PARED = '#'
	SALIDA = 'S'
	LLEGADA = 'E'
	ERROR = "ERROR"
)

var DIRRECCIONES = []direccion{
	{fila: -1, columna: 0, apuntado: "ARRIBA"},
	{fila: 1, columna: 0, apuntado: "ABAJO"},
	{fila: 0, columna: -1, apuntado: "IZQUIERDA"},
	{fila: 0, columna: 1, apuntado: "DERECHA"},
}

type coordenada struct {
	fila, columna int
}

type direccion struct {
	fila, columna int
	apuntado string
}

type origen struct {
	anterior coordenada
	movimiento string
}

func buscarsalidaYllegada(laberinto [][]rune) (salida, llegada coordenada) {
	for i := 0; i < len(laberinto); i++ {
		for j := 0; j < len(laberinto[i]); j++ {
			if laberinto[i][j] == SALIDA {
				salida = coordenada{fila: i, columna: j}
			} else if laberinto[i][j] == LLEGADA {
				llegada = coordenada{fila: i, columna: j}
			}
		}
	}
	return salida, llegada
}

func esParedYEstaAdentro(laberinto [][]rune, c coordenada) bool {
	return c.fila >= 0 && c.fila < len(laberinto) &&
		c.columna >= 0 && c.columna < len(laberinto[0]) &&
		laberinto[c.fila][c.columna] != PARED
}

func invertirVector(vector []string) {
	for i := 0; i < len(vector)/2; i++ {
		vector[i], vector[len(vector)-1-i] = vector[len(vector)-1-i], vector[i]
	}
}

func armarCamino(origenes [][]origen, salida, llegada coordenada) {
	var pasos []string
	actual := llegada
	inicio := salida
	for actual != inicio {
		origenActual := origenes[actual.fila][actual.columna]
		pasos = append(pasos, origenActual.movimiento)
		actual = origenActual.anterior
	}
	invertirVector(pasos)
	for _, paso := range pasos {
		fmt.Println(paso)
	}
}

func resolver(laberinto [][]rune) bool {
	origenes := make([][]origen, len(laberinto))
	for i := 0; i < len(laberinto); i++ {
		origenes[i] = make([]origen, len(laberinto[i]))
	}
	salida, llegada := coordenada{-1,-1}, coordenada{-1,-1}
	salida, llegada = buscarsalidaYllegada(laberinto)
	if salida.fila == -1 || llegada.fila == -1 {
		return false
	}
	celdas := cola.CrearColaEnlazada[coordenada]()
	celdas.Encolar(salida)
	laberinto[salida.fila][salida.columna] = PARED
	for !celdas.EstaVacia() {
		actual := celdas.Desencolar()
		if actual == llegada {
			armarCamino(origenes, salida, llegada)
			return true
		}
		for _, dir := range DIRRECCIONES {
			celdaVecina := coordenada{fila: actual.fila + dir.fila, columna: actual.columna + dir.columna}
			if esParedYEstaAdentro(laberinto, celdaVecina) {
				celdas.Encolar(celdaVecina)
				laberinto[celdaVecina.fila][celdaVecina.columna] = PARED
				origenes[celdaVecina.fila][celdaVecina.columna] = origen{anterior: actual, movimiento: dir.apuntado}
			}
		}
	}
	return false
}

func main() {
	lectura := bufio.NewScanner(os.Stdin)
	for lectura.Scan() {
		partes := strings.Split(lectura.Text(), " ")
		filas, _ := strconv.Atoi(partes[0])
		columnas, _ := strconv.Atoi(partes[1])
		laberinto := make([][]rune, filas)
		for i := 0; i < filas; i++ {
			lectura.Scan()
			linea := lectura.Text()
			laberinto[i] = make([]rune, columnas)
			for j := 0; j < columnas; j++ {
				laberinto[i][j] = rune(linea[j])
			}
		}
		if !resolver(laberinto) {
			fmt.Println(ERROR)
		}
	}
}
