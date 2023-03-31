package services

import (
	"e_commerce-microservices/src/users/models"
	"e_commerce-microservices/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type DAO interface {
	Query(map[string][]string) ([]models.User, bool)
	Create(models.User) bool
	Update(map[string][]string, models.User) bool
	Delete(map[string][]string) bool
}

type Service struct {
	userDao DAO
}

func NewUserService(userDao DAO) *Service {
	return &Service{
		userDao: userDao,
	}
}


func (s *Service) GetUserProfile(email string) (models.User, bool) {
	conditions := map[string][]string{
		"email": {email},
	}
	// Get User details from the database
	dbUser, err := s.GetUsers(conditions)
	return dbUser[0], err
}

func (s *Service) GetUsers(queryParams map[string][]string) ([]models.User, bool) {
	users, err := s.userDao.Query(queryParams)
	return users, err
}

// You shouldnt be able to create a user if the email already exists in the database
func (s *Service) CreateUser(User models.User) {
	User.Password = utils.GenerateHashPassword(User.Password)
	s.userDao.Create(User)
}

// HELPER FUNCTIONS \\

func (s *Service) ValidateUserCredentials(user models.User) (models.User, bool) {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	dbUser, _:= s.GetUsers(conditions)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser[0].Password), []byte(user.Password))
	if err != nil {
		return models.User{}, false
	}
	return dbUser[0], true
}

func (s *Service) CheckUserExists(user models.User) bool {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	fmt.Println(conditions)
	// Check if the user email exists in the database
	dbUser, _ := s.GetUsers(conditions)
	if len(dbUser) == 0 {
		return false
	}
	return true
}

func (s *Service) CheckFormValidity(user models.User) bool {
	if user.Email == "" || user.Password == "" || user.Name == "" || user.Username == "" {
		return false
	}
	return true
}

func (s *Service) CheckLoginFormValidity(user models.User) bool {
	if user.Email == "" || user.Password == "" {
		return false
	}
	return true
}