package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	/////
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"./avl"
	"./lista"
	"bytes"
	//"reflect"
	//"os"
	//"os/exec"
	//"strconv"
)

// json de Tiendas
type json_file struct {
	Datos []struct {
		Indice        string `json:"Indice"`
		Departamentos []struct {
			Nombre  string `json:"Nombre"`
			Tiendas []struct {
				Nombre       string `json:"Nombre"`
				Descripcion  string `json:"Descripcion"`
				Contacto     string `json:"Contacto"`
				Calificacion int    `json:"Calificacion"`
				Logo         string `json:"Logo"`
			} `json:"Tiendas"`
		} `json:"Departamentos"`
	} `json:"Datos"`
}

//json de Inventarios
type json_inventario struct {
	Invetarios []struct {
		Tienda       string `json:"Tienda"`
		Departamento string `json:"Departamento"`
		Calificacion int    `json:"Calificacion"`
		Productos    []struct {
			Nombre      string  `json:"Nombre"`
			Codigo      int     `json:"Codigo"`
			Descripcion string  `json:"Descripcion"`
			Precio      float64 `json:"Precio"`
			Cantidad    int     `json:"Cantidad"`
			Imagen      string  `json:"Imagen"`
		} `json:"Productos"`
	} `json:"Invetarios"`
}

type json_prueba struct {
	Datos []struct {
		Indice        string `json:"Indice"`
		Departamentos []struct {
			Nombre  string `json:"Nombre"`
			Tiendas []struct {
				Nombre       string `json:"Nombre"`
				Descripcion  string `json:"Descripcion"`
				Contacto     string `json:"Contacto"`
				Calificacion int    `json:"Calificacion"`
				Logo         string `json:"Logo"`
			} `json:"Tiendas"`
		} `json:"Departamentos"`
	} `json:"Datos"`
}

//json de Inventarios2
type json_inventario2 struct {
	Invetarios []struct {
		Tienda       string `json:"Tienda"`
		Departamento string `json:"Departamento"`
		Calificacion int    `json:"Calificacion"`
		Productos    []struct {
			Nombre      string  `json:"Nombre"`
			Codigo      int     `json:"Codigo"`
			Descripcion string  `json:"Descripcion"`
			Precio      float64 `json:"Precio"`
			Cantidad    int     `json:"Cantidad"`
			Imagen      string  `json:"Imagen"`
		} `json:"Productos"`
	} `json:"Invetarios"`
}

type json_producto struct{
	Productos    []struct {
		Nombre      string  `json:"Nombre"`
		Codigo      int     `json:"Codigo"`
		Descripcion string  `json:"Descripcion"`
		Precio      float64 `json:"Precio"`
		Cantidad    int     `json:"Cantidad"`
		Imagen      string  `json:"Imagen"`
	} `json:"Productos"`
}

type borrar struct {
	Invetarios []struct {
		Tienda       string `json:"Tienda"`
	} `json:"Invetarios"`
}

var dato_json json_file
var dato_inventario json_inventario
var dato_prueba json_prueba
var dato_inventario2 json_inventario2
var dato_producto json_producto
var dato_borrar borrar
var Invent = avl.NewAVL()
var Tienda_list = lista.NewLista()

func lista_tiendas(){
	for i:=0; i < len(dato_json.Datos); i++{
		var departamentos = dato_json.Datos[i].Departamentos
		//fmt.Println(departamentos)
		for j:=0; j < len(departamentos); j++{
			var tiendas = departamentos[j].Tiendas
			//fmt.Println(tiendas)
			for k:=0; k < len(tiendas); k++{
				var tienda = tiendas[k]
				//fmt.Println(tienda)
				nuevaTienda := lista.Tienda{
					Nombre:       tienda.Nombre,
					Descripcion:  tienda.Descripcion,
					Contacto:     tienda.Contacto,
					Calificacion: tienda.Calificacion,
					Logo:         tienda.Logo,
				}
				Tienda_list.Insertar(nuevaTienda)
			}
		}
	}
}

func avl_inventario(){
	for i:=0; i < len(dato_inventario.Invetarios); i++ {
		var productos = dato_inventario.Invetarios[i].Productos
		//fmt.Println(tiendas)
		for j:=0; j < len(productos); j++{
			var product = productos[j]
			//fmt.Println(product)
			nuevoInventario := avl.Producto{
				Nombre:      product.Nombre, 
				Codigo:		 product.Codigo,      
				Descripcion: product.Descripcion,
				Precio:      product.Precio,      
				Cantidad:    product.Cantidad,
				Imagen:      product.Imagen,
			}
			Invent.Insertar(nuevoInventario)
		}
	}
}

func getTiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_json)
}

func createTiendas(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_json)
	lista_tiendas()
	Tienda_list.Imprimir()
}

func getInventarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_inventario)
}

func createInventarios(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_inventario)
	avl_inventario()
	Invent.Print()

	//Mandando y regresando inventario

	jsonReq, err := json.Marshal(dato_inventario)
	req, err := http.Post("http://localhost:3000/Inventarios2", "application/json; charset = utf-8", bytes.NewBuffer(jsonReq))
	
	if err != nil {
	}

	defer req.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(bodyBytes, &dato_inventario)
	//fmt.Println(string(bodyBytes))
	//fmt.Fprintf(w, string(bodyBytes))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dato_inventario)
}

func getCargaProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_prueba)
}

func createCargaProduct(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_prueba)
}

func getInventario2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_inventario2)
}

func createInventario2(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_inventario2)
}

func getProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_producto)
}

func createProducto(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_producto)
}

func getBorrar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_borrar)
}

func createBorrar(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_borrar)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/Tiendas", getTiendas).Methods("GET")
	router.HandleFunc("/Tiendas", createTiendas).Methods("POST")
	router.HandleFunc("/Inventarios", getInventarios).Methods("GET")
	router.HandleFunc("/Inventarios", createInventarios).Methods("POST")
	router.HandleFunc("/cargaproduct", getCargaProduct).Methods("GET")
	router.HandleFunc("/cargaproduct", createCargaProduct).Methods("POST")
	router.HandleFunc("/Inventarios2", getInventario2).Methods("GET")
	router.HandleFunc("/Inventarios2", createInventario2).Methods("POST")
	router.HandleFunc("/producto", getProducto).Methods("GET")
	router.HandleFunc("/producto", createProducto).Methods("POST")
	router.HandleFunc("/Borrar", getBorrar).Methods("GET")
	router.HandleFunc("/Borrar", createBorrar).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	log.Println("Escuchando en http://localhost:3000")
	http.ListenAndServe(":3000", router)
}