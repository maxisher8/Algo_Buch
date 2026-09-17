
ackage main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"/tdas/cola/cola_enlazada"
)

func Resolver(laberinto [][]rune) {
	filas := len(laberinto)
	columnas := len(laberinto[0])
	visitados := cola.CrearColaEnlazada[int]()

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
		Resolver(laberinto)
	}
}