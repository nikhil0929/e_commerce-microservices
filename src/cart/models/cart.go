package models

// We are now going to make the cart microservice to be stateful. I am going to have a table in the database

/*
- create a new Cart microservice, which would be responsible for managing cart items for each user. This service would have its own database, where it stores each user's current cart items.

- Each cart could be identified by a unique cart ID, and each cart could have one or many cart items. Each cart item would include product ID and quantity. This allows you to track the cart state over time and make it easily recoverable.
*/

// CartItem represents a single item in the cart
type CartItem struct {
	ID         uint `gorm:"primary_key"`
	CartID     uint
	ProductID  uint
	TotalPrice float64
	Quantity   int
}

// Cart represents the cart model with multiple cart items
type Cart struct {
	ID        uint       `gorm:"primary_key"`
	UserID    uint       `gorm:"unique"`
	CartItems []CartItem `gorm:"foreignKey:CartID"`
}

// I am NOT using gorm.Model here because I do not want to have the CreatedAt, UpdatedAt, DeletedAt fields in the database. The cart is not as important as the user, so I do not need to track when the cart was created, updated, or deleted. I just need to track the cart ID, user ID, and cart items.
