package controller

import (
	// "database/sql"
	"api_rpm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomerList(c *gin.Context) {
	var customer []model.Customer
	model.DB.Select("id_customer, nama_customer, alamat_customer, no_handphone, created_at,change_date").Find(&customer)

	c.JSON(http.StatusOK, gin.H{"customers": customer})

}

func GetUserByUsername(c *gin.Context) {
	var user []model.User
	username := c.Param("username")
	if model.DB.Where("username = ?", username).Select("id, username, name, role").Find(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal mencari user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetDaftarKota(c *gin.Context) {
	var daftarkota []model.DaftarKotaPengiriman
	model.DB.Select("id_kota, nama_kota, harga,berat_minimum").Find(&daftarkota)

	c.JSON(http.StatusOK, gin.H{"daftar_kota": daftarkota})

}

func GetUsers(c *gin.Context) {

	var users []model.Users

	if c.GetHeader("Access") == api_key && c.GetHeader("Content-Type") == content_type {
		model.DB.Select("id, username, name, role").Find(&users)
		c.JSON(http.StatusOK, gin.H{"users": users})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Header"})
	}

}

func GetPickupList(c *gin.Context) {

	var pickup []model.DriverShipments

	if c.GetHeader("Access") == api_key && c.GetHeader("Content-Type") == content_type {
		model.DB.Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&pickup)
		c.JSON(http.StatusOK, gin.H{"pickup_list": pickup})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Header"})
	}
	// model.DB.Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&pickup)
	// c.JSON(http.StatusOK, gin.H{"pickup_list": pickup})

}

func GetPhotoByResi(c *gin.Context) {
	var barang []model.BarangShipment
	resi_barang := c.Param("resi_barang")
	if model.DB.Where("resi_barang = ?", resi_barang).Select("nama_barang,pemilik_barang, image_url").Find(&barang).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal mencari gambar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"barang_pickup": barang})
}
