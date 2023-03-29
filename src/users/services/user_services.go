package services

import (
	"e_commerce-microservices/src/users/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type DAO interface {
	Query(map[string][]string) ([]models.Product, bool)
	Create(models.Product) bool
	Update(map[string][]string, models.Product) bool
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


func GetUserProfile(email string) models.User {
	conditions := map[string][]string{
		"email": {email},
	}
	// Get User details from the database
	dbUser := GetUsers(conditions)
	return dbUser[0]
}

func GetUsers(queryParams map[string][]string) []models.User {
	var RecievedUsers []models.User
	RecievedUsers = DB.QueryRecordWithMapConditions(&models.User{}, RecievedUsers, queryParams).([]models.User)
	return RecievedUsers
}

func CreateUser(User models.User) {
	User.Password = Utils.GenerateHashPassword(User.Password)
	DB.CreateRecord(&User)
}

// HELPER FUNCTIONS \\

func ValidateUserCredentials(user models.User) (models.User, bool) {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	dbUser := GetUsers(conditions)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser[0].Password), []byte(user.Password))
	if err != nil {
		return models.User{}, false
	}
	return dbUser[0], true
}

func CheckUserExists(user models.User) bool {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	fmt.Println(conditions)
	// Check if the user email exists in the database
	dbUser := GetUsers(conditions)
	if len(dbUser) == 0 {
		return false
	}
	return true
}

func CheckFormValidity(user models.User) bool {
	if user.Email == "" || user.Password == "" || user.Name == "" || user.Username == "" {
		return false
	}
	return true
}

func CheckLoginFormValidity(user models.User) bool {
	if user.Email == "" || user.Password == "" {
		return false
	}
	return true
}