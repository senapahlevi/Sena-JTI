package home

import (
	"01-Login/platform/database"
	"01-Login/platform/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler for our home page.
//
//	func Handler(ctx *gin.Context) {
//		ctx.HTML(http.StatusOK, "home.html", nil)
//	}
var db *gorm.DB

func SetDatabase(database *database.Database) {
	db = database.DB
}
func Handler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form-input.html", nil)
}

func Create(ctx *gin.Context) {
	fmt.Println("hello create", db)
	var phoneModel models.PhoneInput
	err := ctx.ShouldBindJSON(&phoneModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("hello create 2")

		return
	}
	fmt.Println("hello create 3")

	result := db.Create(&phoneModel)
	if result.Error != nil {
		fmt.Println("hello create 4", result.Error.Error())

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, phoneModel)
	return
}
