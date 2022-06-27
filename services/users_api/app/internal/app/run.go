package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"users_api/internal/controllers"
	"users_api/internal/repository"
	"users_api/internal/service"
	"users_api/internal/storage/mysql"
)

var (
	port         = os.Getenv("API_PORT")
	mySqlConnStr = os.Getenv("MYSQL_CONN_STR")
)

func Run() error {
	router := mux.NewRouter()

	mySqlConn, err := mysql.CreateConnection(mySqlConnStr)
	if err != nil {
		return err
	}

	userRepository := repository.NewUserRepository(mySqlConn)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", userController.DeleteUserByID).Methods("DELETE")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", userController.UpdateUserByID).Methods("PUT")

	log.Println("Users api server started on port " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		return err
	}
	return nil
}
