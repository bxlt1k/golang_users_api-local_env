package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"users_api/internal/dto"
	"users_api/internal/helpers"
	"users_api/internal/service"
)

var (
	clientURL = os.Getenv("CLIENT_URL")
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	users, er := uc.UserService.GetUsers(page)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	encode, err := json.Marshal(users)
	if err != nil {
		log.Println("unable to encode users: " + err.Error())
		helpers.ErrorResponse(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(encode)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	UserId, err := strconv.Atoi(vars["id"])

	users, er := uc.UserService.GetUserByID(UserId)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	encode, err := json.Marshal(users)
	if err != nil {
		log.Println("unable to encode users: " + err.Error())
		helpers.ErrorResponse(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(encode)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("unable to read request body: " + err.Error())
		helpers.ErrorResponse(w, "Некорректный запрос", http.StatusInternalServerError)
		return
	}

	var userRegistrationData *dto.UserRegistrationData
	if err = json.Unmarshal(body, &userRegistrationData); err != nil {
		log.Println("unable to decode request body: " + err.Error())
		helpers.ErrorResponse(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	userID, er := uc.UserService.CreateUser(userRegistrationData)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprintf(
		`{"id": "%d", "firstName": "%s", "lastName": "%s"}`,
		userID, userRegistrationData.FirstName, userRegistrationData.LastName)))
}

func (uc *UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	UserId, err := strconv.Atoi(vars["id"])

	users, er := uc.UserService.DeleteUserByID(UserId)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	encode, err := json.Marshal(users)
	if err != nil {
		log.Println("unable to encode users: " + err.Error())
		helpers.ErrorResponse(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write(encode)
}

func (uc *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	UserId, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("unable to read request body: " + err.Error())
		helpers.ErrorResponse(w, "Некорректный запрос", http.StatusInternalServerError)
		return
	}

	var userUpdateData *dto.UserUpdateData
	if err = json.Unmarshal(body, &userUpdateData); err != nil {
		log.Println("unable to decode request body: " + err.Error())
		helpers.ErrorResponse(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	er := uc.UserService.UpdateUser(UserId, userUpdateData)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
