package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_Annual(symbol string, limit int) ([]models.CashFlow_Annual, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
	symbol,
	reportedCurrency,
	cik,
	fillingDate,
	acceptedDate,
	calendarYear,
	period,
	netIncome,
	depreciationAndAmortization,
	deferredIncomeTax,
	stockBasedCompensation,
	changeInWorkingCapital,
	accountsReceivables,
	inventory,
	accountsPayables,
	otherWorkingCapital,
	otherNonCashItems,
	netCashProvidedByOperatingActivities,
	investmentsInPropertyPlantAndEquipment,
	acquisitionsNet,
	purchasesOfInvestments,
	salesMaturitiesOfInvestments,
	otherInvestingActivites,
	netCashUsedForInvestingActivites,
	debtRepayment,
	commonStockIssued,
	commonStockRepurchased,
	dividendsPaid,
	otherFinancingActivites,
	netCashUsedProvidedByFinancingActivities,
	effectOfForexChangesOnCash,
	netChangeInCash,
	cashAtEndOfPeriod,
	cashAtBeginningOfPeriod,
	operatingCashFlow,
	capitalExpenditure,
	freeCashFlow,
	link,
	finalLink FROM CashFlow_Annual WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListCFA := []models.CashFlow_Annual{}

	if error_show != nil {

		return oListCFA, error_show
	}

	for rows.Next() {
		oCFA := models.CashFlow_Annual{}
		rows.Scan(&oCFA.Date,
			&oCFA.Symbol,
			&oCFA.ReportedCurrency,
			&oCFA.Cik,
			&oCFA.FillingDate,
			&oCFA.AcceptedDate,
			&oCFA.CalendarYear,
			&oCFA.Period,
			&oCFA.NetIncome,
			&oCFA.DepreciationAndAmortization,
			&oCFA.DeferredIncomeTax,
			&oCFA.StockBasedCompensation,
			&oCFA.ChangeInWorkingCapital,
			&oCFA.AccountsReceivables,
			&oCFA.Inventory,
			&oCFA.AccountsPayables,
			&oCFA.OtherWorkingCapital,
			&oCFA.OtherNonCashItems,
			&oCFA.NetCashProvidedByOperatingActivities,
			&oCFA.InvestmentsInPropertyPlantAndEquipment,
			&oCFA.AcquisitionsNet,
			&oCFA.PurchasesOfInvestments,
			&oCFA.SalesMaturitiesOfInvestments,
			&oCFA.OtherInvestingActivites,
			&oCFA.NetCashUsedForInvestingActivites,
			&oCFA.DebtRepayment,
			&oCFA.CommonStockIssued,
			&oCFA.CommonStockRepurchased,
			&oCFA.DividendsPaid,
			&oCFA.OtherFinancingActivites,
			&oCFA.NetCashUsedProvidedByFinancingActivities,
			&oCFA.EffectOfForexChangesOnCash,
			&oCFA.NetChangeInCash,
			&oCFA.CashAtEndOfPeriod,
			&oCFA.CashAtBeginningOfPeriod,
			&oCFA.OperatingCashFlow,
			&oCFA.CapitalExpenditure,
			&oCFA.FreeCashFlow,
			&oCFA.Link,
			&oCFA.FinalLink)
		oListCFA = append(oListCFA, oCFA)
	}

	//OK
	return oListCFA, nil
}

