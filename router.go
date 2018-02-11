// router.go - sets up the routes for this file.
package main

import (
	"github.com/betacraft/yaag/yaag"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gin-gonic/gin.v1"
)

// Router - contains the router for the file
func Router() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Authentication", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})

	router := gin.Default()

	v1 := router.Group("/auth")
	{
		v1.GET("/user", BrowseUsers)
		v1.GET("/user/:id", ReadUser)
		v1.PUT("/user/:id", EditUser)
		v1.POST("/user", AddUser)
		v1.DELETE("user/:id", DeleteUser)
	}
	router.Run()
}
