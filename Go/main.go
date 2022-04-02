package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><meta http-equiv='X-UA-Compatible' content='IE=edge'><meta name='viewport' content='width=device-width, initial-scale=1.0'><title>Document</title></head><body style='background-color: red;text-align:center;border-style:double;border-radius:90px; padding:10px;'><h1>Mi formulario</h1>Ingresa tu nombre: <input type='text'><br>Ingresa tu edad: <input type='number'><br><input type='submit' value='Enviar'></body></html>")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Go web App Stored on port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
