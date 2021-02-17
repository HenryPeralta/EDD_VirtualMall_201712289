package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
	fmt.Println("hola")
	request()
	fmt.Println(re.Datos[0].Indice)
	fmt.Println("hola")
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
	fmt.Println(len(re.Datos))
	fmt.Println(re.Datos[1].Indice)
	fmt.Println(re.Datos[0].Departamentos)
	fmt.Println(re.Datos[0].Departamentos[0].Tiendas)
	fmt.Println(re.Datos[0].Departamentos[0].Tiendas[0].Calificacion)
	fmt.Println(len(re.Datos[1].Departamentos))
	//var i int
	//fmt.Println(i)
	//var j int = 6 
	//var a [16]int
	//fmt.Println(a)
	i := len(re.Datos) * len(re.Datos[0].Departamentos) * 5
	fmt.Println(i)
	//fmt.Println(j)
	arreglo := make([]int, i)
	fmt.Println(arreglo)
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

/*func mostrar(w http.ResponseWriter, r *http.Request){
	fmt.Println(re.Datos)
}*/

