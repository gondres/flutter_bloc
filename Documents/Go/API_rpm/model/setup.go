package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/shipment"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Users{})
	database.AutoMigrate(&Customer{})
	database.AutoMigrate(&DriverShipments{})
	database.AutoMigrate(&BarangShipment{})
	database.AutoMigrate(&DaftarKotaPengiriman{})
	
	DB = database
}
