package controller

import (
	"api_rpm/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePickup(c *gin.Context) {
	var body model.DriverShipments

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := model.DB.Table("driver_shipments").Create(&body)
	if result.Error != nil {

		log.Printf("Error creating shipment: %v", result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}
	log.Printf("success menginput pickup")
	c.JSON(http.StatusOK, gin.H{"message": "success menginput pickup"})
}

func CreateBarang(c *gin.Context) {
	var body []model.BarangShipment

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := model.DB.Table("barang_shipment").Create(&body)
	if result.Error != nil {
		// panic(result.Error)
		log.Printf("Error creating shipment: %v", result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	log.Printf("success menginput barang pickup")
	c.JSON(http.StatusOK, gin.H{"message": "barang berhasil diproses"})
}

func CreateCustomer(c *gin.Context) {
	var body model.Customer

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := model.DB.Table("customers").Create(&body)
	if result.Error != nil {

		log.Printf("Error create new user: %v", result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	log.Printf("success menginput customer baru")
	c.JSON(http.StatusOK, gin.H{"message": "customer baru berhasil diinput."})
}

func CreateKotaPengiriman(c *gin.Context) {
	var body model.DaftarKotaPengiriman

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := model.DB.Table("daftar_kota_pengiriman").Create(&body)
	if result.Error != nil {

		log.Printf("Error create new kota pengiriman: %v", result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	log.Printf("success menginput kota pengiriman")
	c.JSON(http.StatusOK, gin.H{"message": "kota pengiriman baru berhasil diinput."})
}
