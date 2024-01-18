package home

import (
	"01-Login/platform/database"
	"01-Login/platform/models"
	"net/http"
	"strconv"

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
	var phoneModel models.PhoneInput
	err := ctx.ShouldBindJSON(&phoneModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastDigit := phoneModel.Phone[len(phoneModel.Phone)-1]
	lastDigitInt, err := strconv.Atoi(string(lastDigit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if lastDigitInt%2 == 0 {
		phoneModel.IsOdd = 0
	} else {
		phoneModel.IsOdd = 1
	}

	result := db.Create(&phoneModel)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, phoneModel)
	return
}
