package handler

import (
	"gorest-krs-wilayah/kabupaten"
	"math"
	"strconv"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type kabupatenHandler struct {
	kabupatenService kabupaten.Service
}

func NewKabupatenHandler(kabupatenService kabupaten.Service) *kabupatenHandler {
	return &kabupatenHandler{kabupatenService}
}

// ListKabupaten godoc
// @Summary Get Wilayah By Kabupaten
// @Description Get Wilayah By Kabupaten
// @ID get-list-wilayah-by-kabupaten
// @Tags Get Wilayah By Kabupaten
// @Accept */*
// @Produce json
// @Success      200  {object}  kabupaten.Kabupaten
// @Param kdprov query string true "Province Code"
// @Param page query integer false "Page Number"
// @Param pageSize query integer false "Number of Kabupaten per page" default(10)
// @Router /showkabupaten [get]
func (handler *kabupatenHandler) GetKabupaten(c *gin.Context) {

	KodeDepdagriProvinsi := c.Query("kdprov")

	page := convertToIntss(c.DefaultQuery("page", "1"), 1)
	pageSize := convertToIntss(c.DefaultQuery("pageSize", "10"), 10)

	if pageSize <= 0 || pageSize > 50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pageSize value. Must be between 1 and 50",
		})
		return
	}

	kabupatens, err := handler.kabupatenService.GetKabupaten(KodeDepdagriProvinsi, page, pageSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	totalItems, err := handler.kabupatenService.GetTotalKabupatenCount(KodeDepdagriProvinsi)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var kabupatensResponse []kabupaten.KabupatenResponse

	for _, f := range kabupatens {
		kabupatenResponse := convertToKabupatenResponse(f)
		kabupatensResponse = append(kabupatensResponse, kabupatenResponse)
	}
	response := gin.H{
		"data": kabupatensResponse,
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

func convertToIntsss(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}

func convertToKabupatenResponse(f kabupaten.Kabupaten) kabupaten.KabupatenResponse {
	return kabupaten.KabupatenResponse{

		IdProvinsi:    int(f.IdProvinsi),
		KodeProvinsi:  int(f.KodeProvinsi),
		NamaProvinsi:  f.NamaProvinsi,
		IdKabupaten:   int(f.IdKabupaten),
		KodeKabupaten: int(f.KodeKabupaten),
		NamaKabupaten: f.NamaKabupaten,
	}
}