func Si_Find_Quarter(symbol string, limit int) ([]models.CashFlow_Quarter, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
	symbol,
	reportedCurrency,
	cik,
	fillingDate,
	acceptedDate,
	calendarYear,
	period,
	netIncome,
	depreciationAndAmortization,
	deferredIncomeTax,
	stockBasedCompensation,
	changeInWorkingCapital,
	accountsReceivables,
	inventory,
	accountsPayables,
	otherWorkingCapital,
	otherNonCashItems,
	netCashProvidedByOperatingActivities,
	investmentsInPropertyPlantAndEquipment,
	acquisitionsNet,
	purchasesOfInvestments,
	salesMaturitiesOfInvestments,
	otherInvestingActivites,
	netCashUsedForInvestingActivites,
	debtRepayment,
	commonStockIssued,
	commonStockRepurchased,
	dividendsPaid,
	otherFinancingActivites,
	netCashUsedProvidedByFinancingActivities,
	effectOfForexChangesOnCash,
	netChangeInCash,
	cashAtEndOfPeriod,
	cashAtBeginningOfPeriod,
	operatingCashFlow,
	capitalExpenditure,
	freeCashFlow,
	link,
	finalLink FROM CashFlow_Quarter WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListCFQ := []models.CashFlow_Quarter{}

	if error_show != nil {

		return oListCFQ, error_show
	}

	for rows.Next() {
		oCFQ := models.CashFlow_Quarter{}
		rows.Scan(&oCFQ.Date,
			&oCFQ.Symbol,
			&oCFQ.ReportedCurrency,
			&oCFQ.Cik,
			&oCFQ.FillingDate,
			&oCFQ.AcceptedDate,
			&oCFQ.CalendarYear,
			&oCFQ.Period,
			&oCFQ.NetIncome,
			&oCFQ.DepreciationAndAmortization,
			&oCFQ.DeferredIncomeTax,
			&oCFQ.StockBasedCompensation,
			&oCFQ.ChangeInWorkingCapital,
			&oCFQ.AccountsReceivables,
			&oCFQ.Inventory,
			&oCFQ.AccountsPayables,
			&oCFQ.OtherWorkingCapital,
			&oCFQ.OtherNonCashItems,
			&oCFQ.NetCashProvidedByOperatingActivities,
			&oCFQ.InvestmentsInPropertyPlantAndEquipment,
			&oCFQ.AcquisitionsNet,
			&oCFQ.PurchasesOfInvestments,
			&oCFQ.SalesMaturitiesOfInvestments,
			&oCFQ.OtherInvestingActivites,
			&oCFQ.NetCashUsedForInvestingActivites,
			&oCFQ.DebtRepayment,
			&oCFQ.CommonStockIssued,
			&oCFQ.CommonStockRepurchased,
			&oCFQ.DividendsPaid,
			&oCFQ.OtherFinancingActivites,
			&oCFQ.NetCashUsedProvidedByFinancingActivities,
			&oCFQ.EffectOfForexChangesOnCash,
			&oCFQ.NetChangeInCash,
			&oCFQ.CashAtEndOfPeriod,
			&oCFQ.CashAtBeginningOfPeriod,
			&oCFQ.OperatingCashFlow,
			&oCFQ.CapitalExpenditure,
			&oCFQ.FreeCashFlow,
			&oCFQ.Link,
			&oCFQ.FinalLink)
		oListCFQ = append(oListCFQ, oCFQ)
	}

	//OK
	return oListCFQ, nil
}

func Si_Find_AnnualGrowth(symbol string, limit int) ([]models.CashFlow_AnnualGrowth, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
    symbol,
    period,
    growthNetIncome,
    growthDepreciationAndAmortization,
    growthDeferredIncomeTax,
    growthStockBasedCompensation,
    growthChangeInWorkingCapital,
    growthAccountsReceivables,
    growthInventory,
    growthAccountsPayables,
    growthOtherWorkingCapital,
    growthOtherNonCashItems,
    growthNetCashProvidedByOperatingActivites,
    growthInvestmentsInPropertyPlantAndEquipment,
    growthAcquisitionsNet,
    growthPurchasesOfInvestments,
    growthSalesMaturitiesOfInvestments,
    growthOtherInvestingActivites,
    growthNetCashUsedForInvestingActivites,
    growthDebtRepayment,
    growthCommonStockIssued,
    growthCommonStockRepurchased,
    growthDividendsPaid,
    growthOtherFinancingActivites,
    growthNetCashUsedProvidedByFinancingActivities,
    growthEffectOfForexChangesOnCash,
    growthNetChangeInCash,
    growthCashAtEndOfPeriod,
    growthCashAtBeginningOfPeriod,
    growthOperatingCashFlow,
    growthCapitalExpenditure,
    growthFreeCashFlow FROM CashFlow_AnnualGrowth WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListCFAG := []models.CashFlow_AnnualGrowth{}

	if error_show != nil {

		return oListCFAG, error_show
	}

	for rows.Next() {
		oCFAG := models.CashFlow_AnnualGrowth{}
		rows.Scan(&oCFAG.Date,
			&oCFAG.Symbol,
			&oCFAG.Period,
			&oCFAG.GrowthNetIncome,
			&oCFAG.GrowthDepreciationAndAmortization,
			&oCFAG.GrowthDeferredIncomeTax,
			&oCFAG.GrowthStockBasedCompensation,
			&oCFAG.GrowthChangeInWorkingCapital,
			&oCFAG.GrowthAccountsReceivables,
			&oCFAG.GrowthInventory,
			&oCFAG.GrowthAccountsPayables,
			&oCFAG.GrowthOtherWorkingCapital,
			&oCFAG.GrowthOtherNonCashItems,
			&oCFAG.GrowthNetCashProvidedByOperatingActivites,
			&oCFAG.GrowthInvestmentsInPropertyPlantAndEquipment,
			&oCFAG.GrowthAcquisitionsNet,
			&oCFAG.GrowthPurchasesOfInvestments,
			&oCFAG.GrowthSalesMaturitiesOfInvestments,
			&oCFAG.GrowthOtherInvestingActivites,
			&oCFAG.GrowthNetCashUsedForInvestingActivites,
			&oCFAG.GrowthDebtRepayment,
			&oCFAG.GrowthCommonStockIssued,
			&oCFAG.GrowthCommonStockRepurchased,
			&oCFAG.GrowthDividendsPaid,
			&oCFAG.GrowthOtherFinancingActivites,
			&oCFAG.GrowthNetCashUsedProvidedByFinancingActivities,
			&oCFAG.GrowthEffectOfForexChangesOnCash,
			&oCFAG.GrowthNetChangeInCash,
			&oCFAG.GrowthCashAtEndOfPeriod,
			&oCFAG.GrowthCashAtBeginningOfPeriod,
			&oCFAG.GrowthOperatingCashFlow,
			&oCFAG.GrowthCapitalExpenditure,
			&oCFAG.GrowthFreeCashFlow)
		oListCFAG = append(oListCFAG, oCFAG)
	}

	//OK
	return oListCFAG, nil
}

