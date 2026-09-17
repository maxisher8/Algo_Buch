package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.VerTope() })
	require.Panics(t, func() { pila.Desapilar() })
}

func TestPilaDesapilarVerTopeVacio(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.VerTope() })
	require.Panics(t, func() { pila.Desapilar() })
}

func TestPilaRecienCreadaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
}

func TestPilaDesapilarVerTopeYaApilada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.Equal(t, 1, pila.Desapilar())
	require.Panics(t, func() { pila.VerTope() })
	require.Panics(t, func() { pila.Desapilar() })
}

func TestPilaConUnElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 1, pila.VerTope())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaMantieneOrden(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)

	require.False(t, pila.EstaVacia())
	require.Equal(t, 4, pila.VerTope())

	require.Equal(t, 4, pila.Desapilar())
	require.Equal(t, 3, pila.VerTope())

	pila.Apilar(4)

	require.Equal(t, 4, pila.VerTope())
	require.Equal(t, 4, pila.Desapilar())
	require.Equal(t, 3, pila.VerTope())

	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.VerTope())

	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.VerTope())

	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.VerTope() })
	require.Panics(t, func() { pila.Desapilar() })

}

func TestPilaCadenas(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola")

	require.Equal(t, "Hola", pila.VerTope())

	pila.Apilar("como")
	pila.Apilar("estas?")

	require.Equal(t, "estas?", pila.VerTope())
	require.Equal(t, "estas?", pila.Desapilar())
	require.Equal(t, "como", pila.VerTope())

	require.Equal(t, "como", pila.Desapilar())
	require.Equal(t, "Hola", pila.VerTope())

	require.Equal(t, "Hola", pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaCaracter(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[rune]()
	pila.Apilar('H')
	pila.Apilar('o')
	pila.Apilar('l')
	pila.Apilar('a')

	require.Equal(t, 'a', pila.VerTope())
	require.Equal(t, 'a', pila.Desapilar())

	require.Equal(t, 'l', pila.VerTope())
	require.Equal(t, 'l', pila.Desapilar())
	require.Equal(t, 'o', pila.VerTope())

	require.Equal(t, 'o', pila.Desapilar())
	require.Equal(t, 'H', pila.VerTope())

	require.Equal(t, 'H', pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaFlotantes(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	pila.Apilar(1.5)
	pila.Apilar(2.5)
	pila.Apilar(3.5)
	pila.Apilar(4.5)

	require.Equal(t, 4.5, pila.VerTope())
	require.Equal(t, 4.5, pila.Desapilar())

	require.Equal(t, 3.5, pila.VerTope())
	require.Equal(t, 3.5, pila.Desapilar())
	require.Equal(t, 2.5, pila.VerTope())

	require.Equal(t, 2.5, pila.Desapilar())
	require.Equal(t, 1.5, pila.VerTope())

	require.Equal(t, 1.5, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaConVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	const n = 2000
	for i := 0; i < n; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
		if i > 0 {
			require.Equal(t, i-1, pila.VerTope())
		}
	}
	require.True(t, pila.EstaVacia())
}
