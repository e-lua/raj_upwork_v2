package available_traded

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

var FundamentalsRouter_pg *fundamentalsRouter_pg

type fundamentalsRouter_pg struct {
}

func (fr *fundamentalsRouter_pg) AddOneData(c echo.Context) error {

	//Instanciamos una variable del modelo Business Worker
	var incoming_newdata Incoming_NewData

	//Add variables from incomening request
	err := c.Bind(&incoming_newdata)
	if err != nil {
		results := Response_string{Error: true, DataError: "All requested data must be sent, review the structure or values", Data: ""}
		return c.JSON(400, results)
	}

	//Validating incoming values
	if len(incoming_newdata.Symbol) < 1 || len(incoming_newdata.Api_token) < 1 {
		results := Response_string{Error: true, DataError: "The values entered do not comply with the business rules", Data: ""}
		return c.JSON(400, results)
	}

	//Send data to the service
	status, boolerror, dataerror, data := AddOneData_Service(incoming_newdata)

	response := Response_string{
		Error:     boolerror,
		DataError: dataerror,
		Data:      data,
	}

	//Succesfull Response
	return c.JSON(status, response)
}

func (fr *fundamentalsRouter_pg) AddTradableSymbolList(c echo.Context) error {

	//Instanciamos una variable del modelo Business Worker
	var incoming_newdata Incoming_NewData

	//Add variables from incomening request
	err := c.Bind(&incoming_newdata)
	if err != nil {
		results := Response_string{Error: true, DataError: "All requested data must be sent, review the structure or values", Data: ""}
		return c.JSON(400, results)
	}

	//Validating incoming values
	if len(incoming_newdata.Api_token) < 1 {
		results := Response_string{Error: true, DataError: "The values entered do not comply with the business rules", Data: ""}
		return c.JSON(400, results)
	}

	//Send data to the service
	status, boolerror, dataerror, data := AddTradableSymbolList_Service(incoming_newdata)

	response := Response_string{
		Error:     boolerror,
		DataError: dataerror,
		Data:      data,
	}

	//Succesfull Response
	return c.JSON(status, response)
}

func (fr *fundamentalsRouter_pg) AddAllData(c echo.Context) error {

	//Instanciamos una variable del modelo Business Worker
	var incoming_newdata Incoming_NewData_ToUploadAllData

	//Add variables from incomening request
	err := c.Bind(&incoming_newdata)
	if err != nil {
		results := Response_string{Error: true, DataError: "All requested data must be sent, review the structure or values", Data: ""}
		return c.JSON(400, results)
	}

	//Validating incoming values
	if len(incoming_newdata.Api_token) < 1 {
		results := Response_string{Error: true, DataError: "The values entered do not comply with the business rules", Data: ""}
		return c.JSON(400, results)
	}

	//Send data to the service
	status, boolerror, dataerror, data := AddAllData_Service(incoming_newdata)

	response := Response_string{
		Error:     boolerror,
		DataError: dataerror,
		Data:      data,
	}

	//Succesfull Response
	return c.JSON(status, response)
}

