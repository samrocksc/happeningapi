// handlers.go.go - Handles the BREAD operations of a user account
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gin-gonic/gin.v1"
)

// AddUser - Create A User.
func AddUser(c *gin.Context) {
	password := []byte(c.PostForm("password"))
	fmt.Println(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hashedPassword))
	users := Users{
		Email:    c.PostForm("email"),
		Password: string(hashedPassword),
		Username: c.PostForm("username"),
	}

	db := Database()
	db.NewRecord(&users)

	if err := db.Create(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User failed to create"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Default Return"})
	return
}

// BrowseUsers - browses for active users
func BrowseUsers(c *gin.Context) {
	var users []Users
	var _users []TransformedUsers

	db := Database()
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Users Found"})
		return
	}

	for _, item := range users {
		_users = append(_users, TransformedUsers{
			ID:          item.ID,
			IsValidated: item.IsValidated,
			Version:     item.Version,
			Email:       item.Email,
			Username:    item.Username,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _users})
}

// ReadUser - Retrieve Single User
func ReadUser(c *gin.Context) {
	var users Users
	userId := c.Param("id")

	db := Database()
	db.First(&users, userId)

	if users.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
	}

	_users := TransformedUsers{
		ID:          users.ID,
		IsValidated: users.IsValidated,
		Version:     users.Version,
		Email:       users.Email,
		Username:    users.Username,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _users})
}

// EditUser - Edits an active user
func EditUser(c *gin.Context) {
	var users Users
	userId := c.Param("id")
	db := Database()
	db.First(&users, userId)

	if users.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User Found!"})
		return
	}

	db.Model(&users).Where("id = ?", userId).Updates(Users{
		IsValidated: users.IsValidated,
		Version:     users.Version,
		Email:       users.Email,
		Username:    users.Username,
	})

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}

// DeleteUser - Deletes a user from the database
func DeleteUser(c *gin.Context) {
	var users Users
	usersId := c.Param("id")
	db := Database()
	db.First(&users, usersId)

	if users.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User Found!"})
		return
	}

	db.Delete(&users)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
