package dao

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/utils"
	"log"

	"gorm.io/gorm"
)

type CartDao struct {
	connection *gorm.DB
}

func NewCartDao() *CartDao {
	new_cart_dao := &CartDao{
		connection: postgres.ConnectDatabase(),
	}

	return new_cart_dao
}

func (cd *CartDao) RunMigrations(model interface{}) {
	utils.MigrateModel(cd.connection, model)
}

// Query returns all items in the DB with specified conditionals
func (cd *CartDao) QueryRecords(object interface{}, conditions map[string]interface{}) error {
	result := cd.connection.Where(conditions).Find(object)

	if result.Error != nil {
		return result.Error
	}

	log.Println("Queried DB successfully")

	return nil
}

func (cd *CartDao) QueryWithAssociation(object interface{}, conditions map[string]interface{}, association string) error {
	result := cd.connection.Preload(association).Where(conditions).Find(object)

	if result.Error != nil {
		return result.Error
	}

	log.Println("Queried DB with association successfully")

	return nil
}

// Creates a new model in the database from the given model object
func (cd *CartDao) CreateRecord(object interface{}) error {
	result := cd.connection.Create(object)

	if result.Error != nil {
		return result.Error
	}

	log.Println("Created new record in DB")

	return nil
}

// Update updates an item in the DB for the given model object with the given newFields
// NOTES:
//   - object is a pointer to the model object and must include the ID primary key from the DB
//   - newFields is a map of the form: map[string][]string{"field_name": {"new_value"}}
func (cd *CartDao) UpdateRecord(object interface{}, newFields map[string]interface{}) error {

	result := cd.connection.Model(object).Updates(newFields)

	if result.Error != nil {
		return result.Error
	}

	log.Println("Updated record in DB")

	return nil
}

// DeleteCart deletes an in the DB with specified conditions
// NOTES:
//   - object is a pointer to the model object and must include the ID primary key from the DB
//   - conditions is a map of the form: map[string]interface{}{"field_name": "value"}
func (cd *CartDao) DeleteRecord(object interface{}, conditions map[string]interface{}) error {
	result := cd.connection.Where(conditions).Delete(object)

	if result.Error != nil {
		return result.Error
	}

	log.Println("Deleted record in DB")

	return nil
}
