package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	/////
	"encoding/json"
	//"fmt"
	"io/ioutil"
	//"reflect"
	//"os"
	//"os/exec"
	//"strconv"
)

/*type json_file struct {
	Name    string `json:"Name"`
	Content string `json:"Content"`
	Type    string `json:"Type"`
	Respuesta[]string `json:"Respuesta"`
	Traduccion string `json:"Traduccion"`
}*/

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

var dato_json json_file

func getTiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato_json)
}

func createTiendas(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato_json)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/Tiendas", getTiendas).Methods("GET")
	router.HandleFunc("/Tiendas", createTiendas).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	log.Println("Escuchando en http://localhost:3000")
	http.ListenAndServe(":3000", router)
}