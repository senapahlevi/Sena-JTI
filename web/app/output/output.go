package output

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
	ctx.HTML(http.StatusOK, "output-page.html", nil)
}
func EditOutput(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "output-edit.html", nil)
}

func GetAllDataOutput(ctx *gin.Context) {
	var phoneModel []models.PhoneInput

	result := db.Find(&phoneModel)
	if result.Error != nil {
		fmt.Println("hello create 4", result.Error.Error())

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, phoneModel)
	return
}
func GetDataOutputID(ctx *gin.Context) {
	var phoneModel []models.PhoneInput
	id := ctx.Param("id")

	result := db.Where("id = ?", id).First(&phoneModel)
	if result.Error != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, phoneModel)
	return
}
func DeleteOutput(ctx *gin.Context) {
	id := ctx.Param("id")

	var phoneModel []models.PhoneInput

	result := db.Where("is_odd = ?", id).First(&phoneModel)
	if result.Error != nil {
		fmt.Println("hello create 4", result.Error.Error())

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	result = db.Delete(&phoneModel)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, phoneModel)
	return
}
func UpdateDataOutput(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println("id", id)
	var phoneModel models.PhoneInput

	result := db.Model(&phoneModel).Where("id = ?", id).First(&phoneModel)
	if result.Error != nil {
		fmt.Println("hello create 4", result.Error.Error())

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	fmt.Println("hello create cobi", phoneModel)

	err := ctx.ShouldBindJSON(&phoneModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result = db.Save(&phoneModel)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	fmt.Println("hello create cobi2", result)

	ctx.JSON(http.StatusOK, phoneModel)
	return
}
