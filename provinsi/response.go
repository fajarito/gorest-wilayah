package provinsi

type ProvinsiResponse struct {
	IdProvinsi   int    `json:"id_provinsi"`
	KodeProvinsi int    `json:"kode_depdagri_provinsi"`
	NamaProvinsi string `json:"nama_provinsi"`
}
