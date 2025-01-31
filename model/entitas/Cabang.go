package entitas

import "gorm.io/gorm"

type Cabang struct {
	ID         int    `json:"id" gorm:"Column:id;primaryKey;autoIncrement"`
	KodeCabang string `json:"kode_cabang" gorm:"Column:kode_cabang;unique"`
	Nama       string `json:"nama" gorm:"Column:nama"`
	Dihapus    gorm.DeletedAt
}
