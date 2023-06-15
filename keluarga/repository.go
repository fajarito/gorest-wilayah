package keluarga

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	// FindByID(nik string) (Stunting, error)
	// FindByPusHamil(nik string) (Stunting, error)
	// FindByStunting(nik string) (Stunting, error)
	// FindByBaduta(nik string) (Stunting, error)
	//KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krk, error)
	GetKeluarga(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, page int, pageSize int) ([]Keluarga, error)
	GetTotalKeluargaCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) (int64, error)
	GetKeluargaBeresiko(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int, page int, pageSize int) ([]Keluarga, error)
	GetTotalKeluargaBeresikoCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int) (int64, error)

	// FindByKec(ProvinsiID int, KabupatenID int, KecamatanID int) ([]Krs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetKeluarga(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, page int, pageSize int) ([]Keluarga, error) {
	var keluargas []Keluarga

	// err := r.db.Find(&faskeses, ProvinsiID).Error
	offset := (page - 1) * pageSize
	err := r.db.Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ? AND kode_depdagri_kelurahan = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan).Offset(offset).Limit(pageSize).Find(&keluargas).Error

	if err != nil {
		return nil, err
	}

	return keluargas, nil
}

func (r *repository) GetTotalKeluargaCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) (int64, error) {
	var count int64
	if err := r.db.Model(&Keluarga{}).Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ? AND kode_depdagri_kelurahan = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) GetKeluargaBeresiko(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int, page int, pageSize int) ([]Keluarga, error) {
	var keluargas []Keluarga

	// err := r.db.Find(&faskeses, ProvinsiID).Error
	// offset := (page - 1) * pageSize
	// err := r.db.Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ? AND kode_depdagri_kelurahan = ? AND risiko_stunting= 'V'", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan).Offset(offset).Limit(pageSize).Find(&keluargas).Error

	// if err != nil {
	// 	return nil, err
	// }

	// return keluargas, nil
	offset := (page - 1) * pageSize
	query := r.db.Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ? AND kode_depdagri_kelurahan = ? AND risiko_stunting= 'V'", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan)

	if filter2 != 0 {
		allowedFilter2Values := []int{1, 2, 3, 4, 0}
		isFilter2Valid := false

		for _, value := range allowedFilter2Values {
			if filter2 == value {
				isFilter2Valid = true
				break
			}
		}

		if !isFilter2Valid {
			return nil, errors.New("Invalid filter2 value. Must be one of: 1, 2, 3, 4, 0")
		}

		query = query.Where("kesejahteraan_prioritas = ?", filter2)
	}

	err := query.Offset(offset).Limit(pageSize).Find(&keluargas).Error
	if err != nil {
		return nil, err
	}

	return keluargas, nil
}

// func (r *repository) GetTotalKeluargaBeresikoCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) (int64, error) {
// 	var count int64
// 	if err := r.db.Model(&Keluarga{}).Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ? AND kode_depdagri_kelurahan = ? AND risiko_stunting= 'V'", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan).Count(&count).Error; err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

func (r *repository) GetTotalKeluargaBeresikoCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int) (int64, error) {
	var count int64
	query := r.db.Model(&Keluarga{}).Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ? AND kode_depdagri_kelurahan = ? AND risiko_stunting = 'V'", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan)

	if filter2 != 0 {
		query = query.Where("kesejahteraan_prioritas = ?", filter2)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
