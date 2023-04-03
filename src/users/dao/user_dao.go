package dao

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/src/users/models"
)

type UserDao struct{
	connection *postgres.DB
}

func NewUserDao(connection *postgres.DB) *UserDao {
	return &UserDao{
		connection: connection,
	}
}

// Query returns all users in the DB with specified query parameters
func (us *UserDao) Query(queryParams map[string][]string) ([]models.User, bool) {
	var RecievedUsers []models.User
	rp, err := us.connection.QueryRecordWithMapConditions(&models.User{}, RecievedUsers, queryParams)
	//log.Println("rp: ", rp, "err: ", err)
	RecievedUsers = rp.([]models.User)
	return RecievedUsers, err
}

// Create creates a new user in the DB with specified query parameters and new fields
func (us *UserDao) Create(newFields models.User) bool {
	return us.connection.CreateRecord(&newFields)
}

// Update updates a user in the DB with specified query parameters and new fields
func (us *UserDao) Update(conditions map[string][]string, newFields models.User) bool {
	return us.connection.UpdateRecord(models.User{}, conditions, newFields)
}

// DeleteProduct deletes a user in the DB with specified query parameters
func (us *UserDao) Delete(conditions map[string][]string) bool {
	return us.connection.DeleteRecord(models.User{}, conditions)
}