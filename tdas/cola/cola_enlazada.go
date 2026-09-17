package cola

const (
	_PANIC_VACIO = "La cola esta vacia"
)

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

func crearNodoCola[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato: dato, siguiente: nil}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{nil, nil}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic(_PANIC_VACIO)
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(valorNuevo T) {
	nuevoNodo := crearNodoCola(valorNuevo)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
		cola.ultimo = nuevoNodo
	} else {
		cola.ultimo.siguiente = nuevoNodo
		cola.ultimo = nuevoNodo
	}
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic(_PANIC_VACIO)
	}
	valor := cola.primero.dato
	cola.primero = cola.primero.siguiente
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return valor
}
