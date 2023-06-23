package kecamatan

// Buat Struct nya terlebih dahulu

type Kecamatan struct {
	// gorm.Model

	IdProvinsi    int64  `gorm:"column:id_provinsi"`
	KodeProvinsi  int64  `gorm:"column:kode_depdagri_provinsi"`
	NamaProvinsi  string `gorm:"column:nama_provinsi"`
	IdKabupaten   int64  `gorm:"column:id_kabupaten"`
	KodeKabupaten int64  `gorm:"column:kode_depdagri_kabupaten"`
	NamaKabupaten string `gorm:"column:nama_kabupaten"`
	IdKecamatan   int64  `gorm:"column:id_kecamatan"`
	KodeKecamatan int64  `gorm:"column:kode_depdagri_kecamatan"`
	NamaKecamatan string `gorm:"column:nama_kecamatan"`
}

func (Kecamatan) TableName() string {
	return "sch_bkkbn.view_wilayah_kecamatan_2022_krs"
}
