package lista

import (
	"fmt"
)

//lugar donde se almacena la informacion
type nodo struct {
	anterior  *nodo
	siguiente *nodo
	tienda      Tienda
}

func (n *nodo) GetSiguiente() *nodo {
	return n.siguiente
}

func (n *nodo) GetTienda() Tienda {
	return n.tienda
}

type Tienda struct {
	Fila         int
	Columna      int
	Calificacion int
	Nombre       string
	Descripcion  string
	Contacto     string
	Departamento string
}

//estructura para almacenar nodos de informacion
type Lista struct {
	inicio *nodo
	ultimo *nodo
	tamaño int
}

func (m *Lista) GetInicio() *nodo {
	return m.inicio
}

func (m *Lista) Size() int {
	return m.tamaño
}

//crear una nueva lista
func NewLista() *Lista {
	return &Lista{nil, nil, 0}
}

//insertar un nodo
func (m *Lista) Insertar(valor Tienda) {
	nuevo := &nodo{nil, nil, valor}

	if m.inicio == nil {
		m.inicio = nuevo
		m.ultimo = nuevo
	} else {
		m.ultimo.siguiente = nuevo
		nuevo.anterior = m.ultimo
		m.ultimo = nuevo
	}
	m.tamaño++
}

//imprimir la lista
func (m *Lista) Imprimir() {
	aux := m.inicio
	for aux != nil {
		fmt.Print("<-[", aux.tienda, "]->")
		aux = aux.siguiente
	}
	fmt.Println()
	fmt.Println("tamaño de la lista = ", m.tamaño)
}

//buscar elementos dentro de una lista
func (m *Lista) Buscar(nombre string, departamento string, calificacion int) *nodo {
	aux := m.inicio
	for aux != nil {
		if aux.tienda.Nombre == nombre {
			fmt.Println(aux.tienda.Nombre)
			if aux.tienda.Departamento == departamento {
				if aux.tienda.Calificacion == calificacion {
					fmt.Println("si se encontro el nodo")
					return aux
				}
				aux = aux.siguiente
			}
			aux = aux.siguiente
		}
		aux = aux.siguiente
	}
	fmt.Println("no se encontro el nodo")
	return nil
}

//eliminar nodo de la lista
/*func (m *Lista) Eliminar(valor Tiendas){
	aux := m.Buscar(valor)

	if m.inicio == aux{
		m.inicio = aux.siguiente
		aux.siguiente.anterior = nil
		aux.siguiente = nil
	}else if m.ultimo == aux{
		m.ultimo = aux.anterior
		aux.anterior.siguiente = nil
		aux.anterior = nil
	}else{
		aux.anterior.siguiente = aux.siguiente
		aux.siguiente.anterior = aux.anterior
		aux.anterior = nil
		aux.siguiente = nil
	}
	m.tamaño --
}*/