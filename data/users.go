package data

import "fmt"

// ErrUserNotFound is an error raised when a product can not be found in the database
var ErrUserNotFound = fmt.Errorf("Product not found")

// User defines the structure for an API User
// swagger:model
type User struct {
	// the id of the user
	//
	// required: false
	// min:1
	Id int `json:id`

	// the name of the user
	//
	// required: true
	// max length: 255
	Name string `json:"name"`

	// the email of the user
	//
	// required: true
	// max length: 255
	Email string `json:"email"`

	// the password of the user
	//
	// required: true
	// max length: 255
	Password string `json:"password"`

	// the id of the group that the user belongs to
	//
	// required: true
	// min: 1
	GroupId int `json:"groupId"`
}

// Users defines a slice of Users
type Users []*User

// GetUsers returns all products from the database
func GetUsers() Users {

}

// GetUserById returns a single user with the specified id
// If the user is not found this func retuns UserNotFound error
func GetUserById(id int) (*User, error) {

}

// UpdateUser replaces a user with the given item
// If the user is not found this func returns UserNotFound error
func UpdateUser(u User) error {

}

// AddUser adds a new user to the database
func AddUser(u User) {

}

// DeleteUser deletes an user from the database
func DeleteUser(id int) error {

}