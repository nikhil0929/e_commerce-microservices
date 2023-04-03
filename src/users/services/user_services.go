package services

import (
	"e_commerce-microservices/src/users/models"
	"e_commerce-microservices/utils"
	"fmt"
	"log"

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
	if err == false {
		log.Println("Error getting user from database")
		return models.User{}, false
	}
	return dbUser[0], true
}

func (s *Service) GetUsers(queryParams map[string][]string) ([]models.User, bool) {
	users, _ := s.userDao.Query(queryParams)
	// log.Println(users, err)
	if len(users) == 0 {
		return []models.User{}, false
	}
	return users, true
}

// You shouldnt be able to create a user if the email already exists in the database
func (s *Service) CreateUser(User models.User) {
	User.Password = utils.GenerateHashPassword(User.Password)
	s.userDao.Create(User)
}

// HELPER FUNCTIONS \\

func (s *Service) ValidateUserCredentials(user models.User) (models.User, bool) {
	// conditions := map[string][]string{
	// 	"email": {user.Email},
	// }
	dbUser, exists := s.GetUserProfile(user.Email)
	if exists == false {
		return models.User{}, false
	}
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return models.User{}, false
	}
	return dbUser, true
}

// Do I really need this function? I can just use the GetUserProfile function which returns a bool
// Refactor this later
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