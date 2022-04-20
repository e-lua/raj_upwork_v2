package available_traded

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	models "github.com/Aphofisis/raj_upwork_v2/models"
	all "github.com/Aphofisis/raj_upwork_v2/repositories/all"
	balanceSheet "github.com/Aphofisis/raj_upwork_v2/repositories/balanceSheet"
	cashFlow "github.com/Aphofisis/raj_upwork_v2/repositories/cashFlow"
	financialRatios "github.com/Aphofisis/raj_upwork_v2/repositories/financialRatios"
	incomeStatement "github.com/Aphofisis/raj_upwork_v2/repositories/incomeStatement"
	keyMetrics "github.com/Aphofisis/raj_upwork_v2/repositories/keyMetrics"
	profile "github.com/Aphofisis/raj_upwork_v2/repositories/profile"
	tradableSymbols "github.com/Aphofisis/raj_upwork_v2/repositories/tradableSymbols"
)

//AddAllData_Service add the data for each symbol
func AddAllData_Service(input_data Incoming_NewData_ToUploadAllData) (int, bool, string, string) {

	output_index := input_data.Index

	for {

		log.Println("===================================== " + strconv.Itoa(output_index) + " =====================================")
		get_respuesta_trad, _ := tradableSymbols.Si_Find_WithLimit(output_index)
		counter := 0
		for _, val := range get_respuesta_trad {

			var inco_newdata Incoming_NewData
			inco_newdata.Symbol = val.Symbol
			inco_newdata.Api_token = input_data.Api_token

			log.Println("--->>>>>>>>SYMBOL TO CHARGE: " + val.Symbol + " " + strconv.Itoa(output_index) + "." + strconv.Itoa(counter))

			_, boolerror, dataerror, _ := AddOneData_Service(inco_newdata)
			if boolerror {
				log.Println(val.Symbol, "result ? ", dataerror)
			}

			time.Sleep(15 * time.Second)

			counter = counter + 1
		}

		if output_index == 1048 {
			log.Println("****===============================END ALL===============================****")
			break
		}

		output_index = output_index + 1
		time.Sleep(1 * time.Minute)
	}

	return 200, false, "", "OK"
}

//AddTradableSymbolList_Service add the tradable symbols
func AddTradableSymbolList_Service(input_data Incoming_NewData) (int, bool, string, string) {

	var get_respuesta_trad []models.TradableSymbols_Income

	source_data, error_get := http.Get("https://fmpcloud.io/api/v3/available-traded/list?apikey=" + input_data.Api_token)
	if error_get != nil {
		return 403, true, "Internal error at the moment to get the data from TradableSymbols, details: " + error_get.Error(), ""
	}
	error_decode_respuesta := json.NewDecoder(source_data.Body).Decode(&get_respuesta_trad)
	if error_decode_respuesta != nil {
		return 403, true, "Internal error at the moment to get the data from TradableSymbols, details: " + error_get.Error(), ""
	}
	log.Print("-------->Traded list-> extracted")

	error_add_tr := tradableSymbols.Si_Add(get_respuesta_trad)
	if error_add_tr != nil {
		return 403, true, "Internal error when Tradable List data load started: " + error_add_tr.Error(), ""
	}

	return 200, false, "", "OK"
}

