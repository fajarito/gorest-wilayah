package handler

import (
	"gorest-krs-wilayah/kelurahan"
	"math"
	"strconv"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type kelurahanHandler struct {
	kelurahanService kelurahan.Service
}

func NewKelurahanHandler(kelurahanService kelurahan.Service) *kelurahanHandler {
	return &kelurahanHandler{kelurahanService}
}

// ListKelurahan godoc
// @Summary Get Wilayah By Kelurahan
// @Description Get Wilayah Sasaran By Kelurahan
// @ID get-list-wilayah-by-kelurahan
// @Tags Get Wilayah By Kelurahan
// @Accept */*
// @Produce json
// @Success      200  {object}  kelurahan.Kelurahan
// @Param kdprov query string true "Province Code"
// @Param kdkab query string true "Kabupaten Code"
// @Param kdkec query string true "Kecamatan Code"
// @Param page query integer false "Page Number"
// @Param pageSize query integer false "Number of Kelurahan per page" default(10)
// @Router /showkelurahan [get]
func (handler *kelurahanHandler) GetKelurahan(c *gin.Context) {

	KodeDepdagriProvinsi := c.Query("kdprov")
	KodeDepdagriKabupaten := c.Query("kdkab")
	KodeDepdagriKecamatan := c.Query("kdkec")

	page := convertToInts(c.DefaultQuery("page", "1"), 1)
	pageSize := convertToInts(c.DefaultQuery("pageSize", "10"), 10)

	if pageSize <= 0 || pageSize > 50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pageSize value. Must be between 1 and 50",
		})
		return
	}

	kelurahans, err := handler.kelurahanService.GetKelurahan(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, page, pageSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	totalItems, err := handler.kelurahanService.GetTotalKelurahanCount(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var kelurahansResponse []kelurahan.KelurahanResponse

	for _, f := range kelurahans {
		kelurahanResponse := convertToKelurahanResponse(f)
		kelurahansResponse = append(kelurahansResponse, kelurahanResponse)
	}
	response := gin.H{
		"data": kelurahansResponse,
	}

	if totalItems > 0 {
		totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
		response["pagination"] = gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalPages":  totalPages,
			"totalItems":  totalItems,
		}
	}
	c.JSON(http.StatusOK, response)

}

func convertToInts(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}

func convertToInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return val
	}
	return val
}

func convertToKelurahanResponse(f kelurahan.Kelurahan) kelurahan.KelurahanResponse {
	return kelurahan.KelurahanResponse{

		IdProvinsi:    int(f.IdProvinsi),
		KodeProvinsi:  int(f.KodeProvinsi),
		NamaProvinsi:  f.NamaProvinsi,
		IdKabupaten:   int(f.IdKabupaten),
		KodeKabupaten: int(f.KodeKabupaten),
		NamaKabupaten: f.NamaKabupaten,
		IdKecamatan:   int(f.IdKecamatan),
		KodeKecamatan: int(f.KodeKecamatan),
		NamaKecamatan: f.NamaKecamatan,
		IdKelurahan:   int(f.IdKelurahan),
		KodeKelurahan: int(f.KodeKelurahan),
		NamaKelurahan: f.NamaKelurahan,
	}
}
