package model

type DriverShipments struct {
	NoResi              string `gorm:"PrimaryKey" json:"no_resi"`
	NamaPengirim        string `gorm:"type:text" json:"nama_pengirim"`
	AlamatPengirim      string `gorm:"type:varchar(255)" json:"alamat_pengirim"`
	NoHandphonePengirim string `gorm:"type:varchar(100)" json:"no_handphone_pengirim"`
	AsalPengiriman      string `gorm:"type:text" json:"asal_pengiriman"`
	TujuanPengiriman    string `gorm:"type:text" json:"tujuan_pengiriman"`
	NamaPenerima        string `gorm:"type:text" json:"nama_penerima"`
	AlamatPenerima      string `gorm:"type:varchar(255)" json:"alamat_penerima"`
	NoHandphonePenerima string `gorm:"type:varchar(100)" json:"no_handphone_penerima"`
	IsiBarang           string `gorm:"type:varchar(50)" json:"isi_barang"`
	BeratBarang         int64  `gorm:"type:int" json:"berat_barang"`
	VolumeBarang        int64  `gorm:"type:int" json:"volume_barang"`
	JumlahKoli          int64  `gorm:"type:int" json:"jumlah_koli"`
	JenisLayanan        string `gorm:"type:text" json:"jenis_layanan"`
	NamaDriver          string `gorm:"type:text" json:"nama_driver"`
	TanggalPengiriman   string `gorm:"type:varchar(100)" json:"tanggal_pengiriman"`
	TotalHarga          int64  `gorm:"type:int" json:"total_harga"`
	StatusPickup        string `gorm:"type:text" json:"status_pickup"`
	CreatedAt           string `gorm:"type:varchar(100)" json:"created_at"`
	ChangeDate          string `gorm:"type:varchar(100)" json:"change_date"`
}