func (fr *fundamentalsRouter_pg) GetIncomeStatement(c echo.Context) error {

	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_IncomenStatement_Annual{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetIncomeStatementAnnual_Service(symbol, limit_int)
		results := Response_IncomenStatement_Annual{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetIncomeStatementQuarter_Service(symbol, limit_int)
		results := Response_IncomenStatement_Quarter{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}

}

func (fr *fundamentalsRouter_pg) GetIncomeStatementGrowth(c echo.Context) error {

	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_IncomenStatement_AnnualGrowth{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetIncomeStatementAnnualGrowth_Service(symbol, limit_int)
		results := Response_IncomenStatement_AnnualGrowth{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetIncomeStatementQuarterGrowth_Service(symbol, limit_int)
		results := Response_IncomenStatement_QuarterGrowth{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}
}

func (fr *fundamentalsRouter_pg) GetBalanceSheetStatement(c echo.Context) error {

	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_BalanceSheet_Annual{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetBalanceSheetAnnual_Service(symbol, limit_int)
		results := Response_BalanceSheet_Annual{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetBalanceSheetQuarter_Service(symbol, limit_int)
		results := Response_BalanceSheet_Quarter{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}
}

func (fr *fundamentalsRouter_pg) GetBalanceSheetStatementGrowth(c echo.Context) error {

	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_BalanceSheet_AnnualGrowth{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetBalanceSheetAnnualGrowth_Service(symbol, limit_int)
		results := Response_BalanceSheet_AnnualGrowth{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetBalanceSheetQuarterGrowth_Service(symbol, limit_int)
		results := Response_BalanceSheet_QuarterGrowth{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}
}

func (fr *fundamentalsRouter_pg) GetCashFlow(c echo.Context) error {

	//Recibimos la fecha de la carta
	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_CashFlow_Annual{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetCashFlowAnnual_Service(symbol, limit_int)
		results := Response_CashFlow_Annual{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetCashFlowQuarter_Service(symbol, limit_int)
		results := Response_CashFlow_Quarter{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}
}

func (fr *fundamentalsRouter_pg) GetCashFlowGrowth(c echo.Context) error {

	//Recibimos la fecha de la carta
	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_CashFlow_AnnualGrowth{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetCashFlowAnnualGrowth_Service(symbol, limit_int)
		results := Response_CashFlow_AnnualGrowth{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetCashFlowQuarterGrowth_Service(symbol, limit_int)
		results := Response_CashFlow_QuarterGrowth{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}
}

func (fr *fundamentalsRouter_pg) GetRatios(c echo.Context) error {

	//Recibimos la fecha de la carta
	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_FinancialRatio_Annual{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetFinancialRatioAnnual_Service(symbol, limit_int)
		results := Response_FinancialRatio_Annual{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetFinancialRatioQuarter_Service(symbol, limit_int)
		results := Response_FinancialRatio_Quarter{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}
}

func (fr *fundamentalsRouter_pg) GetRatiosTTM(c echo.Context) error {

	//Recibimos la fecha de la carta
	symbol := c.Param("symbol")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_FinancialRatio_AnnualTTM{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	status, boolerror, dataerror, data := GetFinancialRatioAnnualTTM_Service(symbol, limit_int)
	results := Response_FinancialRatio_AnnualTTM{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (fr *fundamentalsRouter_pg) GetKeyMetrics(c echo.Context) error {

	//Recibimos la fecha de la carta
	symbol := c.Param("symbol")

	//Recibimos el limit
	period := c.Request().URL.Query().Get("period")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_KeyMetrics_Annual{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	if period != "quarter" {
		status, boolerror, dataerror, data := GetKeyMetricsAnnual_Service(symbol, limit_int)
		results := Response_KeyMetrics_Annual{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		status, boolerror, dataerror, data := GetKeyMetricsQuarter_Service(symbol, limit_int)
		results := Response_KeyMetrics_Quarter{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}
}

func (fr *fundamentalsRouter_pg) GetKeyMetricsTTM(c echo.Context) error {

	//Recibimos la fecha de la carta
	symbol := c.Param("symbol")

	//Recibimos el limit
	limit := c.Request().URL.Query().Get("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Validating incoming values
	if limit_int <= 0 {
		results := Response_KeyMetrics_CompanyTTM{Error: true, DataError: "The values entered do not comply with the business rules", Data: nil}
		return c.JSON(400, results)
	}

	status, boolerror, dataerror, data := GetKeyMetricsCompanyTTM_Service(symbol, limit_int)
	results := Response_KeyMetrics_CompanyTTM{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (fr *fundamentalsRouter_pg) GetAvailableTraded(c echo.Context) error {

	status, boolerror, dataerror, data := GetAvailableTraded_Service()
	results := Response_TradableSymbols{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (fr *fundamentalsRouter_pg) GetCompanyProfile(c echo.Context) error {

	//Recibimos la fecha de la carta
	symbol := c.Param("symbol")

	status, boolerror, dataerror, data := GetCompanyProfile_Service(symbol)
	results := Response_CompanyProfile{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (fr *fundamentalsRouter_pg) GetIndustries(c echo.Context) error {

	status, boolerror, dataerror, data := GetIndustries_Service()
	results := Response_ListString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (fr *fundamentalsRouter_pg) GetSectors(c echo.Context) error {

	status, boolerror, dataerror, data := GetSectors_Service()
	results := Response_ListString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
