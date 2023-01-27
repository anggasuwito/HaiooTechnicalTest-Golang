package models

func (Cart) TableName() string {
	return "cart"
}

type Cart struct {
	KodeProduk string `json:"kode_produk" gorm:"type:uuid;not null;primaryKey"`
	NamaProduk string `json:"nama_produk"`
	Kuantitas  int    `json:"kuantitas"`
}

type CartRequestBody struct {
	NamaProduk string `json:"nama_produk" binding:"required"`
	Kuantitas  int    `json:"kuantitas" binding:"required"`
}

type CartRequestParameter struct {
	NamaProduk string `form:"nama_produk"`
	Kuantitas  int    `form:"kuantitas"`
}
