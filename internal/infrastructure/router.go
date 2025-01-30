package infrastructure

import (
	"database/sql"
	"fmt"
	"net/http"
	"nilus-challenge-backend/internal/domain/user"
	handler "nilus-challenge-backend/internal/infrastructure/http"
	user_repo "nilus-challenge-backend/internal/infrastructure/repository/user"
)

func NewRouter(db *sql.DB) {
	userPostgresRepo := user_repo.NewPostgresUserRepository(db)

	fmt.Println(userPostgresRepo)

	userService := user.NewService(userPostgresRepo)
	userHandler := handler.NewUserHandler(userService)

	api := "/api/v1"

	http.HandleFunc("GET "+api+"/users", userHandler.GetUsers)
	http.HandleFunc("GET "+api+"/users/{id}", userHandler.GetUserByID)
	http.HandleFunc("POST "+api+"/users", userHandler.CreateUser)
	http.HandleFunc("PUT "+api+"/users/{id}", userHandler.UpdateUser)
	http.HandleFunc("DELETE "+api+"/users/{id}", userHandler.DeleteUser)
	http.HandleFunc("PUT "+api+"/users/{id}/opt-out", userHandler.OptOutUser)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}