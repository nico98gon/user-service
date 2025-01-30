package main

import (
	"fmt"
	"io"
	"net/http"
	"nilus-challenge-backend/internal/infrastructure"

	"os"
)

func main() {
	db := infrastructure.NewDBConnection()
	defer db.Close()

	infrastructure.NewRouter(db)

	notificationServicePort := os.Getenv("NOTIFICATION_SERVICE_PORT")
	var noficiationServiceURL string
	if notificationServicePort == "" {
		noficiationServiceURL = "http://notification-service:8082"
	} else {
		noficiationServiceURL = fmt.Sprintf("http://notification-service:%s", notificationServicePort)
	}

	fmt.Println("Servicio de usuarios escuchando en el puerto 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}

	http.HandleFunc("/notification", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(noficiationServiceURL)
		if err != nil {
			http.Error(w, "Error al conectar con notification-service: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "Respuesta de notification-service: %s", body)
	})
}
