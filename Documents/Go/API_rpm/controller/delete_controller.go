package controller

import (
	// "database/sql"
	"api_rpm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id_user")

	// Perform deletion operation
	if result := model.DB.Delete(&model.User{}, id); result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id_customer")

	// Perform deletion operation
	if result := model.DB.Delete(&model.Customer{}, id); result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}

func DeleteKota(c *gin.Context) {
	id := c.Param("id_kota")

	// Perform deletion operation
	if result := model.DB.Delete(&model.DaftarKotaPengiriman{}, id); result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}