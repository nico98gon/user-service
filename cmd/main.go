package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Servicio de usuarios en funcionamiento!")
	})

	http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://notification-service:8082/")
		if err != nil {
			http.Error(w, "Error al conectar con notification-service: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "Respuesta de notification-service: %s", body)
	})

	fmt.Println("Servicio de usuarios escuchando en el puerto 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
