package keluarga

// Buat Struct nya terlebih dahulu

type Keluarga struct {
	// gorm.Model

	// IdFrm                  string `gorm:"column:id_frm"`
	IdFrm                  string `gorm:"column:id_frm"`
	IdProvinsi             int64  `gorm:"column:id_provinsi"`
	KodeProvinsi           int64  `gorm:"column:kode_depdagri_provinsi"`
	NamaProvinsi           string `gorm:"column:nama_provinsi"`
	IdKabupaten            int64  `gorm:"column:id_kabupaten"`
	KodeKabupaten          int64  `gorm:"column:kode_depdagri_kabupaten"`
	NamaKabupaten          string `gorm:"column:nama_kabupaten"`
	IdKecamatan            int64  `gorm:"column:id_kecamatan"`
	KodeKecamatan          int64  `gorm:"column:kode_depdagri_kecamatan"`
	NamaKecamatan          string `gorm:"column:nama_kecamatan"`
	IdKelurahan            int64  `gorm:"column:id_kelurahan"`
	KodeKelurahan          int64  `gorm:"column:kode_depdagri_kelurahan"`
	NamaKelurahan          string `gorm:"column:nama_kelurahan"`
	KodeRw                 string `gorm:"column:kode_depdagri_rw"`
	NamaRw                 string `gorm:"column:nama_rw"`
	KodeRt                 string `gorm:"column:kode_depdagri_rt"`
	NamaRt                 string `gorm:"column:nama_rt"`
	Periode                int64  `gorm:"column:periode"`
	KodeKeluarga           string `gorm:"column:kode_keluarga"`
	Nik                    string `gorm:"column:nik"`
	NamaKepalaKeluarga     string `gorm:"column:nama_kepala_keluarga"`
	NamaIstri              string `gorm:"column:nama_istri"`
	Baduta                 string `gorm:"column:baduta"`
	Balita                 string `gorm:"column:balita"`
	Pus                    string `gorm:"column:pus"`
	PusHamil               string `gorm:"column:pus_hamil"`
	SumberAirLayakTidak    string `gorm:"column:sumber_air_layak_tidak"`
	JambanLayakTidak       string `gorm:"column:jamban_layak_tidak"`
	TerlaluMuda            string `gorm:"column:terlalu_muda"`
	TerlaluTua             string `gorm:"column:terlalu_tua"`
	TerlaluDekat           string `gorm:"column:terlalu_dekat"`
	TerlaluBanyak          string `gorm:"column:terlalu_banyak"`
	BukanPesertaKbModern   string `gorm:"column:bukan_peserta_kb_modern"`
	KesejahteraanPrioritas int64  `gorm:"column:kesejahteraan_prioritas"`
	RisikoStunting         string `gorm:"column:risiko_stunting"`
	Latitude               string `gorm:"column:latitude"`
	Longitude              string `gorm:"column:longitude"`
}

func (Keluarga) TableName() string {
	return "sch_bkkbn.view_keluarga_2022_krs"
}
