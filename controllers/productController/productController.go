package productController

import (
	"encoding/json"
	"net/http"

	"github.com/RyanRamadhan11/Shop_RestApi_App/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var existingProduct models.Product
	id := c.Param("id")

	// Ambil data produk yang sudah ada dari database
	if err := models.DB.First(&existingProduct, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Produk tidak ditemukan"})
		return
	}

	// Bind data JSON ke struktur yang sudah ada
	if err := c.ShouldBindJSON(&existingProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Lakukan pembaruan pada produk yang sudah ada
	if err := models.DB.Save(&existingProduct).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal menyimpan perubahan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": existingProduct})
}



func Delete(c *gin.Context) {

	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}