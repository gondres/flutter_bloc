package model

type BarangShipment struct {
	IdBarang      string `gorm:"PrimaryKey" json:"id_barang"`
	ResiBarang    string `gorm:"type:varchar(255)" json:"resi_barang"`
	NamaBarang    string `gorm:"type:varchar(255)" json:"nama_barang"`
	PemilikBarang string `gorm:"type:text" json:"pemilik_barang"`
	ImageUrl      string `gorm:"type:varchar(255)" json:"image_url"`
	CreatedAt     string `jgorm:"type:varchar(100)" son:"created_at"`
	ChangeDate    string `gorm:"type:varchar(100)" json:"change_date"`
}

func (BarangShipment) TableName() string {
	return "barang_shipment"
}
