package routes

import (
	"fhir-sirs/app/config"
	"fhir-sirs/app/server"
	"fhir-sirs/pkg/api/v1/ckg_tb/models"
	"fhir-sirs/pkg/api/v1/ckg_tb/usecase"
	"fhir-sirs/pkg/api/v1/ckg_tb/utils"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	usecase *usecase.DataCKGTB
}

func NewHTTP(u *usecase.DataCKGTB, g *echo.Group) {
	h := HTTP{usecase: u}
	g.GET("/skrining", h.getData)
	g.POST("/status-pasien", h.postData)

	conf := config.GetConfig()
	fmt.Printf("env %+v\n", conf)
}

func (h *HTTP) getData(c echo.Context) error {
	tanggal := c.QueryParam("tgl")

	// if all query params are empty, return 400
	if tanggal == "" {
		return server.ResponseStatusBadRequest(
			c,
			"Query parameter 'tgl' is required",
			nil,
			nil,
			nil,
		)
	}

	halaman := 1
	hal := c.QueryParam("page")
	if !utils.IsNotEmptyString(&hal) {
		hal = "1"
	}
	num, err := strconv.Atoi(hal)
	if err == nil && num > 0 {
		halaman = num
	}

	data, err := h.usecase.GeTbtDataFiltered(tanggal, halaman)
	if err != nil {
		return server.ResponseStatusInternalServerError(c, err.Error(), nil, nil, err)
	}
	return server.ResponseStatusOK(c, "success", data, nil, nil)
}

func (h *HTTP) postData(c echo.Context) error {
	// Siapkan variabel untuk menampung data dari body request.
	var payload []models.StatusPasienTBInput

	// Bind body JSON dari request ke dalam variabel payload.
	if err := c.Bind(&payload); err != nil {
		return server.ResponseStatusBadRequest(
			c,
			"Invalid request body",
			nil,
			nil,
			err,
		)
	}

	result, err := h.usecase.PostTbPatientStatus(payload)
	if err != nil {
		return server.ResponseStatusBadRequest(
			c,
			err.Error(),
			nil,
			nil,
			err,
		)
	}
	return server.ResponseStatusOK(c, "Data received successfully", result, nil, nil)
}
