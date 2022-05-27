package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User2 struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func main() {
	r := gin.Default()
	dsn := "root:153359157aA@tcp(127.0.0.1:3306)/a?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&User2{})
	if err != nil {
		return
	}
	r.POST("/login", func(c *gin.Context) {
		var User1, User3 User2
		err := c.ShouldBind(&User1)
		if err != nil {
			log.Println(err.Error())
		}
		res := db.Where(&User2{
			Username: User1.Username,
			Password: User1.Password,
		}).First(&User3)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"status": "error",
			})
			db.Create(&User1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "successful",
			})
		}
	})
	r.POST("/register", func(c *gin.Context) {
		var User1, User3 User2
		err := c.ShouldBind(&User1)
		if err != nil {
			log.Println(err.Error())
		}
		res := db.Where(&User2{
			Username: User1.Username,
			Password: User1.Password,
		}).First(&User3)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"status": "successful",
			})
			db.Create(&User1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "have registered",
			})
			fmt.Println(res.RowsAffected)
		}
	})
	err = r.Run()
	if err != nil {
		return
	}
}
