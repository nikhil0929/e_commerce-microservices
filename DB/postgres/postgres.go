package postgres

// All Database related operations in this directory (DB)
// ex: FetchFromDB, DeleteFromDB
// all database functions here can be used by 'Services' to fetch/update/insert/delete data from DB
// *** place DB connection file in this folder as well ***
import (
	"fmt"
	"log"

	"e_commerce-microservices/config"
	"e_commerce-microservices/utils" // THIS IS THE CORRECT WAY TO IMPORT PACKAGES

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database connection object
type DB struct {
	connection *gorm.DB
}

func newDBConnection_postgres() *DB {
	return &DB{
		connection: ConnectDatabase(),
	}
}

// Creates a new database connection with the given credentials
// Returns DB object
func ConnectDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.Db_username, config.Db_password, config.Db_name, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")
	return db
}

// Creates new record instance in the database from the given model object
func (pdb *DB) CreateRecord(object interface{}) bool {

	result := pdb.connection.Create(object)
	if result.Error != nil {
		// panic(result.Error)
		return false
	}
	log.Println(utils.CreateLogMessage("Created record", object))
	return true
}

// Queries the database for the given set of fields and some string conditions specified as a map/struct
// Returns the queried object
func (pdb *DB) QueryRecordWithMapConditions(modelObject interface{}, outputObject interface{}, conditions interface{}) (interface{}, bool) {
	filterQuerable := pdb.connection.Model(&modelObject)
	newQuery := utils.CreateConditionClause(filterQuerable, conditions.(map[string][]string))
	result := newQuery.Find(&outputObject)
	if result.Error != nil {
		// panic(result.Error)
		return nil, false
	}
	// save outputObject in an array variable

	log.Println(utils.CreateLogMessage("Queried record", modelObject))
	return outputObject, true
}

// Updates the given model object in the database with new fields specified as a map/struct
func (pdb *DB) UpdateRecord(modelObject interface{}, conditions interface{}, newVals interface{}) bool{
	filterQuerable := pdb.connection.Model(&modelObject)
	newQuery := utils.CreateConditionClause(filterQuerable, conditions.(map[string][]string))
	result := newQuery.Updates(newVals)
	if result.Error != nil {
		// panic(result.Error)
		return false
	}
	log.Println(utils.CreateLogMessage("Updated record", modelObject))
	return true
}

// Deletes the given model object from the database with the given conditions specified as a map
func (pdb *DB) DeleteRecord(modelObject interface{}, conditions interface{}) bool{
	filterQuerable := pdb.connection.Model(&modelObject)
	newQuery := utils.CreateConditionClause(filterQuerable, conditions.(map[string][]string))
	result := newQuery.Delete(&modelObject)
	if result.Error != nil {
		// panic(result.Error)
		return false
	}
	log.Println(utils.CreateLogMessage("Deleted record", modelObject))
	return true
}