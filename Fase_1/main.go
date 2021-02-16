package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Tienda struct{
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
}

type Departamento struct{
	Nombre string
	Tiendas []Tienda
}

type Dato struct{
	Indice string
	Departamentos []Departamento 
}

type Sobre struct{
	Datos []Dato
}

var re Sobre

func main(){
	request()
}

func servidor(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Bienvenido a mi Servidor")
}

func getArreglo(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "[1,2,3,4]")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(re)
}

func metodoPost(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &re)
	fmt.Println(re)
	fmt.Fprint(w, re)
}

func request(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", servidor)
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/getArreglo", metodoPost).Methods("POST")
	log.Println("Escuchando en http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

