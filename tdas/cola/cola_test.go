package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}

func TestColaDesencolarVerPrimeroVacio(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}

func TestColaRecienCreadaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
}

func TestColaDesencolarVerPrimeroYaEncolada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.Equal(t, 1, cola.Desencolar())
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}

func TestColaConUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaMantieneOrden(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.VerPrimero())
	cola.Encolar(5)
	require.Equal(t, 2, cola.VerPrimero())
	require.Equal(t, 2, cola.Desencolar())
	require.Equal(t, 3, cola.VerPrimero())
	require.Equal(t, 3, cola.Desencolar())
	require.Equal(t, 4, cola.VerPrimero())
	require.Equal(t, 4, cola.Desencolar())
	require.Equal(t, 5, cola.VerPrimero())
	require.Equal(t, 5, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}

func TestColaCadenas(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Hola")
	require.Equal(t, "Hola", cola.VerPrimero())
	cola.Encolar("como")
	cola.Encolar("estas?")
	require.Equal(t, "Hola", cola.VerPrimero())
	require.Equal(t, "Hola", cola.Desencolar())
	require.Equal(t, "como", cola.VerPrimero())
	require.Equal(t, "como", cola.Desencolar())
	require.Equal(t, "estas?", cola.VerPrimero())
	require.Equal(t, "estas?", cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaCaracter(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[rune]()
	cola.Encolar('H')
	cola.Encolar('o')
	cola.Encolar('l')
	cola.Encolar('a')
	require.Equal(t, 'H', cola.VerPrimero())
	require.Equal(t, 'H', cola.Desencolar())
	require.Equal(t, 'o', cola.VerPrimero())
	require.Equal(t, 'o', cola.Desencolar())
	require.Equal(t, 'l', cola.VerPrimero())
	require.Equal(t, 'l', cola.Desencolar())
	require.Equal(t, 'a', cola.VerPrimero())
	require.Equal(t, 'a', cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaFlotantes(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	cola.Encolar(1.5)
	cola.Encolar(2.5)
	cola.Encolar(3.5)
	cola.Encolar(4.5)
	require.Equal(t, 1.5, cola.VerPrimero())
	require.Equal(t, 1.5, cola.Desencolar())
	require.Equal(t, 2.5, cola.VerPrimero())
	require.Equal(t, 2.5, cola.Desencolar())
	require.Equal(t, 3.5, cola.VerPrimero())
	require.Equal(t, 3.5, cola.Desencolar())
	require.Equal(t, 4.5, cola.VerPrimero())
	require.Equal(t, 4.5, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaConVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	const n = 2000
	for i := 0; i < n; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for i := 0; i < n; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
		if i < n-1 {
			require.Equal(t, i+1, cola.VerPrimero())
		}
	}
	require.True(t, cola.EstaVacia())
}
