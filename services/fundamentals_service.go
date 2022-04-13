package available_traded

import (
	"encoding/json"
	"log"
	"net/http"

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

func AddAllData_Service(input_data Incoming_NewData) (int, bool, string, string) {

	var get_respuesta_trad []models.TradableSymbols
	//Trader list

	source_data, error_get := http.Get("https://fmpcloud.io/api/v3/available-traded/list?apikey=" + input_data.Api_token)
	if error_get != nil {
		return 403, true, "Internal error at the moment to get the data from TradableSymbols, details: " + error_get.Error(), ""
	}
	error_decode_respuesta := json.NewDecoder(source_data.Body).Decode(&get_respuesta_trad)
	if error_decode_respuesta != nil {
		return 403, true, "Internal error at the moment to get the data from TradableSymbols, details: " + error_get.Error(), ""
	}
	log.Print("-------->Traded list-> extracted")

	counter := 0
	for _, val := range get_respuesta_trad {

		var inco_newdata Incoming_NewData
		inco_newdata.Symbol = val.Symbol
		inco_newdata.Api_token = input_data.Api_token

		_, boolerror, dataerror, _ := AddOneData_Service(inco_newdata)
		if boolerror {
			log.Println(val.Symbol, input_data.Api_token)
			log.Println(val.Symbol, "result ? ", dataerror)
		}

		counter = counter + 1
	}

	return 200, false, "", "OK"
}

func AddTradableSymbolList_Service(input_data Incoming_NewData) (int, bool, string, string) {

	var get_respuesta_trad []models.TradableSymbols

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
	error_decode_respuesta_cp := json.NewDecoder(source_data.Body).Decode(&get_respuesta_companyprofile)
	if error_decode_respuesta_cp != nil {
		return 403, true, "Internal error at the moment to get the data from TradableSymbols, details: " + error_decode_respuesta_cp.Error(), ""
	}
	log.Print("-------->Company profile-> extracted")

	//Income Statement Annual
	source_data_isa, error_get_isa := http.Get("https://fmpcloud.io/api/v3/income-statement/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isa []models.IncomeStatement_Annual
	if error_get_isa != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Annual, details: " + error_get_isa.Error(), ""
	}
	error_decode_respuesta_isa := json.NewDecoder(source_data_isa.Body).Decode(&get_respuesta_isa)
	if error_decode_respuesta_isa != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Annual, details: " + error_decode_respuesta_isa.Error(), ""
	}
	log.Print("-------->Income Statement Annual-> extracted")

	//Income Statement Quarter
	source_data_isq, error_get_isq := http.Get("https://fmpcloud.io/api/v3/income-statement/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isq []models.IncomeStatement_Quarter
	if error_get_isq != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Quarter, details: " + error_get_isq.Error(), ""
	}
	error_decode_respuesta := json.NewDecoder(source_data_isq.Body).Decode(&get_respuesta_isq)
	if error_decode_respuesta != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Quarter, details: " + error_get_isq.Error(), ""
	}
	log.Print("-------->Income Statement Quarter-> extracted")

	//Income Statement Annual Growth
	source_data_isag, error_get_isag := http.Get("https://fmpcloud.io/api/v3/income-statement-growth/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isa_isag []models.IncomeStatement_AnnualGrowth
	if error_get_isa != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Annual Growth, details: " + error_get_isag.Error(), ""
	}
	error_decode_respuesta_isag := json.NewDecoder(source_data_isag.Body).Decode(&get_respuesta_isa_isag)
	if error_decode_respuesta_isag != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Annual Growth, details: " + error_decode_respuesta_isag.Error(), ""
	}
	log.Print("-------->Income Statement Annual Growth-> extracted")

	//Income Statement Quarter Growth
	source_data_isqg, error_get_isqg := http.Get("https://fmpcloud.io/api/v3/income-statement-growth/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_isqg []models.IncomeStatement_QuarterGrowth
	if error_get_isqg != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Quarter Growth, details: " + error_get_isqg.Error(), ""
	}
	error_decode_respuesta_isqg := json.NewDecoder(source_data_isqg.Body).Decode(&get_respuesta_isqg)
	if error_decode_respuesta_isqg != nil {
		return 403, true, "Internal error at the moment to get the data from Income Statement Quarter Growth, details: " + error_decode_respuesta_isqg.Error(), ""
	}
	log.Print("-------->Income Statement Quarter Growth-> extracted")

	//Balance Sheet Annual
	source_data_bsa, error_get_bsa := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsa []models.BalanceSheet_Annual
	if error_get_isa != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Annual, details: " + error_get_bsa.Error(), ""
	}
	error_decode_respuesta_bsa := json.NewDecoder(source_data_bsa.Body).Decode(&get_respuesta_bsa)
	if error_decode_respuesta_bsa != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Annual, details: " + error_decode_respuesta_bsa.Error(), ""
	}
	log.Print("-------->Balance Sheet Annual-> extracted")

	//Balance Sheet Quarter
	source_data_bsq, error_get_bsq := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsq []models.BalanceSheet_Quarter
	if error_get_bsq != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Quarter, details: " + error_get_bsq.Error(), ""
	}
	error_decode_respuesta_bsq := json.NewDecoder(source_data_bsq.Body).Decode(&get_respuesta_bsq)
	if error_decode_respuesta_bsq != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Quarter, details: " + error_decode_respuesta_bsq.Error(), ""
	}
	log.Print("-------->Balance Sheet Quarter-> extracted")

	//Balance Sheet Annual Growth
	source_data_bsag, error_get_bsag := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement-growth/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsag []models.BalanceSheet_AnnualGrowth
	if error_get_bsag != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Annual Growth, details: " + error_get_bsag.Error(), ""
	}
	error_decode_respuesta_bsag := json.NewDecoder(source_data_bsag.Body).Decode(&get_respuesta_bsag)
	if error_decode_respuesta_bsag != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Annual Growth, details: " + error_decode_respuesta_bsag.Error(), ""
	}
	log.Print("-------->Balance Sheet Annual Growth-> extracted")

	//Balance Sheet Quarter Growth
	source_data_bsqg, error_get_bsqg := http.Get("https://fmpcloud.io/api/v3/balance-sheet-statement-growth/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_bsqg []models.BalanceSheet_QuarterGrowth
	if error_get_bsag != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Quarter Growth, details: " + error_get_bsqg.Error(), ""
	}
	error_decode_respuesta_bsqg := json.NewDecoder(source_data_bsqg.Body).Decode(&get_respuesta_bsqg)
	if error_decode_respuesta_bsqg != nil {
		return 403, true, "Internal error at the moment to get the data from Balance Sheet Quarter Growth, details: " + error_decode_respuesta_bsqg.Error(), ""
	}
	log.Print("-------->Balance Sheet Quarter Growth-> extracted")

	//Cash Flow Annual
	source_data_cfa, error_get_cfa := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfa []models.CashFlow_Annual
	if error_get_cfa != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Annual, details: " + error_get_isa.Error(), ""
	}
	error_decode_respuesta_cfa := json.NewDecoder(source_data_cfa.Body).Decode(&get_respuesta_cfa)
	if error_decode_respuesta_cfa != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Annual, details: " + error_decode_respuesta_cfa.Error(), ""
	}
	log.Print("-------->Cash Flow Annual-> extracted")

	//Cash Flow Quarter
	source_data_cfq, error_get_cfq := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfq []models.CashFlow_Quarter
	if error_get_cfq != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Quarter, details: " + error_get_cfq.Error(), ""
	}
	error_decode_respuesta_cfq := json.NewDecoder(source_data_cfq.Body).Decode(&get_respuesta_cfq)
	if error_decode_respuesta_cfq != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Quarter, details: " + error_decode_respuesta_cfq.Error(), ""
	}
	log.Print("-------->Cash Flow Quarter-> extracted")

	//Cash Flow Annual Growth
	source_data_cfag, error_get_cfag := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement-growth/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfag []models.CashFlow_AnnualGrowth
	if error_get_cfag != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Annual Growth, details: " + error_get_isa.Error(), ""
	}
	error_decode_respuesta_cfag := json.NewDecoder(source_data_cfag.Body).Decode(&get_respuesta_cfag)
	if error_decode_respuesta_cfag != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Annual Growth, details: " + error_decode_respuesta_cfag.Error(), ""
	}
	log.Print("-------->Cash Flow Annual Growth-> extracted")

	//Cash Flow Quarter Growth
	source_data_cfqg, error_get_cfqg := http.Get("https://fmpcloud.io/api/v3/cash-flow-statement-growth/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_cfqg []models.CashFlow_QuarterGrowth
	if error_get_cfqg != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Quarter Growth, details: " + error_get_cfqg.Error(), ""
	}
	error_decode_respuesta_cfqg := json.NewDecoder(source_data_cfqg.Body).Decode(&get_respuesta_cfqg)
	if error_decode_respuesta_cfqg != nil {
		return 403, true, "Internal error at the moment to get the data from Cash Flow Quarter Growth, details: " + error_decode_respuesta_cfqg.Error(), ""
	}
	log.Print("-------->Cash Flow Quarter Growth-> extracted")

	//Financial Ratio Annual
	source_data_fra, error_get_fra := http.Get("https://fmpcloud.io/api/v3/ratios/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_fra []models.FinancialRatio_Annual
	if error_get_fra != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Annual, details: " + error_get_fra.Error(), ""
	}
	error_decode_respuesta_fra := json.NewDecoder(source_data_fra.Body).Decode(&get_respuesta_fra)
	if error_decode_respuesta_fra != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Annual, details: " + error_decode_respuesta_fra.Error(), ""
	}
	log.Print("-------->Financial Ratio Annual-> extracted")

	//Financial Ratio Quarter
	source_data_frq, error_get_frq := http.Get("https://fmpcloud.io/api/v3/ratios/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_frq []models.FinancialRatio_Quarter
	if error_get_frq != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Quarter, details: " + error_get_frq.Error(), ""
	}
	error_decode_respuesta_frq := json.NewDecoder(source_data_frq.Body).Decode(&get_respuesta_frq)
	if error_decode_respuesta_frq != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Quarter, details: " + error_decode_respuesta_frq.Error(), ""
	}
	log.Print("-------->Financial Ratio Quarter-> extracted")

	//Financial Ratio AnnualTTM
	source_data_frattm, error_get_frattm := http.Get("https://fmpcloud.io/api/v3/ratios-ttm/" + input_data.Symbol + "?apikey=" + input_data.Api_token)
	var get_respuesta_frattm []models.FinancialRatio_AnnualTTM
	if error_get_frattm != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Annual, details: " + error_get_frattm.Error(), ""
	}
	error_decode_respuesta_frattm := json.NewDecoder(source_data_frattm.Body).Decode(&get_respuesta_frattm)
	if error_decode_respuesta_frattm != nil {
		return 403, true, "Internal error at the moment to get the data from Financial Ratio Annual, details: " + error_decode_respuesta_frattm.Error(), ""
	}
	log.Print("-------->Financial Ratio Annual-> extracted")

	//Key Metrics Company TTM
	source_data_kmcttm, error_get_kmcttm := http.Get("https://fmpcloud.io/api/v3/key-metrics-ttm/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_kmcttm []models.KeyMetrics_CompanyTTM
	if error_get_kmcttm != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Company TTM, details: " + error_get_kmcttm.Error(), ""
	}
	error_decode_respuesta_kmcttm := json.NewDecoder(source_data_kmcttm.Body).Decode(&get_respuesta_kmcttm)
	if error_decode_respuesta_kmcttm != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Company TTM, details: " + error_decode_respuesta_kmcttm.Error(), ""
	}
	log.Print("-------->Key Metrics Company TTM-> extracted")

	//Key Metrics Annual
	source_data_kma, error_get_kma := http.Get("https://fmpcloud.io/api/v3/key-metrics/" + input_data.Symbol + "?limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_kma []models.KeyMetrics_Annual
	if error_get_kma != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Annual, details: " + error_get_kma.Error(), ""
	}
	error_decode_respuesta_kma := json.NewDecoder(source_data_kma.Body).Decode(&get_respuesta_kma)
	if error_decode_respuesta_kma != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Annual, details: " + error_decode_respuesta_kma.Error(), ""
	}
	log.Print("-------->Key Metrics Annual-> extracted")

	//Key Metrics Quarter
	source_data_kmq, error_get_kmq := http.Get("https://fmpcloud.io/api/v3/key-metrics/" + input_data.Symbol + "?period=quarter&limit=800000&apikey=" + input_data.Api_token)
	var get_respuesta_kmq []models.KeyMetrics_Quarter
	if error_get_kmq != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Quarter, details: " + error_get_kmq.Error(), ""
	}
	error_decode_respuesta_kmq := json.NewDecoder(source_data_kmq.Body).Decode(&get_respuesta_kmq)
	if error_decode_respuesta_kmq != nil {
		return 403, true, "Internal error at the moment to get the data from Key Metrics Quarter, details: " + error_decode_respuesta_kmq.Error(), ""
	}
	log.Print("-------->Key Metrics Quarter-> extracted")
	log.Print("-------->SUCCESSFUL EXTRACTION DATA PROCESS")

	error_add := all.Si_Add(input_data.Symbol, get_respuesta_companyprofile, get_respuesta_isa, get_respuesta_isq, get_respuesta_isa_isag, get_respuesta_isqg, get_respuesta_bsa, get_respuesta_bsq, get_respuesta_bsag, get_respuesta_bsqg, get_respuesta_cfa, get_respuesta_cfq, get_respuesta_cfag, get_respuesta_cfqg, get_respuesta_fra, get_respuesta_frq, get_respuesta_frattm, get_respuesta_kmcttm, get_respuesta_kma, get_respuesta_kmq)
	if error_add != nil {
		log.Println(403, true, "Internal error when data load started: "+error_add.Error(), " ")
	}

	return 201, false, "", "OK"
}

func GetIncomeStatementAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_Annual_Response) {

	incomeAnnual, error_find_all := incomeStatement.Si_Find_Annual(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching IncomeStatement Annual data, details: " + error_find_all.Error(), incomeAnnual
	}

	return 201, false, "", incomeAnnual
}

func GetIncomeStatementQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_Quarter_Response) {

	incomeQuarter, error_find_all := incomeStatement.Si_Find_Quarter(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching IncomeStatement Quarter data, details: " + error_find_all.Error(), incomeQuarter
	}

	return 201, false, "", incomeQuarter
}

func GetIncomeStatementAnnualGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_AnnualGrowth_Response) {

	incomeAnnualGrwoth, error_find_all := incomeStatement.Si_Find_AnnualGrowth(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching IncomeStatement Annual Growth data, details: " + error_find_all.Error(), incomeAnnualGrwoth
	}

	return 201, false, "", incomeAnnualGrwoth
}

func GetIncomeStatementQuarterGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.IncomeStatement_QuarterGrowth_Response) {

	incomeQuarterGrwoth, error_find_all := incomeStatement.Si_Find_QuarterGrowth(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching IncomeStatement Quarter Growth data, details: " + error_find_all.Error(), incomeQuarterGrwoth
	}

	return 201, false, "", incomeQuarterGrwoth
}

func GetBalanceSheetAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_Annual_Response) {

	balanceSheetAnnual, error_find_all := balanceSheet.Si_Find_Annual(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching BalanceSheet Annual data, details: " + error_find_all.Error(), balanceSheetAnnual
	}

	return 201, false, "", balanceSheetAnnual
}

func GetBalanceSheetQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_Quarter_Response) {

	balanceSheetQuarter, error_find_all := balanceSheet.Si_Find_Quarter(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching BalanceSheet Quarter data, details: " + error_find_all.Error(), balanceSheetQuarter
	}

	return 201, false, "", balanceSheetQuarter
}

func GetBalanceSheetAnnualGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_AnnualGrowth_Response) {

	balanceSheetAnnualGrowth, error_find_all := balanceSheet.Si_Find_AnnualGrowth(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching BalanceSheet Annual Growth data, details: " + error_find_all.Error(), balanceSheetAnnualGrowth
	}

	return 201, false, "", balanceSheetAnnualGrowth
}

func GetBalanceSheetQuarterGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.BalanceSheet_QuarterGrowth_Response) {

	balanceSheetQuarterGrowth, error_find_all := balanceSheet.Si_Find_QuarterGrowth(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching BalanceSheet Quarter Growth data, details: " + error_find_all.Error(), balanceSheetQuarterGrowth
	}

	return 201, false, "", balanceSheetQuarterGrowth
}

func GetCashFlowAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_Annual_Response) {

	cashFlowAnnual, error_find_all := cashFlow.Si_Find_Annual(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching CashFlow Annual data, details: " + error_find_all.Error(), cashFlowAnnual
	}

	return 201, false, "", cashFlowAnnual
}

func GetCashFlowQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_Quarter_Response) {

	cashFlowQuarter, error_find_all := cashFlow.Si_Find_Quarter(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching CashFlow Quarter data, details: " + error_find_all.Error(), cashFlowQuarter
	}

	return 201, false, "", cashFlowQuarter
}

func GetCashFlowAnnualGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_AnnualGrowth_Response) {

	cashFlowAnnualGrowth, error_find_all := cashFlow.Si_Find_AnnualGrowth(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching CashFlow Annual Growth data, details: " + error_find_all.Error(), cashFlowAnnualGrowth
	}

	return 201, false, "", cashFlowAnnualGrowth
}

func GetCashFlowQuarterGrowth_Service(symbol string, limit_int int) (int, bool, string, []models.CashFlow_QuarterGrowth_Response) {

	cashFlowQuarterGrowth, error_find_all := cashFlow.Si_Find_QuarterGrowth(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching CashFlow Quarter Growth data, details: " + error_find_all.Error(), cashFlowQuarterGrowth
	}

	return 201, false, "", cashFlowQuarterGrowth
}

func GetFinancialRatioAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.FinancialRatio_Annual_Response) {

	financialRatioAnnual, error_find_all := financialRatios.Si_Find_Annual(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching FinancialRatio Annual data, details: " + error_find_all.Error(), financialRatioAnnual
	}

	return 201, false, "", financialRatioAnnual
}

func GetFinancialRatioQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.FinancialRatio_Quarter_Response) {

	financialRatioQuarter, error_find_all := financialRatios.Si_Find_Quarter(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching FinancialRatio Quarter data, details: " + error_find_all.Error(), financialRatioQuarter
	}

	return 201, false, "", financialRatioQuarter
}

func GetFinancialRatioAnnualTTM_Service(symbol string, limit_int int) (int, bool, string, []models.FinancialRatio_AnnualTTM_Response) {

	financialRatioAnnualTTM, error_find_all := financialRatios.Si_Find_AnnualTTM(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching FinancialRatio Annual TTM data, details: " + error_find_all.Error(), financialRatioAnnualTTM
	}

	return 201, false, "", financialRatioAnnualTTM
}

func GetKeyMetricsAnnual_Service(symbol string, limit_int int) (int, bool, string, []models.KeyMetrics_Annual_Response) {

	keymetricsAnnual, error_find_all := keyMetrics.Si_Find_Annual(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching KeyMetrics Annual data, details: " + error_find_all.Error(), keymetricsAnnual
	}

	return 201, false, "", keymetricsAnnual
}

func GetKeyMetricsQuarter_Service(symbol string, limit_int int) (int, bool, string, []models.KeyMetrics_Quarter_Response) {

	keymetricsQuarter, error_find_all := keyMetrics.Si_Find_Quarter(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching KeyMetrics Quarter data, details: " + error_find_all.Error(), keymetricsQuarter
	}

	return 201, false, "", keymetricsQuarter
}

func GetKeyMetricsCompanyTTM_Service(symbol string, limit_int int) (int, bool, string, []models.KeyMetrics_CompanyTTM_Response) {

	keymetricsCompanyTTM, error_find_all := keyMetrics.Si_Find_companyTTM(symbol, limit_int)
	if error_find_all != nil {
		return 403, true, "Internal error when searching KeyMetrics Company TTM data, details: " + error_find_all.Error(), keymetricsCompanyTTM
	}

	return 201, false, "", keymetricsCompanyTTM
}

func GetAvailableTraded_Service() (int, bool, string, []models.TradableSymbols) {

	tradableSymbols, error_find_all := tradableSymbols.Si_Find()
	if error_find_all != nil {
		return 403, true, "Internal error when searching Tradable Symbols data, details: " + error_find_all.Error(), tradableSymbols
	}

	return 201, false, "", tradableSymbols
}

func GetCompanyProfile_Service(symbols string) (int, bool, string, models.CompanyProfile) {

	profile, error_find_all := profile.Si_Find_Profile(symbols)
	if error_find_all != nil {
		return 403, true, "Internal error when searching Profile data, details: " + error_find_all.Error(), profile
	}

	return 201, false, "", profile
}

func GetIndustryAndSector_Service() (int, bool, string, []models.IndustryAndSector) {

	industries_sectors, error_find_all := all.Si_Find_IndustriesAndSector()
	if error_find_all != nil {
		return 403, true, "Internal error when searching Industries and Sectors, details: " + error_find_all.Error(), industries_sectors
	}

	return 201, false, "", industries_sectors
}
