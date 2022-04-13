package all

import (
	"log"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Delete(symbol string) error {

	//BEGIN
	tx, error_tx := models.SingleStoreCN.Begin()
	if error_tx != nil {
		tx.Rollback()
		return error_tx
	}

	//Company profile
	q_cop := `DELETE FROM CompanyProfile WHERE symbol=?`
	_, error_delete_cop := tx.Query(q_cop, symbol)
	if error_delete_cop != nil {
		tx.Rollback()
		return error_delete_cop
	}
	log.Print("DELETED COMPANY PROFILE....Done")

	//Income Statement Annual
	q_isa := `DELETE FROM IncomeStatement_Annual WHERE symbol=?`
	_, error_delete_isa := tx.Query(q_isa, symbol)
	if error_delete_isa != nil {
		tx.Rollback()
		return error_delete_isa
	}
	log.Print("DELETED INCOME STATEMENT ANNUAL....Done")

	//Income Statement Quarter
	q_isq := `DELETE FROM IncomeStatement_Quarter WHERE symbol=?`
	_, error_delete_isq := tx.Query(q_isq, symbol)
	if error_delete_isq != nil {
		tx.Rollback()
		return error_delete_isq
	}
	log.Print("DELETED INCOME STATEMENT QUARTER....Done")

	//Income Statement Annual Growth
	q_isag := `DELETE FROM IncomeStatement_AnnualGrowth WHERE symbol=?`
	_, error_delete_isag := tx.Query(q_isag, symbol)
	if error_delete_isag != nil {
		tx.Rollback()
		return error_delete_isag
	}
	log.Print("DELETED INCOME STATEMENT ANNUAL GROWTH....Done")

	//Income Statement Quarter Growth
	q_isqg := `DELETE FROM IncomeStatement_QuarterGrowth WHERE symbol=?`
	_, error_delete_isqg := tx.Query(q_isqg, symbol)
	if error_delete_isqg != nil {
		tx.Rollback()
		return error_delete_isqg
	}
	log.Print("DELETED INCOME STATEMENT QUARTER GROWTH....Done")

	//Balance Sheet Annual
	q_bsa := `DELETE FROM BalanceSheet_Annual WHERE symbol=?`
	_, error_delete_bsa := tx.Query(q_bsa, symbol)
	if error_delete_bsa != nil {
		tx.Rollback()
		return error_delete_bsa
	}
	log.Print("DELETED BALANCE SHEET ANNUAL....Done")

	//Balance Sheet Quarter
	q_bsq := `DELETE FROM BalanceSheet_Quarter WHERE symbol=?`
	_, error_delete_bsq := tx.Query(q_bsq, symbol)
	if error_delete_bsq != nil {
		tx.Rollback()
		return error_delete_bsq
	}
	log.Print("DELETED BALANCE SHEET QUARTER....Done")

	//Balance Sheet Annual Growth
	q_bsag := `DELETE FROM BalanceSheet_AnnualGrowth WHERE symbol=?`
	_, error_delete_bsag := tx.Query(q_bsag, symbol)
	if error_delete_bsag != nil {
		tx.Rollback()
		return error_delete_bsag
	}
	log.Print("DELETED BALANCE SHEET ANNUAL GROWTH....Done")

	//Balance Sheet Quarter Growth
	q_bsqg := `DELETE FROM BalanceSheet_QuarterGrowth WHERE symbol=?`
	_, error_delete_bsqg := tx.Query(q_bsqg, symbol)
	if error_delete_bsqg != nil {
		tx.Rollback()
		return error_delete_bsqg
	}
	log.Print("DELETED BALANCE SHEET QUARTER GROWTH....Done")

	//CashFlow Annual
	q_cfa := `DELETE FROM CashFlow_Annual WHERE symbol=?`
	_, error_delete_cfa := tx.Query(q_cfa, symbol)
	if error_delete_cfa != nil {
		tx.Rollback()
		return error_delete_cfa
	}
	log.Print("DELETED CASHFLOW ANNUAL....Done")

	//CashFlow Quarter
	q_cfq := `DELETE FROM CashFlow_Quarter WHERE symbol=?`
	_, error_delete_cfq := tx.Query(q_cfq, symbol)
	if error_delete_cfq != nil {
		tx.Rollback()
		return error_delete_cfq
	}
	log.Print("DELETED CASHFLOW QUARTER....Done")

	//CashFlow Annual Growth
	q_cfag := `DELETE FROM CashFlow_AnnualGrowth WHERE symbol=?`
	_, error_delete_cfag := tx.Query(q_cfag, symbol)
	if error_delete_cfag != nil {
		tx.Rollback()
		return error_delete_cfag
	}
	log.Print("DELETED CASHFLOW ANNUAL GROWTH....Done")

	//CashFlow Quarter Growth
	q_cfqg := `DELETE FROM CashFlow_QuarterGrowth WHERE symbol=?`
	_, error_delete_cfqg := tx.Query(q_cfqg, symbol)
	if error_delete_cfqg != nil {
		tx.Rollback()
		return error_delete_cfqg
	}
	log.Print("DELETED CASHFLOW QUARTER GROWTH....Done")

	//Financial Ratio Annual
	q_fra := `DELETE FROM FinancialRatio_Annual WHERE symbol=?`
	_, error_delete_fra := tx.Query(q_fra, symbol)
	if error_delete_fra != nil {
		tx.Rollback()
		return error_delete_fra
	}
	log.Print("DELETED FINANCIAL RATIO ANNUAL....Done")

	//Financial Ratio Quarter
	q_frq := `DELETE FROM FinancialRatio_Quarter WHERE symbol=?`
	_, error_delete_frq := tx.Query(q_frq, symbol)
	if error_delete_frq != nil {
		tx.Rollback()
		return error_delete_frq
	}
	log.Print("DELETED FINANCIAL RATIO QUARTER....Done")

	//Financial Ratio Annual TTM
	q_frattm := `DELETE FROM FinancialRatio_AnnualTTM WHERE symbol=?`
	_, error_delete_frattm := tx.Query(q_frattm, symbol)
	if error_delete_frattm != nil {
		tx.Rollback()
		return error_delete_frattm
	}
	log.Print("DELETED FINANCIAL RATIO ANNUAL TTM....Done")

	//Key Metrics Annual
	q_kma := `DELETE FROM KeyMetrics_Annual WHERE symbol=?`
	_, error_delete_kma := tx.Query(q_kma, symbol)
	if error_delete_kma != nil {
		tx.Rollback()
		return error_delete_kma
	}
	log.Print("DELETED KEY METRICS ANNUAL....Done")

	//Key Metrics Quarter
	q_kmq := `DELETE FROM KeyMetrics_Quarter WHERE symbol=?`
	_, error_delete_kmq := tx.Query(q_kmq, symbol)
	if error_delete_kmq != nil {
		tx.Rollback()
		return error_delete_kmq
	}
	log.Print("DELETED KEY METRICS QUARTER....Done")

	//Key Metrics Company TTM
	q_kmcttm := `DELETE FROM KeyMetrics_CompanyTTM WHERE symbol=?`
	_, error_delete_kmcttm := tx.Query(q_kmcttm, symbol)
	if error_delete_kmcttm != nil {
		tx.Rollback()
		return error_delete_kmcttm
	}
	log.Print("DELETED KEY METRICS COMPANY TTM....Done")

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit()
	if err_commit != nil {
		tx.Rollback()
		return err_commit
	}

	log.Print("-------->SUCCESSFUL DELETING DATA PROCESS")

	return nil
}
