package model

type DaftarKotaPengiriman struct {
	IdKota       int64  `gorm:"PrimaryKey" json:"id_kota"`
	NamaKota     string `gorm:"type:text" json:"nama_kota"`
	Harga        int64  `gorm:"type:int" json:"harga"`
	BeratMinimum int64  `gorm:"type:int" json:"berat_minimum"`
}

func (DaftarKotaPengiriman) TableName() string {
	return "daftar_kota_pengiriman"
}
