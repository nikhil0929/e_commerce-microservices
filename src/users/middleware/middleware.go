package middleware

import (
	"e_commerce-microservices/src/users/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorized(ctx *gin.Context) {

	// check if token is present in header
	sgToken := ctx.Request.Header["Token"]
	if len(sgToken) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		ctx.Abort()
		return
	}

	// validate token
	claims, isValid := auth.ValidateJWT(sgToken[0])
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}
	ctx.Request.Header.Set("email", claims.Email)
	ctx.Request.Header.Set("username", claims.Username)
	ctx.Request.Header.Set("role", claims.Role)
	ctx.Next()
}

// I NEED TO PUT THE 'ParseSession' MIDDLEWARE IN CART MICROSERVICE

// func ParseSession(ctx *gin.Context) {
// 	session := sessions.Default(ctx)
// 	var userCart models.Cart
// 	val := session.Get("cart")
// 	if val == nil {
// 		log.Println("Created new cart")
// 		userCart = models.Cart{}
// 	} else {
// 		log.Println("Found existing cart")
// 		userCart = val.(Models.Cart)
// 	}
// 	session.Set("cart", userCart)
// 	session.Save()
// 	ctx.Next()
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, mysession")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}