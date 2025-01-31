package entitas

import "gorm.io/gorm"

type Asuransi struct {
	ID             int          `json:"id" gorm:"Column:id;primaryKey;autoIncrement"`
	NomorPolis     string       `json:"nomor_polis" gorm:"Column:nomor_polis"`
	NomorAplikasi  string       `json:"nomor_aplikasi" gorm:"Column:nomor_aplikasi"`
	Usia           int          `json:"usia" gorm:"Column:usia"`
	JangaWaktu     int          `json:"jangka_waktu" gorm:"Column:jangka_waktu"`
	Konstruksi     int          `json:"konstruksi" gorm:"Column:konstruksi"`
	Alamat         string       `json:"alamat" gorm:"Column:alamat"`
	Provinsi       string       `json:"provinsi" gorm:"Column:provinsi"`
	Kota           string       `json:"kota" gorm:"Column:kota"`
	Daerah         string       `json:"daerah" gorm:"Column:daerah"`
	IsGempaBumi    int          `json:"is_gempa_bumi" gorm:"Column:is_gempa_bumi"`
	Approve        int          `json:"approve" gorm:"Column:approve"`
	Status         string       `json:"status" gorm:"Column:status"`
	Premi          float64      `json:"premi" gorm:"Column:premi"`
	HargaBangunan  float64      `json:"harga_bangunan" gorm:"Column:harga_bangunan"`
	Total          float64      `json:"total" gorm:"Column:total"`
	UserID         int          `json:"user_id" gorm:"Column:user_id"`
	TipaBangunanID int          `json:"tipe_bangunan_id" gorm:"Column:tipe_bangunan_id"`
	User           User         `json:"author,omitempty" form:"author" gorm:"foreignkey:user_id"`
	TipeBangunan   TipeBangunan `json:"tipe_bangunan,omitempty" form:"tipe_bangunan" gorm:"foreignkey:tipe_bangunan_id"`
	Dihapus        gorm.DeletedAt
}
