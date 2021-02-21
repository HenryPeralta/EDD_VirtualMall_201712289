package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"
	"./lista"
	"github.com/gorilla/mux"
	"strconv"
)

type Sobre struct {
	Datos []struct {
		Indice        string `json:"Indice"`
		Departamentos []struct {
			Nombre  string `json:"Nombre"`
			Tiendas []struct {
				Nombre       string `json:"Nombre"`
				Descripcion  string `json:"Descripcion"`
				Contacto     string `json:"Contacto"`
				Calificacion int    `json:"Calificacion"`
			} `json:"Tiendas"`
		} `json:"Departamentos"`
	} `json:"Datos"`
}

var dato_json Sobre
var filas int
var columnas int
var linealizada [] *lista.Lista



func matriz(){
	contFilas := 0
	
	for i := 0; i < len(dato_json.Datos); i++{
		var departamentos = dato_json.Datos[i].Departamentos

		var contColumnas int = 0
		
		for j := 0; j<len(departamentos); j++{
			var tiendas = departamentos[j].Tiendas
		
			var calificacion1 = lista.NewLista()
			var calificacion2 = lista.NewLista()
			var calificacion3 = lista.NewLista()
			var calificacion4 = lista.NewLista()
			var calificacion5 = lista.NewLista()

		
		 	for k := 0; k<len(tiendas); k++{
				var tienda = tiendas[k]
				nuevaTienda := lista.Tienda{
					Fila 		:contFilas,
					Columna 	:contColumnas,
					Calificacion:tienda.Calificacion,
					Nombre 		:tienda.Nombre,
					Descripcion :tienda.Descripcion,
					Contacto 	:tienda.Contacto,
					Departamento:departamentos[j].Nombre,
				}

				if(tiendas[k].Calificacion == 1){
					calificacion1.Insertar(nuevaTienda)
				}

				if(tiendas[k].Calificacion == 2){
					calificacion2.Insertar(nuevaTienda)
				}

				if(tiendas[k].Calificacion == 3){
					calificacion3.Insertar(nuevaTienda)
				}

				if(tiendas[k].Calificacion == 4){
					calificacion4.Insertar(nuevaTienda)
				}

				if(tiendas[k].Calificacion == 5){
					calificacion5.Insertar(nuevaTienda)
				}

			}

			for index :=0;index<len(linealizada);index++{
				if(linealizada[index] == nil){
					linealizada[index] = calificacion1
					linealizada[index + 1] = calificacion2
					linealizada[index + 2] = calificacion3
					linealizada[index + 3] = calificacion4
					linealizada[index + 4] = calificacion5
					break
				}
			}

			contColumnas++
		}
		contFilas++
	}
}


func main(){
	lis := lista.NewLista()

	lis.Imprimir()
	fmt.Println("Se termino de imprimir la lista")	
//	lis.Buscar(8)
//	lis.Buscar(1)
	fmt.Println("-----------------------")
//	lis.Eliminar(4)
	lis.Imprimir()
	request()
	fmt.Println(dato_json.Datos[0].Indice)
	fmt.Println("hola")
}

func servidor(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Bienvenido a mi Servidor")
}

func getArreglo(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "[1,2,3,4]")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_json)
}

func metodoPost(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_json)
	//fmt.Println(dato_json)
	//fmt.Fprint(w, dato_json)

	filas = len(dato_json.Datos)
	columnas = len(dato_json.Datos[0].Departamentos)
	linealizada = make([]*lista.Lista, filas * columnas * 5)
	matriz()

	generarReporte()

	// BUSQUEDA --------------------------------------
	// busqueda := buscar("Aurora","Deportes",7)
	// if busqueda {
	// 	fmt.Println("Se encontro");
	// } else {
	// 	fmt.Println("No lo encontre")
	// }

	// IMPRIMIR Y RECORRER ----------------------------------
	// fmt.Println(linealizada);
	// for index :=0;index<len(linealizada);index++{
	// 	listaTiendas := linealizada[index]
	// 	fmt.Println(index);

	// 	nodoLista := listaTiendas.GetInicio();
	// 	for nodoLista != nil {
	// 		tienda := nodoLista.GetTienda()
	// 		fmt.Println("		" + tienda.Nombre + " - " + tienda.Contacto)
	// 		nodoLista = nodoLista.GetSiguiente()
	// 	}	
	// }
}

func buscar(nombre string, departamento string, calificacion int) bool{ //lista.Tienda
	for index :=0;index<len(linealizada);index++{
		listaTiendas := linealizada[index]

		nodoLista := listaTiendas.GetInicio();
		for nodoLista != nil {
			tienda := nodoLista.GetTienda()
			fmt.Print(tienda.Nombre + " == " + tienda.Departamento + " == ")
			fmt.Print(tienda.Calificacion)
			fmt.Println(" == ")

			if tienda.Departamento == departamento {
				if tienda.Nombre == nombre {
					if tienda.Calificacion == calificacion {
						return true //return tienda
					}
				}
			}
			nodoLista = nodoLista.GetSiguiente()
		}	
	}
	return false //return nil
}

func generarReporte(){
    for index :=1 ;index<=len(linealizada);index++{
		listaTiendas := linealizada[index - 1]
		fmt.Println("	node"+ strconv.Itoa(index) +"[label=\""+ strconv.Itoa(index) +"\"];");
		
		if listaTiendas.Size() > 0 {
			fmt.Println("		node"+ strconv.Itoa(index) +"->node" + strconv.Itoa(index) + "0s;");
		}

		nodoLista := listaTiendas.GetInicio();
		j := 0
		for nodoLista != nil {
			tienda := nodoLista.GetTienda()
			
			fmt.Println("		  node"+ strconv.Itoa(index) + "" + strconv.Itoa(j) + "s[label=\""+ (tienda.Nombre) +"\"];");
			if j != listaTiendas.Size() - 1 {
				fmt.Println("		    node"+strconv.Itoa(index)+ "" + strconv.Itoa(j) +"s->node"+ strconv.Itoa(index) + "" + strconv.Itoa(j + 1) + "s;");
			}
			
			nodoLista = nodoLista.GetSiguiente()
			j++;
		}	
	}
}


func request(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", servidor)
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/getArreglo", metodoPost).Methods("POST")
	//router.HandleFunc("/mostrar", mostrar)
	log.Println("Escuchando en http://localhost:3000")
	http.ListenAndServe(":3000", router)
}


