package handler

import (
	"gorest-krs-wilayah/kecamatan"
	"math"
	"strconv"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type kecamatanHandler struct {
	kecamatanService kecamatan.Service
}

func NewKecamatanHandler(kecamatanService kecamatan.Service) *kecamatanHandler {
	return &kecamatanHandler{kecamatanService}
}

// ListKecamatan godoc
// @Summary Get Wilayah By Kecamatan
// @Description Get Wilayah By Kecamatan
// @ID get-list-wilayah-by-kecamatan
// @Tags Get Wilayah By Kecamatan
// @Accept */*
// @Produce json
// @Success      200  {object}  kecamatan.Kecamatan
// @Param kdprov query string true "Province Code"
// @Param kdkab query string true "Kabupaten Code"
// @Param page query integer false "Page Number"
// @Param pageSize query integer false "Number of Kecamatan per page" default(10)
// @Router /showkecamatan [get]
func (handler *kecamatanHandler) GetKecamatan(c *gin.Context) {

	KodeDepdagriProvinsi := c.Query("kdprov")
	KodeDepdagriKabupaten := c.Query("kdkab")

	page := convertToIntss(c.DefaultQuery("page", "1"), 1)
	pageSize := convertToIntss(c.DefaultQuery("pageSize", "10"), 10)

	if pageSize <= 0 || pageSize > 50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pageSize value. Must be between 1 and 50",
		})
		return
	}

	kecamatans, err := handler.kecamatanService.GetKecamatan(KodeDepdagriProvinsi, KodeDepdagriKabupaten, page, pageSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	totalItems, err := handler.kecamatanService.GetTotalKecamatanCount(KodeDepdagriProvinsi, KodeDepdagriKabupaten)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var kecamatansResponse []kecamatan.KecamatanResponse

	for _, f := range kecamatans {
		kecamatanResponse := convertToKecamatanResponse(f)
		kecamatansResponse = append(kecamatansResponse, kecamatanResponse)
	}
	response := gin.H{
		"data": kecamatansResponse,
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

func convertToIntss(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}

func convertToKecamatanResponse(f kecamatan.Kecamatan) kecamatan.KecamatanResponse {
	return kecamatan.KecamatanResponse{

		IdProvinsi:    int(f.IdProvinsi),
		KodeProvinsi:  int(f.KodeProvinsi),
		NamaProvinsi:  f.NamaProvinsi,
		IdKabupaten:   int(f.IdKabupaten),
		KodeKabupaten: int(f.KodeKabupaten),
		NamaKabupaten: f.NamaKabupaten,
		IdKecamatan:   int(f.IdKecamatan),
		KodeKecamatan: int(f.KodeKecamatan),
		NamaKecamatan: f.NamaKecamatan,
	}
}
