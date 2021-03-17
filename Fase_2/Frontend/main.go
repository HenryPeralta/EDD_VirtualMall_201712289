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

var dato_json json_file
var dato_inventario json_inventario
var Invent = avl.NewAVL()

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
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/Tiendas", getTiendas).Methods("GET")
	router.HandleFunc("/Tiendas", createTiendas).Methods("POST")
	router.HandleFunc("/Inventarios", getInventarios).Methods("GET")
	router.HandleFunc("/Inventarios", createInventarios).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	log.Println("Escuchando en http://localhost:3000")
	http.ListenAndServe(":3000", router)
}