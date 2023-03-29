package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e_commerce-microservices/src/users/models"
)

type Service interface {
	GetProducts(map[string][]string) ([]models.User, bool)
	CreateProduct(models.User) bool
	UpdateProduct(map[string][]string, models.User) bool
	DeleteProduct(map[string][]string) bool

	UserSignUp()
	UserSignIn()
	GetUserProfile()
	Logout()
}

type api struct {
	product_service Service
}

func NewProductController(product_service Service) *api {
	return &api{
		product_service: product_service,
	}
}

// Public level commands

/// Endpoint: /signup
/*
	POST: (Required: email, password, username, name. Optional: phone, address)
		- Check form validity and has required fields
		- Check if user already exists
		- Create user
		- Return success message to client
*/
func UserSignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil { // Bind client form data to user struct
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if Services.CheckFormValidity(user) { // Check form validity and has required fields
		if Services.CheckUserExists(user) { // Check if user already exists
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}
		Services.CreateUser(user) // Create user in DB
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
	}

}

// Endpoint: /login
/*
	POST: (email, password)
		- Check form validity and has required fields
		- Validate user credentials
		- Generate JWT token
		- Return token to client
*/
func UserSignIn(c *gin.Context) {
	var user models.User
	// Bind client form data to user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// MAYBE CLEAN UP THIS FUNCTION IT IS REALLY UGLY (DO IT ON ALL ENDPOINTS)
	if Services.CheckLoginFormValidity(user) { // Check form validity and has required fields
		dbUser, isValid := Services.ValidateUserCredentials(user) // Validate user credentials
		if isValid {
			token, isSuccess := Authenticator.GenerateJWT(dbUser) // Generate JWT token with required claims
			if !isSuccess {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"token": token}) // Return token to client
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
	}
}

// Endpoint: /profile
/*
	GET: (Required: token)
		- Request first goes to middleware to validate tokenf
		- If token is valid, parse claims
		- Return user profile to client
*/
func GetUserProfile(c *gin.Context) {
	// claims := c.MustGet("claims").(map[string]interface{})
	email := c.GetHeader("email")
	user := Services.GetUserProfile(email)
	c.JSON(http.StatusOK, user)
}

// TODO: NEED TO IMPLEMENT
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}