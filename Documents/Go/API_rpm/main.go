package main

import (
	"api_rpm/controller"
	"api_rpm/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	model.ConnectDatabase()

	r.POST("/shipment/login", controller.LoginHandler)
	r.POST("/shipment/register", controller.HashAccount)
	r.POST("/shipment/pickup", controller.ValidateJWT(controller.CreatePickup))
	r.POST("/shipment/barang", controller.ValidateJWT(controller.CreateBarang))
	r.POST("/shipment/customer", controller.ValidateJWT(controller.CreateCustomer))
	r.POST("/shipment/kota", controller.ValidateJWT(controller.CreateKotaPengiriman))

	r.PUT("/shipment/update/:no_resi", controller.UpdateStatus)
	r.PUT("/shipment/update/customer/:id_customer", controller.UpdateCustomers)
	r.PUT("/shipment/update/kota/:id_kota", controller.UpdateKota)
	r.PUT("/shipment/update/user/:id_user", controller.UpdateUser)

	r.GET("/shipment/pickup/list", controller.ValidateJWT(controller.GetPickupList))
	r.GET("/shipment/barang/:resi_barang", controller.ValidateJWT(controller.GetPhotoByResi))
	r.GET("/shipment/users", controller.GetUsers)
	r.GET("/shipment/customer", controller.ValidateJWT(controller.GetCustomerList))
	r.GET("/shipment/home", controller.ValidateJWT(controller.GetUsers))
	r.GET("/shipment/kota/list", controller.ValidateJWT(controller.GetDaftarKota))
	r.GET("/shipment/user/:username", controller.ValidateJWT(controller.GetUserByUsername))
	r.GET("/shipment/filter/:param", controller.FilterPickupByNoResi)
	r.GET("/shipment/filter/date/:startdate/:enddate", controller.FilterPickupByDateTest)

	r.DELETE("/shipment/delete/user/:id_user", controller.DeleteUser)
	r.DELETE("/shipment/delete/customer/:id_customer", controller.DeleteCustomer)
	r.DELETE("/shipment/delete/kota/:id_kota", controller.DeleteKota)
	r.Run()
}