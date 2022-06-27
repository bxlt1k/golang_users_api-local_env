package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"users_api/internal/dto"
	"users_api/internal/errors"
	"users_api/internal/models"
	"users_api/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetUsers(UserPage int) ([]*models.User, *errors.ApiError) {
	users, err := us.UserRepository.GetUsers(int64(UserPage))
	if err != nil {
		log.Println("unable to get users: " + err.Error())
		return nil, errors.InternalServerError(err)
	}

	return users, nil
}

func (us *UserService) GetUserByID(UserID int) (*models.User, *errors.ApiError) {
	user, err := us.UserRepository.GetUserByID(int64(UserID))
	if err != nil {
		log.Println("unable to get user by id: " + err.Error())
		return nil, errors.InternalServerError(err)
	}

	return user, nil
}

func (us *UserService) CreateUser(userData *dto.UserRegistrationData) (int, *errors.ApiError) {
	candidate, err := us.UserRepository.GetUserByEmail(userData.Email)
	if err != nil {
		log.Println("unable to get user by email: " + err.Error())
		return -1, errors.InternalServerError(err)
	}

	if candidate != nil {
		return -1, errors.BadRequestError(fmt.Sprintf("Пользователь с электронной почтой %s уже существует", userData.Email),
			fmt.Errorf("the user with the email %s already exists", userData.Email))
	}

	if userData.Password == "" {
		return -1, errors.BadRequestError(fmt.Sprintf("Укажите пароль"), fmt.Errorf("empty password"))
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 3)
	if err != nil {
		log.Println("unable to hash password: " + err.Error())
		return -1, errors.InternalServerError(err)
	}

	user := &models.User{
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Password:  string(hashPassword),
	}

	id, err := us.UserRepository.SaveUser(user)
	if err != nil {
		log.Println("unable to save user: " + err.Error())
		return -1, errors.InternalServerError(err)
	}

	return int(id), nil

}

func (us *UserService) DeleteUserByID(UserID int) (*models.User, *errors.ApiError) {
	candidate, err := us.UserRepository.GetUserByID(int64(UserID))
	if err != nil {
		log.Println("unable to get user by id: " + err.Error())
		return nil, errors.InternalServerError(err)
	}

	if candidate == nil {
		return nil, errors.BadRequestError(fmt.Sprintf("Пользователь с ID %d удален или еще не создан", UserID),
			fmt.Errorf("the user with the ID %d deleted or not created", UserID))
	}

	user, err := us.UserRepository.DeleteUserByID(int64(UserID))
	if err != nil {
		log.Println("unable to get users: " + err.Error())
		return nil, errors.InternalServerError(err)
	}

	return user, nil
}

func (us *UserService) UpdateUser(UserID int, userData *dto.UserUpdateData) *errors.ApiError {
	candidate, err := us.UserRepository.GetUserByID(int64(UserID))
	if err != nil {
		log.Println("unable to get user by id: " + err.Error())
		return errors.InternalServerError(err)
	}

	if candidate == nil {
		return errors.BadRequestError(fmt.Sprintf("Пользователь с ID %d удален или еще не создан", UserID),
			fmt.Errorf("the user with the ID %d deleted or not created", UserID))
	}

	candidate, err = us.UserRepository.GetUserByEmail(userData.Email)
	if err != nil {
		log.Println("unable to get user by email: " + err.Error())
		return errors.InternalServerError(err)
	}

	if candidate != nil {
		return errors.BadRequestError(fmt.Sprintf("Пользователь с электронной почтой %s уже существует", userData.Email),
			fmt.Errorf("the user with the email %s already exists", userData.Email))
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 3)
	if err != nil {
		log.Println("unable to hash password: " + err.Error())
		return errors.InternalServerError(err)
	}

	user := &models.User{
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Password:  string(hashPassword),
	}

	err = us.UserRepository.UpdateUser(int64(UserID), user)
	log.Println(err)
	if err != nil {
		log.Println("unable to update user: " + err.Error())
		return errors.InternalServerError(err)
	}

	return nil
}
