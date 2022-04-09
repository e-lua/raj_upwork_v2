package api

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	service "github.com/Aphofisis/raj_upwork_v2/services"
)

func Handlers() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/api/v1")

	//V1->Fundamentals
	router_fundamentals := version_1.Group("/fundamentals")
	router_fundamentals.PUT("", service.FundamentalsRouter_pg.AddAllData)
	router_fundamentals.GET("/income-statement/:symbol", service.FundamentalsRouter_pg.GetIncomeStatement)
	router_fundamentals.GET("/income-statement-growth/:symbol", service.FundamentalsRouter_pg.GetIncomeStatementGrowth)
	router_fundamentals.GET("/balance-sheet-statement/:symbol", service.FundamentalsRouter_pg.GetBalanceSheetStatement)
	router_fundamentals.GET("/balance-sheet-statement-growth/:symbol", service.FundamentalsRouter_pg.GetBalanceSheetStatementGrowth)
	router_fundamentals.GET("/cash-flow-statement/:symbol", service.FundamentalsRouter_pg.GetCashFlow)
	router_fundamentals.GET("/cash-flow-statement-growth/:symbol", service.FundamentalsRouter_pg.GetCashFlowGrowth)
	router_fundamentals.GET("/ratios/:symbol", service.FundamentalsRouter_pg.GetRatios)
	router_fundamentals.GET("/ratios-ttm/:symbol", service.FundamentalsRouter_pg.GetRatiosTTM)
	router_fundamentals.GET("/key-metrics/:symbol", service.FundamentalsRouter_pg.GetKeyMetrics)
	router_fundamentals.GET("/key-metrics-ttm/:symbol", service.FundamentalsRouter_pg.GetKeyMetricsTTM)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "6500"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Restricted Access")
}
