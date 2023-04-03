package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e_commerce-microservices/src/users/auth"
	"e_commerce-microservices/src/users/models"
)

type Service interface {
	GetUserProfile(string) (models.User, bool)
	GetUsers(queryParams map[string][]string) ([]models.User, bool)
	CreateUser(User models.User)
	ValidateUserCredentials(user models.User) (models.User, bool)
	CheckUserExists(user models.User) bool
	CheckFormValidity(user models.User) bool
	CheckLoginFormValidity(user models.User) bool
}

type api struct {
	user_service Service
}

func NewUserController(user_service Service) *api {
	return &api{
		user_service: user_service,
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
func (a *api) UserSignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil { // Bind client form data to user struct
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if a.user_service.CheckFormValidity(user) { // Check form validity and has required fields
		if a.user_service.CheckUserExists(user) { // Check if user already exists
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}
		a.user_service.CreateUser(user) // Create user in DB
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
func (a *api) UserSignIn(c *gin.Context) {
	var user models.User
	// Bind client form data to user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// MAYBE CLEAN UP THIS FUNCTION IT IS REALLY UGLY (DO IT ON ALL ENDPOINTS)
	if a.user_service.CheckLoginFormValidity(user) { // Check form validity and has required fields
		dbUser, isValid := a.user_service.ValidateUserCredentials(user) // Validate user credentials
		if isValid {
			token, isSuccess := auth.GenerateJWT(dbUser) // Generate JWT token with required claims
			if !isSuccess {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"token": token}) // Return token to client
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
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
func (a *api) GetUserProfile(c *gin.Context) {
	// claims := c.MustGet("claims").(map[string]interface{})
	email := c.GetHeader("email")
	user, isValid := a.user_service.GetUserProfile(email)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// TODO: NEED TO IMPLEMENT
func (a *api) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}