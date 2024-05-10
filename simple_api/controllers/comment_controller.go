package controllers

import (
	"example/simple_api/config"
	"os"

	"example/simple_api/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/google/uuid"
	"gorm.io/gorm"
)

func All(c *gin.Context) {

	var comments []entity.Comment
	config.DB.Find(&comments)
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func Index(c *gin.Context) {
	var comment entity.Comment
	id := c.Param("id")

	if err := config.DB.First(&comment, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func Create(c *gin.Context) {
	var comment entity.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Create(&comment).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat insert new data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Input data berhasil"})
}

func Update(c *gin.Context) {
	var comment entity.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id := comment.Id
	if config.DB.Model(&comment).Where("id = ?", id).Updates(&comment).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat update data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ubah data berhasil"})
}

func Delete(c *gin.Context) {
	var comment entity.Comment
	id := c.Param("id")
	if config.DB.Delete(&comment, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak dapat dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil hapus data"})
}

func Upload(c *gin.Context) {
	data, errReq := c.MultipartForm()
	if errReq != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": errReq.Error()})
	}

	file := data.File["file"][0]
	errSaveFile := c.SaveUploadedFile(file, "./image/"+file.Filename)

	if errSaveFile != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": errSaveFile.Error()})
	}

	_, _ = os.Open("./image/" + file.Filename)
	c.JSON(http.StatusOK, gin.H{"data": "./image/" + file.Filename})
}
