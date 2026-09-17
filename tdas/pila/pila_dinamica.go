package pila

const (
	TAM_INICIAL        = 0
	PANIC_VACIO        = "La pila esta vacia"
	FACTOR_REDIMENSION = 2
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{make([]T, TAM_INICIAL), TAM_INICIAL}
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic(PANIC_VACIO)
	}
	return pila.datos[pila.cantidad-1]
}

// Apilar agrega un nuevo elemento a la pila.
func (pila *pilaDinamica[T]) Apilar(valorNuevo T) {
	if pila.cantidad == cap(pila.datos) {
		if cap(pila.datos) == 0 {
			pila.redimension(1)
		} else {
			pila.redimension(cap(pila.datos) * FACTOR_REDIMENSION)
		}
	}
	pila.datos[pila.cantidad] = valorNuevo
	pila.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic(PANIC_VACIO)
	}
	valor := pila.datos[pila.cantidad-1]
	pila.cantidad--
	if pila.cantidad < cap(pila.datos)/4 {
		pila.redimension(cap(pila.datos) / FACTOR_REDIMENSION)
	}
	return valor
}

func (pila *pilaDinamica[T]) redimension(tamanio int) {
	nuevapila := make([]T, tamanio)
	copy(nuevapila, pila.datos[:pila.cantidad])
	pila.datos = nuevapila
}
