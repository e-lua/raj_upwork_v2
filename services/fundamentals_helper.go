package available_traded

import "github.com/Aphofisis/raj_upwork_v2/models"

/*-----RESPONSE DATA-----*/
type Response_IncomenStatement_Annual struct {
	Error     bool                            `json:"error"`
	DataError string                          `json:"dataError"`
	Data      []models.IncomeStatement_Annual `json:"data"`
}
type Response_IncomenStatement_Quarter struct {
	Error     bool                             `json:"error"`
	DataError string                           `json:"dataError"`
	Data      []models.IncomeStatement_Quarter `json:"data"`
}
type Response_IncomenStatement_AnnualGrowth struct {
	Error     bool                                  `json:"error"`
	DataError string                                `json:"dataError"`
	Data      []models.IncomeStatement_AnnualGrowth `json:"data"`
}
type Response_IncomenStatement_QuarterGrowth struct {
	Error     bool                                   `json:"error"`
	DataError string                                 `json:"dataError"`
	Data      []models.IncomeStatement_QuarterGrowth `json:"data"`
}

type Response_BalanceSheet_Annual struct {
	Error     bool                         `json:"error"`
	DataError string                       `json:"dataError"`
	Data      []models.BalanceSheet_Annual `json:"data"`
}
type Response_BalanceSheet_Quarter struct {
	Error     bool                          `json:"error"`
	DataError string                        `json:"dataError"`
	Data      []models.BalanceSheet_Quarter `json:"data"`
}
type Response_BalanceSheet_AnnualGrowth struct {
	Error     bool                               `json:"error"`
	DataError string                             `json:"dataError"`
	Data      []models.BalanceSheet_AnnualGrowth `json:"data"`
}
type Response_BalanceSheet_QuarterGrowth struct {
	Error     bool                                `json:"error"`
	DataError string                              `json:"dataError"`
	Data      []models.BalanceSheet_QuarterGrowth `json:"data"`
}

type Response_CashFlow_Annual struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      []models.CashFlow_Annual `json:"data"`
}
type Response_CashFlow_Quarter struct {
	Error     bool                      `json:"error"`
	DataError string                    `json:"dataError"`
	Data      []models.CashFlow_Quarter `json:"data"`
}
type Response_CashFlow_AnnualGrowth struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      []models.CashFlow_AnnualGrowth `json:"data"`
}
type Response_CashFlow_QuarterGrowth struct {
	Error     bool                            `json:"error"`
	DataError string                          `json:"dataError"`
	Data      []models.CashFlow_QuarterGrowth `json:"data"`
}

type Response_FinancialRatio_Annual struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      []models.FinancialRatio_Annual `json:"data"`
}
type Response_FinancialRatio_Quarter struct {
	Error     bool                            `json:"error"`
	DataError string                          `json:"dataError"`
	Data      []models.FinancialRatio_Quarter `json:"data"`
}
type Response_FinancialRatio_AnnualTTM struct {
	Error     bool                              `json:"error"`
	DataError string                            `json:"dataError"`
	Data      []models.FinancialRatio_AnnualTTM `json:"data"`
}

type Response_KeyMetrics_Annual struct {
	Error     bool                       `json:"error"`
	DataError string                     `json:"dataError"`
	Data      []models.KeyMetrics_Annual `json:"data"`
}
type Response_KeyMetrics_Quarter struct {
	Error     bool                        `json:"error"`
	DataError string                      `json:"dataError"`
	Data      []models.KeyMetrics_Quarter `json:"data"`
}
type Response_KeyMetrics_CompanyTTM struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      []models.KeyMetrics_CompanyTTM `json:"data"`
}

type Response_string struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

/*-----INCOMING DATA-----*/
type Incoming_NewData struct {
	Symbol             string `json:"symbol"`
	Api_token          string `json:"api_token"`
	WithTradedList     bool   `json:"with_tradedlist"`
	WithCompanyProfile bool   `json:"with_companyprofile"`
}
type Incoming_Request struct {
	Code          string `json:"code"`
	Currency_code string `json:"currency_code"`
}