//AddOneData_Service add the data just for one symbol
func AddOneData_Service(input_data Incoming_NewData) (int, bool, string, string) {

	log.Print("-------->VALIDATING IF THE DATA ALREADY EXISTS")
	symbol, _ := all.Si_Find(input_data.Symbol)
	if symbol != "" {

		log.Print("-------->THIS DATA ALREADY EXISTS -STARTING EXTRACTION DATA PROCESS")
		error_delete := all.Si_Delete(input_data.Symbol)
		if error_delete != nil {
			return 500, true, "Error trying to delete duplicate data", ""
		}
	}

	log.Print("-------->STARTING EXTRACTION DATA PROCESS")

	var get_respuesta_companyprofile []models.CompanyProfile

	source_data, error_get := http.Get("https://fmpcloud.io/api/v3/profile/" + input_data.Symbol + "?apikey=" + input_data.Api_token)
	if error_get != nil {
		return 403, true, "Internal error at the moment to get the data from TradableSymbols, details: " + error_get.Error(), ""
	}
	json.NewDecoder(source_data.Body).Decode(&get_respuesta_companyprofile)

	log.Print("-------->Company profile-> extracted")

	time.Sleep(1 * time.Second)

	//Income Statement Annual
	source_data_isa, error_get_isa := http.Get("https://fmpcloud.io/api/v3/income-statement/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isa []models.IncomeStatement_Annual
	if error_get_isa != nil {
		log.Println(source_data_isa)
		return 403, true, "Internal error at the moment to get the data from Income Statement Annual, details: " + error_get_isa.Error(), ""
	}
	json.NewDecoder(source_data_isa.Body).Decode(&get_respuesta_isa)

	log.Print("-------->Income Statement Annual-> extracted")

	time.Sleep(1 * time.Second)

	//Income Statement Quarter
	source_data_isq, error_get_isq := http.Get("https://fmpcloud.io/api/v3/income-statement/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isq []models.IncomeStatement_Quarter
	if error_get_isq != nil {
		log.Println(source_data_isq)
		return 403, true, "Internal error at the moment to get the data from Income Statement Quarter, details: " + error_get_isq.Error(), ""
	}
	json.NewDecoder(source_data_isq.Body).Decode(&get_respuesta_isq)

	log.Print("-------->Income Statement Quarter-> extracted")

	time.Sleep(1 * time.Second)

	//Income Statement Annual Growth
	source_data_isag, error_get_isag := http.Get("https://fmpcloud.io/api/v3/income-statement-growth/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isa_isag []models.IncomeStatement_AnnualGrowth
	if error_get_isa != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Annual Growth, details: " + error_get_isag.Error(), ""
	}
	json.NewDecoder(source_data_isag.Body).Decode(&get_respuesta_isa_isag)

	log.Print("-------->Income Statement Annual Growth-> extracted")

	time.Sleep(1 * time.Second)

	//Income Statement Quarter Growth
	source_data_isqg, error_get_isqg := http.Get("https://fmpcloud.io/api/v3/income-statement-growth/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isqg []models.IncomeStatement_QuarterGrowth
	if error_get_isqg != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Quarter Growth, details: " + error_get_isqg.Error(), ""
	}
	json.NewDecoder(source_data_isqg.Body).Decode(&get_respuesta_isqg)

	log.Print("-------->Income Statement Quarter Growth-> extracted")

	time.Sleep(1 * time.Second)

	//Balance Sheet Annual
	source_data_bsa, error_get_bsa := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsa []models.BalanceSheet_Annual
	if error_get_isa != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Annual, details: " + error_get_bsa.Error(), ""
	}
	json.NewDecoder(source_data_bsa.Body).Decode(&get_respuesta_bsa)

	log.Print("-------->Balance Sheet Annual-> extracted")

	time.Sleep(1 * time.Second)

	//Balance Sheet Quarter
	source_data_bsq, error_get_bsq := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsq []models.BalanceSheet_Quarter
	if error_get_bsq != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Quarter, details: " + error_get_bsq.Error(), ""
	}
	json.NewDecoder(source_data_bsq.Body).Decode(&get_respuesta_bsq)

	log.Print("-------->Balance Sheet Quarter-> extracted")

	time.Sleep(1 * time.Second)

	//Balance Sheet Annual Growth
	source_data_bsag, error_get_bsag := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement-growth/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsag []models.BalanceSheet_AnnualGrowth
	if error_get_bsag != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Annual Growth, details: " + error_get_bsag.Error(), ""
	}
	json.NewDecoder(source_data_bsag.Body).Decode(&get_respuesta_bsag)

	log.Print("-------->Balance Sheet Annual Growth-> extracted")

	time.Sleep(1 * time.Second)

	//Balance Sheet Quarter Growth
	source_data_bsqg, error_get_bsqg := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement-growth/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsqg []models.BalanceSheet_QuarterGrowth
	if error_get_bsag != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Quarter Growth, details: " + error_get_bsqg.Error(), ""
	}
	json.NewDecoder(source_data_bsqg.Body).Decode(&get_respuesta_bsqg)

	log.Print("-------->Balance Sheet Quarter Growth-> extracted")

	time.Sleep(1 * time.Second)

	//Cash Flow Annual
	source_data_cfa, error_get_cfa := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfa []models.CashFlow_Annual
	if error_get_cfa != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Annual, details: " + error_get_isa.Error(), ""
	}
	json.NewDecoder(source_data_cfa.Body).Decode(&get_respuesta_cfa)

	log.Print("-------->Cash Flow Annual-> extracted")

	time.Sleep(1 * time.Second)

	//Cash Flow Quarter
	source_data_cfq, error_get_cfq := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfq []models.CashFlow_Quarter
	if error_get_cfq != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Quarter, details: " + error_get_cfq.Error(), ""
	}
	json.NewDecoder(source_data_cfq.Body).Decode(&get_respuesta_cfq)

	log.Print("-------->Cash Flow Quarter-> extracted")

	time.Sleep(1 * time.Second)

	//Cash Flow Annual Growth
	source_data_cfag, error_get_cfag := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement-growth/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfag []models.CashFlow_AnnualGrowth
	if error_get_cfag != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Annual Growth, details: " + error_get_isa.Error(), ""
	}
	json.NewDecoder(source_data_cfag.Body).Decode(&get_respuesta_cfag)

	log.Print("-------->Cash Flow Annual Growth-> extracted")

	time.Sleep(1 * time.Second)

	//Cash Flow Quarter Growth
	source_data_cfqg, error_get_cfqg := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement-growth/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfqg []models.CashFlow_QuarterGrowth
	if error_get_cfqg != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Quarter Growth, details: " + error_get_cfqg.Error(), ""
	}
	json.NewDecoder(source_data_cfqg.Body).Decode(&get_respuesta_cfqg)

	log.Print("-------->Cash Flow Quarter Growth-> extracted")

	time.Sleep(1 * time.Second)

	//Financial Ratio Annual
	source_data_fra, error_get_fra := http.Get("https://fmpcloud.io/api/v3/ratios/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_fra []models.FinancialRatio_Annual
	if error_get_fra != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Annual, details: " + error_get_fra.Error(), ""
	}
	json.NewDecoder(source_data_fra.Body).Decode(&get_respuesta_fra)

	log.Print("-------->Financial Ratio Annual-> extracted")

	time.Sleep(1 * time.Second)

	//Financial Ratio Quarter
	source_data_frq, error_get_frq := http.Get("https://fmpcloud.io/api/v3/ratios/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_frq []models.FinancialRatio_Quarter
	if error_get_frq != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Quarter, details: " + error_get_frq.Error(), ""
	}
	json.NewDecoder(source_data_frq.Body).Decode(&get_respuesta_frq)

	log.Print("-------->Financial Ratio Quarter-> extracted")

	time.Sleep(1 * time.Second)

	//Financial Ratio AnnualTTM
	source_data_frattm, error_get_frattm := http.Get("https://fmpcloud.io/api/v3/ratios-ttm/" + input_data.Symbol + "?apikey=" + input_data.Api_token)
	var get_respuesta_frattm []models.FinancialRatio_AnnualTTM
	if error_get_frattm != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Annual, details: " + error_get_frattm.Error(), ""
	}
	json.NewDecoder(source_data_frattm.Body).Decode(&get_respuesta_frattm)

	log.Print("-------->Financial Ratio Annual-> extracted")

	time.Sleep(1 * time.Second)

	//Key Metrics Company TTM
	source_data_kmcttm, error_get_kmcttm := http.Get("https://fmpcloud.io/api/v3/key-metrics-ttm/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_kmcttm []models.KeyMetrics_CompanyTTM
	if error_get_kmcttm != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Company TTM, details: " + error_get_kmcttm.Error(), ""
	}
	json.NewDecoder(source_data_kmcttm.Body).Decode(&get_respuesta_kmcttm)

	log.Print("-------->Key Metrics Company TTM-> extracted")

	time.Sleep(1 * time.Second)

	//Key Metrics Annual
	source_data_kma, error_get_kma := http.Get("https://fmpcloud.io/api/v3/key-metrics/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_kma []models.KeyMetrics_Annual
	if error_get_kma != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Annual, details: " + error_get_kma.Error(), ""
	}
	json.NewDecoder(source_data_kma.Body).Decode(&get_respuesta_kma)

	log.Print("-------->Key Metrics Annual-> extracted")

	time.Sleep(1 * time.Second)

	//Key Metrics Quarter
	source_data_kmq, error_get_kmq := http.Get("https://fmpcloud.io/api/v3/key-metrics/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_kmq []models.KeyMetrics_Quarter
	if error_get_kmq != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Quarter, details: " + error_get_kmq.Error(), ""
	}
	json.NewDecoder(source_data_kmq.Body).Decode(&get_respuesta_kmq)

	log.Print("-------->Key Metrics Quarter-> extracted")
	time.Sleep(1 * time.Second)

	log.Print("-------->SUCCESSFUL EXTRACTION DATA PROCESS")

	error_add := all.Si_Add(input_data.Symbol, get_respuesta_companyprofile, get_respuesta_isa, get_respuesta_isq, get_respuesta_isa_isag, get_respuesta_isqg, get_respuesta_bsa, get_respuesta_bsq, get_respuesta_bsag, get_respuesta_bsqg, get_respuesta_cfa, get_respuesta_cfq, get_respuesta_cfag, get_respuesta_cfqg, get_respuesta_fra, get_respuesta_frq, get_respuesta_frattm, get_respuesta_kmcttm, get_respuesta_kma, get_respuesta_kmq)
	if error_add != nil {
		log.Println(403, true, "Internal error when data load started: "+error_add.Error(), " ")
	}

	return 201, false, "", "OK"
}

func GetIncomeStatementAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_Annual_Response) {

	incomeAnnual, _ := incomeStatement.Si_Find_Annual(symbol, limit_int)

	return 201, false, "", incomeAnnual
}

func GetIncomeStatementQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_Quarter_Response) {

	incomeQuarter, _ := incomeStatement.Si_Find_Quarter(symbol, limit_int)

	return 201, false, "", incomeQuarter
}

func GetIncomeStatementAnnualGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_AnnualGrowth_Response) {

	incomeAnnualGrwoth, _ := incomeStatement.Si_Find_AnnualGrowth(symbol, limit_int)

	return 201, false, "", incomeAnnualGrwoth
}

func GetIncomeStatementQuarterGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_QuarterGrowth_Response) {

	incomeQuarterGrwoth, _ := incomeStatement.Si_Find_QuarterGrowth(symbol, limit_int)

	return 201, false, "", incomeQuarterGrwoth
}

func GetBalanceSheetAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_Annual_Response) {

	balanceSheetAnnual, _ := balanceSheet.Si_Find_Annual(symbol, limit_int)

	return 201, false, "", balanceSheetAnnual
}

func GetBalanceSheetQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_Quarter_Response) {

	balanceSheetQuarter, _ := balanceSheet.Si_Find_Quarter(symbol, limit_int)

	return 201, false, "", balanceSheetQuarter
}

func GetBalanceSheetAnnualGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_AnnualGrowth_Response) {

	balanceSheetAnnualGrowth, _ := balanceSheet.Si_Find_AnnualGrowth(symbol, limit_int)

	return 201, false, "", balanceSheetAnnualGrowth
}

func GetBalanceSheetQuarterGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_QuarterGrowth_Response) {

	balanceSheetQuarterGrowth, _ := balanceSheet.Si_Find_QuarterGrowth(symbol, limit_int)

	return 201, false, "", balanceSheetQuarterGrowth
}

func GetCashFlowAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_Annual_Response) {

	cashFlowAnnual, _ := cashFlow.Si_Find_Annual(symbol, limit_int)

	return 201, false, "", cashFlowAnnual
}

func GetCashFlowQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_Quarter_Response) {

	cashFlowQuarter, _ := cashFlow.Si_Find_Quarter(symbol, limit_int)

	return 201, false, "", cashFlowQuarter
}

func GetCashFlowAnnualGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_AnnualGrowth_Response) {

	cashFlowAnnualGrowth, _ := cashFlow.Si_Find_AnnualGrowth(symbol, limit_int)

	return 201, false, "", cashFlowAnnualGrowth
}

func GetCashFlowQuarterGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_QuarterGrowth_Response) {

	cashFlowQuarterGrowth, _ := cashFlow.Si_Find_QuarterGrowth(symbol, limit_int)

	return 201, false, "", cashFlowQuarterGrowth
}

func GetFinancialRatioAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.FinancialRatio_Annual_Response) {

	financialRatioAnnual, _ := financialRatios.Si_Find_Annual(symbol, limit_int)

	return 201, false, "", financialRatioAnnual
}

func GetFinancialRatioQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.FinancialRatio_Quarter_Response) {

	financialRatioQuarter, _ := financialRatios.Si_Find_Quarter(symbol, limit_int)

	return 201, false, "", financialRatioQuarter
}

func GetFinancialRatioAnnualTTM_Service(symbol string, limit_int int) (int, bool, string, []models.FinancialRatio_AnnualTTM_Response) {

	financialRatioAnnualTTM, _ := financialRatios.Si_Find_AnnualTTM(symbol, limit_int)

	return 201, false, "", financialRatioAnnualTTM
}

func GetKeyMetricsAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.KeyMetrics_Annual_Response) {

	keymetricsAnnual, _ := keyMetrics.Si_Find_Annual(symbol, limit_int)

	return 201, false, "", keymetricsAnnual
}

func GetKeyMetricsQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.KeyMetrics_Quarter_Response) {

	keymetricsQuarter, _ := keyMetrics.Si_Find_Quarter(symbol, limit_int)

	return 201, false, "", keymetricsQuarter
}

func GetKeyMetricsCompanyTTM_Service(symbol string, limit_int int) (int, bool, string, []models.KeyMetrics_CompanyTTM_Response) {

	keymetricsCompanyTTM, _ := keyMetrics.Si_Find_companyTTM(symbol, limit_int)

	return 201, false, "", keymetricsCompanyTTM
}

func GetAvailableTraded_Service() (int, bool, string, []models.TradableSymbols) {

	tradableSymbols, _ := tradableSymbols.Si_Find()
	return 201, false, "", tradableSymbols
}

func GetCompanyProfile_Service(symbols string) (int, bool, string, models.CompanyProfile) {

	profile, _ := profile.Si_Find_Profile(symbols)

	return 201, false, "", profile
}

func GetIndustries_Service() (int, bool, string, []string) {

	industries, _ := all.Si_Find_Industries()

	return 201, false, "", industries
}

func GetSectors_Service() (int, bool, string, []string) {

	sectors, _ := all.Si_Find_Sectors()

	return 201, false, "", sectors
}
