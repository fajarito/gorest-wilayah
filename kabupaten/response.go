package kabupaten

type KabupatenResponse struct {
	IdProvinsi    int    `json:"id_provinsi"`
	KodeProvinsi  int    `json:"kode_depdagri_provinsi"`
	NamaProvinsi  string `json:"nama_provinsi"`
	IdKabupaten   int    `json:"id_kabupaten"`
	KodeKabupaten int    `json:"kode_depdagri_kabupaten"`
	NamaKabupaten string `json:"nama_kabupaten"`
}
