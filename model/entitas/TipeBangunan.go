package entitas

import "gorm.io/gorm"

type TipeBangunan struct {
	ID           int     `json:"id" gorm:"Column:id;primaryKey;autoIncrement"`
	KodeBangunan string  `json:"kode_bangunan" gorm:"Column:kode_bangunan;unique"`
	NamaBangunan string  `json:"nama_bangunan" gorm:"Column:nama_bangunan"`
	Tarif        float64 `json:"tarif" gorm:"Column:tarif"`
	Dihapus      gorm.DeletedAt
}

type RequestAddTipeBangunan struct {
	KodeBangunan string  `json:"kode_bangunan"`
	NamaBangunan string  `json:"nama_bangunan"`
	Tarif        float64 `json:"tarif"`
}
