package controller

import (
	// "database/sql"
	"api_rpm/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func UpdateStatus(c *gin.Context) {
	var pickup model.DriverShipments

	no_resi := c.Param("no_resi")

	if err := c.ShouldBindJSON(&pickup); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&pickup).Where("no_resi = ?", no_resi).Updates(&pickup).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})
}

func UpdateCustomers(c *gin.Context) {
	var customer model.Customer

	id_customer := c.Param("id_customer")

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&customer).Where("id_customer = ?", id_customer).Updates(&customer).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})
}

func UpdateKota(c *gin.Context) {
	var kota model.DaftarKotaPengiriman

	id_kota := c.Param("id_kota")

	if err := c.ShouldBindJSON(&kota); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&kota).Where("id_kota = ?", id_kota).Updates(&kota).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})
}

func UpdateUser(c *gin.Context) {
	var user model.Users

	id := c.Param("id_user")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Password = string(bytes)

	if model.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})
}

// func FilterPickupByNoResi(c *gin.Context) {
// 	var listPickup []model.DriverShipments
// 	param := c.Param("param")
// 	if model.DB.Where("no_resi = ? OR nama_pengirim = ? OR status_pickup = ?", param, param, param).Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&listPickup).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal mencari filter pickup"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"pickup_list": listPickup})
// }

// func FilterPickupByDate(c *gin.Context) {
// 	var listPickup []model.DriverShipments
// 	dateStr := c.Param("date")
// 	date, err := time.Parse("02-01-2006", dateStr)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
// 		return
// 	}
// 	if model.DB.Where("tanggal_pengiriman = ?", date).Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&listPickup).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal mencari filter melalui tanggal"})
// 		return
// 	}

// 	// 	c.JSON(http.StatusOK, gin.H{"pickup_list": listPickup})
// }

// func FilterPickupByDateTest(c *gin.Context) {
// 	var listPickup []model.DriverShipments
// 	startdateStr := c.Param("startdate")
// 	enddateStr := c.Param("enddate")
// 	startdate, err := time.Parse("02-01-2006", startdateStr)
// 	enddate, err := time.Parse("02-01-2006", enddateStr)
// 	fmt.Printf(startdate.String())
// 	fmt.Printf(enddate.String())
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
// 		return
// 	}
// 	if model.DB.Where("tanggal_pengiriman BETWEEN ? AND ?", startdateStr, enddateStr).Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&listPickup).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak ditemukan"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"pickup_list": listPickup})
// }