package routes

import (
	"fhir-sirs/app/server"
	rlreport "fhir-sirs/pkg/api/v1/rlreport"
	rlreportUC "fhir-sirs/pkg/api/v1/rlreport/usecase"

	"github.com/labstack/echo/v4"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

type HTTP struct {
	rlSvc rlreport.Service
}

func NewHTTP(rlreportSvc *rlreportUC.RLReport, err *echo.Group) {
	h := HTTP{rlSvc: rlreportSvc}

	g := err.Group("", echotrace.Middleware())
	g.GET("/examples", h.getExamples)

	err.GET("/rl34", h.GetRL34)
	err.GET("/rl35", h.GetRL35)
	err.GET("/rl51", h.GetRL51)
}

// the function to be called in routing
func (h *HTTP) getExamples(c echo.Context) error {

	results := h.rlSvc.GetRL34()

	return server.ResponseStatusOK(c, "Success", results, nil, nil)
}

func (h *HTTP) GetRL34(c echo.Context) error {
	return server.ResponseStatusOK(c, "Success", h.rlSvc.GetRL34(), nil, nil)
}

func (h *HTTP) GetRL35(c echo.Context) error {
	return server.ResponseStatusOK(c, "Success", h.rlSvc.GetRL35(), nil, nil)
}

func (h *HTTP) GetRL51(c echo.Context) error {
	return server.ResponseStatusOK(c, "Success", h.rlSvc.GetRL51(), nil, nil)
}
