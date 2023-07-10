package model

type Customer struct {
	IdCustomer     int64  `gorm:"PrimaryKey" json:"id_customer"`
	NamaCustomer   string `gorm:"type:text" json:"nama_customer"`
	AlamatCustomer string `gorm:"type:text" json:"alamat_customer"`
	NoHandphone    string `gorm:"type:varchar(100)" json:"no_handphone"`
	CreatedAt      string `gorm:"type:varchar(100)" json:"created_at"`
	ChangeDate     string `gorm:"type:varchar(100)" json:"change_date"`
}

func (Customer) TableName() string {
	return "customers"
}
