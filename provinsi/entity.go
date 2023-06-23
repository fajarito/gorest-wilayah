package provinsi

// Buat Struct nya terlebih dahulu

type Provinsi struct {
	// gorm.Model

	IdProvinsi   int64  `gorm:"column:id_provinsi"`
	KodeProvinsi int64  `gorm:"column:kode_depdagri_provinsi"`
	NamaProvinsi string `gorm:"column:nama_provinsi"`
}

func (Provinsi) TableName() string {
	return "sch_bkkbn.view_wilayah_provinsi_2022_krs"
}
