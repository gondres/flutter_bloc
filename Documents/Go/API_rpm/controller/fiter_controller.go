package controller

import (
	"api_rpm/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FilterPickupByNoResi(c *gin.Context) {
	var listPickup []model.DriverShipments
	param := c.Param("param")
	if model.DB.Where("no_resi = ? OR nama_pengirim = ? OR status_pickup = ?", param, param, param).Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&listPickup).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal mencari filter pickup"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pickup_list": listPickup})
}

func FilterPickupByDate(c *gin.Context) {
	var listPickup []model.DriverShipments
	dateStr := c.Param("date")
	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
		return
	}
	if model.DB.Where("tanggal_pengiriman = ?", date).Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&listPickup).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal mencari filter melalui tanggal"})
		return
	}

	// 	c.JSON(http.StatusOK, gin.H{"pickup_list": listPickup})
}

func FilterPickupByDateTest(c *gin.Context) {
	var listPickup []model.DriverShipments
	startdateStr := c.Param("startdate")
	enddateStr := c.Param("enddate")
	startdate, err := time.Parse("02-01-2006", startdateStr)
	enddate, err := time.Parse("02-01-2006", enddateStr)
	fmt.Printf(startdate.String())
	fmt.Printf(enddate.String())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
		return
	}
	if model.DB.Where("tanggal_pengiriman BETWEEN ? AND ?", startdateStr, enddateStr).Select("no_resi, nama_pengirim, alamat_pengirim,no_handphone_pengirim,asal_pengiriman, tujuan_pengiriman, nama_penerima,alamat_penerima, no_handphone_penerima, isi_barang, berat_barang, volume_barang, jumlah_koli, jenis_layanan,nama_driver, tanggal_pengiriman,total_harga, status_pickup, created_at, change_date").Find(&listPickup).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pickup_list": listPickup})
}
