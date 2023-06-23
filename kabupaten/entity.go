package kabupaten

// Buat Struct nya terlebih dahulu

type Kabupaten struct {
	// gorm.Model

	IdProvinsi    int64  `gorm:"column:id_provinsi"`
	KodeProvinsi  int64  `gorm:"column:kode_depdagri_provinsi"`
	NamaProvinsi  string `gorm:"column:nama_provinsi"`
	IdKabupaten   int64  `gorm:"column:id_kabupaten"`
	KodeKabupaten int64  `gorm:"column:kode_depdagri_kabupaten"`
	NamaKabupaten string `gorm:"column:nama_kabupaten"`
}

func (Kabupaten) TableName() string {
	return "sch_bkkbn.view_wilayah_kabupaten_2022_krs"
}