func Si_Find_QuarterGrowth(symbol string, limit int) ([]models.CashFlow_QuarterGrowth, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
    symbol,
    period,
    growthNetIncome,
    growthDepreciationAndAmortization,
    growthDeferredIncomeTax,
    growthStockBasedCompensation,
    growthChangeInWorkingCapital,
    growthAccountsReceivables,
    growthInventory,
    growthAccountsPayables,
    growthOtherWorkingCapital,
    growthOtherNonCashItems,
    growthNetCashProvidedByOperatingActivites,
    growthInvestmentsInPropertyPlantAndEquipment,
    growthAcquisitionsNet,
    growthPurchasesOfInvestments,
    growthSalesMaturitiesOfInvestments,
    growthOtherInvestingActivites,
    growthNetCashUsedForInvestingActivites,
    growthDebtRepayment,
    growthCommonStockIssued,
    growthCommonStockRepurchased,
    growthDividendsPaid,
    growthOtherFinancingActivites,
    growthNetCashUsedProvidedByFinancingActivities,
    growthEffectOfForexChangesOnCash,
    growthNetChangeInCash,
    growthCashAtEndOfPeriod,
    growthCashAtBeginningOfPeriod,
    growthOperatingCashFlow,
    growthCapitalExpenditure,
    growthFreeCashFlow FROM CashFlow_QuarterGrowth WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListCFQG := []models.CashFlow_QuarterGrowth{}

	if error_show != nil {

		return oListCFQG, error_show
	}

	for rows.Next() {
		oCFQG := models.CashFlow_QuarterGrowth{}
		rows.Scan(&oCFQG.Date,
			&oCFQG.Symbol,
			&oCFQG.Period,
			&oCFQG.GrowthNetIncome,
			&oCFQG.GrowthDepreciationAndAmortization,
			&oCFQG.GrowthDeferredIncomeTax,
			&oCFQG.GrowthStockBasedCompensation,
			&oCFQG.GrowthChangeInWorkingCapital,
			&oCFQG.GrowthAccountsReceivables,
			&oCFQG.GrowthInventory,
			&oCFQG.GrowthAccountsPayables,
			&oCFQG.GrowthOtherWorkingCapital,
			&oCFQG.GrowthOtherNonCashItems,
			&oCFQG.GrowthNetCashProvidedByOperatingActivites,
			&oCFQG.GrowthInvestmentsInPropertyPlantAndEquipment,
			&oCFQG.GrowthAcquisitionsNet,
			&oCFQG.GrowthPurchasesOfInvestments,
			&oCFQG.GrowthSalesMaturitiesOfInvestments,
			&oCFQG.GrowthOtherInvestingActivites,
			&oCFQG.GrowthNetCashUsedForInvestingActivites,
			&oCFQG.GrowthDebtRepayment,
			&oCFQG.GrowthCommonStockIssued,
			&oCFQG.GrowthCommonStockRepurchased,
			&oCFQG.GrowthDividendsPaid,
			&oCFQG.GrowthOtherFinancingActivites,
			&oCFQG.GrowthNetCashUsedProvidedByFinancingActivities,
			&oCFQG.GrowthEffectOfForexChangesOnCash,
			&oCFQG.GrowthNetChangeInCash,
			&oCFQG.GrowthCashAtEndOfPeriod,
			&oCFQG.GrowthCashAtBeginningOfPeriod,
			&oCFQG.GrowthOperatingCashFlow,
			&oCFQG.GrowthCapitalExpenditure,
			&oCFQG.GrowthFreeCashFlow)
		oListCFQG = append(oListCFQG, oCFQG)
	}

	//OK
	return oListCFQG, nil
}
