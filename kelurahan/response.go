package kelurahan

type KelurahanResponse struct {
	IdProvinsi    int    `json:"id_provinsi"`
	KodeProvinsi  int    `json:"kode_depdagri_provinsi"`
	NamaProvinsi  string `json:"nama_provinsi"`
	IdKabupaten   int    `json:"id_kabupaten"`
	KodeKabupaten int    `json:"kode_depdagri_kabupaten"`
	NamaKabupaten string `json:"nama_kabupaten"`
	IdKecamatan   int    `json:"id_kecamatan"`
	KodeKecamatan int    `json:"kode_depdagri_kecamatan"`
	NamaKecamatan string `json:"nama_kecamatan"`
	IdKelurahan   int    `json:"id_kelurahan"`
	KodeKelurahan int    `json:"kode_depdagri_kelurahan"`
	NamaKelurahan string `json:"nama_kelurahan"`
}
