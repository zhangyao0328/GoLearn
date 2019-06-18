package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/gin-gonic/gin"
)

var db *gorm.DB
var err error

type Person struct {
	ID uint `json:”id”`
	FirstName string `json:”firstname”`
	LastName string `json:”lastname”`
}

func main() {

	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/", GetProjects)
	r.GET("/people/:id", GetPerson)

	r.Run(":8080")
}

func GetProjects(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
	c.JSON(200, person)
	}
}
