package kelurahan

import "encoding/json"

type WilayahInput struct {
	Nama   string      `json:"nama" binding:"required"`
	Kode   json.Number `json:"kode" binding:"required,number"`
	Alamat string      `json:"alamat" binding:"required"`
}
