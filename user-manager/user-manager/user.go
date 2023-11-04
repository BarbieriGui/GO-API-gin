package main

import (
	"errors"
	"net/http"

	//"github.com/mattn/go-sqlite3"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
}

var users = []user{ // need to implement password, role and uuid
	{Name: "John", Lastname: "Shepard", Email: "j.shepard@test.com", Phone: 91929394},
	{Name: "Marcos", Lastname: "Coisos", Email: "m.coisos@test.com", Phone: 94939291},
	{Name: "Richard", Lastname: "Jr", Email: "r.jr@test.com", Phone: 91949293},
}

// function to list users stored in user struct
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// function to add new users to the user struct
func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// functions to call a user by email
func userByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := getUserByEmail(email)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func getUserByEmail(email string) (*user, error) {
	for i, u := range users {
		if u.Email == email {
			return &users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

// function to edit the users information in the user struct
// func manageUser(c *gin.Context) {
// }
