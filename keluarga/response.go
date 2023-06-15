package keluarga

type KeluargaResponse struct {
	IdFrm                  string `json:"id_frm"`
	IdProvinsi             int    `json:"id_provinsi"`
	KodeProvinsi           int    `json:"kode_depdagri_provinsi"`
	NamaProvinsi           string `json:"nama_provinsi"`
	IdKabupaten            int    `json:"id_kabupaten"`
	KodeKabupaten          int    `json:"kode_depdagri_kabupaten"`
	NamaKabupaten          string `json:"nama_kabupaten"`
	IdKecamatan            int    `json:"id_kecamatan"`
	KodeKecamatan          int    `json:"kode_depdagri_kecamatan"`
	NamaKecamatan          string `json:"nama_kecamatan"`
	IdKelurahan            int    `json:"id_kelurahan"`
	KodeKelurahan          int    `json:"kode_depdagri_kelurahan"`
	NamaKelurahan          string `json:"nama_kelurahan"`
	KodeRw                 string `json:"kode_depdagri_rw"`
	NamaRw                 string `json:"nama_rw"`
	KodeRt                 string `json:"kode_depdagri_rt"`
	NamaRt                 string `json:"nama_rt"`
	Periode                int    `json:"periode"`
	KodeKeluarga           string `json:"kode_keluarga"`
	Nik                    string `json:"nik"`
	NamaKepalaKeluarga     string `json:"nama_kepala_keluarga"`
	NamaIstri              string `json:"nama_istri"`
	Baduta                 string `json:"baduta"`
	Balita                 string `json:"balita"`
	Pus                    string `json:"pus"`
	PusHamil               string `json:"pus_hamil"`
	SumberAirLayakTidak    string `json:"sumber_air_layak_tidak"`
	JambanLayakTidak       string `json:"jamban_layak_tidak"`
	TerlaluMuda            string `json:"terlalu_muda"`
	TerlaluTua             string `json:"terlalu_tua"`
	TerlaluDekat           string `json:"terlalu_dekat"`
	TerlaluBanyak          string `json:"terlalu_banyak"`
	BukanPesertaKbModern   string `json:"bukan_peserta_kb_modern"`
	KesejahteraanPrioritas int    `json:"kesejahteraan_prioritas"`
	RisikoStunting         string `json:"risiko_stunting"`
	Latitude               string `json:"latitude"`
	Longitude              string `json:"longitude"`
}
