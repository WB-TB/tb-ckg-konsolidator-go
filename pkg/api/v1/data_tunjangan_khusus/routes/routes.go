package routes

import (
	"fhir-sirs/app/server"
	"fhir-sirs/pkg/api/v1/data_tunjangan_khusus/usecase"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	usecase *usecase.DataTunjanganKhusus
}

func NewHTTP(u *usecase.DataTunjanganKhusus, g *echo.Group) {
	h := HTTP{usecase: u}
	g.GET("/data_tunjangan_khusus", h.getData)
}

func (h *HTTP) getData(c echo.Context) error {
	orgID := c.QueryParam("organization_id")
	nikDrSp := c.QueryParam("nik_drSp")
	tanggal := c.QueryParam("tanggal")

	//data := h.usecase.GetDummyDataFiltered(orgID, nikDrSp, tanggal)
	//data := h.usecase.GetDummyData()

	// if all query params are empty, return 400
	if orgID == "" && nikDrSp == "" && tanggal == "" {
		return server.ResponseStatusBadRequest(
			c,
			"at least one query param (organization_id, nik_drSp, tanggal) is required",
			nil,
			nil,
			nil,
		)
	}

	data, err := h.usecase.GetRealDataFiltered(orgID, nikDrSp, tanggal)
	if err != nil {
		return server.ResponseStatusInternalServerError(c, "DB error", nil, nil, err)
	}
	return server.ResponseStatusOK(c, "success", data, nil, nil)
}
