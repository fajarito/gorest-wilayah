package handler

import (
	"gorest-krs-wilayah/provinsi"
	"math"
	"strconv"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type provinsiHandler struct {
	provinsiService provinsi.Service
}

func NewProvinsiHandler(provinsiService provinsi.Service) *provinsiHandler {
	return &provinsiHandler{provinsiService}
}

// ListProvinsi godoc
// @Summary Get Wilayah by Province
// @Description Get Wilayah by Province
// @ID get-wilayah-by-province
// @Tags Get Wilayah by Province
// @Accept */*
// @Produce json
// @Success      200  {object}  provinsi.Provinsi
// @Param pageSize query integer false "Number of Province per page" default(34)
// @Router /showprovinsi [get]
func (handler *provinsiHandler) GetProvinsi(c *gin.Context) {

	page := convertToIntss(c.DefaultQuery("page", "1"), 1)
	pageSize := convertToIntss(c.DefaultQuery("pageSize", "40"), 40)

	if pageSize <= 0 || pageSize > 40 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pageSize value. Must be between 1 and 50",
		})
		return
	}

	provinsis, err := handler.provinsiService.GetProvinsi(page, pageSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	totalItems, err := handler.provinsiService.GetTotalProvinsiCount()
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var provinsisResponse []provinsi.ProvinsiResponse

	for _, f := range provinsis {
		provinsiResponse := convertToProvinsiResponse(f)
		provinsisResponse = append(provinsisResponse, provinsiResponse)
	}
	response := gin.H{
		"data": provinsisResponse,
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

func convertToIntssss(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}

func convertToProvinsiResponse(f provinsi.Provinsi) provinsi.ProvinsiResponse {
	return provinsi.ProvinsiResponse{

		IdProvinsi:   int(f.IdProvinsi),
		KodeProvinsi: int(f.KodeProvinsi),
		NamaProvinsi: f.NamaProvinsi,
	}
}
