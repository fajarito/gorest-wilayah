package kelurahan

import (
	"gorm.io/gorm"
)

type Repository interface {
	// FindByID(nik string) (Stunting, error)
	// FindByPusHamil(nik string) (Stunting, error)
	// FindByStunting(nik string) (Stunting, error)
	// FindByBaduta(nik string) (Stunting, error)
	//KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krk, error)
	// GetKelurahanDetail(IdKelurahan int) ([]Kelurahan, error)
	GetKelurahan(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, page int, pageSize int) ([]Kelurahan, error)
	GetTotalKelurahanCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) (int64, error)

	// FindByKec(ProvinsiID int, KabupatenID int, KecamatanID int) ([]Krs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) GetKelurahanDetail(IdKelurahan string) (Kelurahan, error) {
// 	var kelurahan Kelurahan
// 	// err := r.db.Find(&kelurahan, TempatPelayananKBID).Error
// 	err := r.db.Where("id_kelurahan = ?", IdKelurahan).Find(&kelurahan).Error
// 	return kelurahan, err
// }

func (r *repository) GetKelurahan(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, page int, pageSize int) ([]Kelurahan, error) {
	var kelurahans []Kelurahan

	// err := r.db.Find(&faskeses, ProvinsiID).Error
	offset := (page - 1) * pageSize
	err := r.db.Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ?", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan).Offset(offset).Limit(pageSize).Find(&kelurahans).Error

	if err != nil {
		return nil, err
	}

	return kelurahans, nil
}

func (r *repository) GetTotalKelurahanCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) (int64, error) {
	var count int64
	if err := r.db.Model(&Kelurahan{}).Where("kode_depdagri_provinsi = ? AND kode_depdagri_kabupaten = ? AND kode_depdagri_kecamatan = ? ", KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
