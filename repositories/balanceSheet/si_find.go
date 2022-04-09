package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_Annual(symbol string, limit int) ([]models.BalanceSheet_Annual, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
	symbol,
	reportedCurrency,
	cik,
	fillingDate,
	acceptedDate,
	calendarYear,
	period,
	cashAndCashEquivalents,
	shortTermInvestments,
	cashAndShortTermInvestments,
	netReceivables,
	inventory,
	otherCurrentAssets,
	totalCurrentAssets,
	propertyPlantEquipmentNet,
	goodwill,
	intangibleAssets,
	goodwillAndIntangibleAssets,
	longTermInvestments,
	taxAssets,
	otherNonCurrentAssets,
	totalNonCurrentAssets,
	otherAssets,
	totalAssets,
	accountPayables,
	shortTermDebt,
	taxPayables,
	deferredRevenue,
	otherCurrentLiabilities,
	totalCurrentLiabilities,
	longTermDebt,
	deferredRevenueNonCurrent,
	deferredTaxLiabilitiesNonCurrent,
	otherNonCurrentLiabilities,
	totalNonCurrentLiabilities,
	otherLiabilities,
	capitalLeaseObligations,
	totalLiabilities,
	preferredStock,
	commonStock,
	retainedEarnings,
	accumulatedOtherComprehensiveIncomeLoss,
	othertotalStockholdersEquity,
	totalStockholdersEquity,
	totalLiabilitiesAndStockholdersEquity,
	minorityInterest,
	totalEquity,
	totalLiabilitiesAndTotalEquity,
	totalInvestments,
	totalDebt,
	netDebt,
	link,
	finalLink FROM BalanceSheet_Annual WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListBSA := []models.BalanceSheet_Annual{}

	if error_show != nil {

		return oListBSA, error_show
	}

	for rows.Next() {
		oBSA := models.BalanceSheet_Annual{}
		rows.Scan(&oBSA.Date,
			&oBSA.Symbol,
			&oBSA.ReportedCurrency,
			&oBSA.Cik,
			&oBSA.FillingDate,
			&oBSA.AcceptedDate,
			&oBSA.CalendarYear,
			&oBSA.Period,
			&oBSA.CashAndCashEquivalents,
			&oBSA.ShortTermInvestments,
			&oBSA.CashAndShortTermInvestments,
			&oBSA.NetReceivables,
			&oBSA.Inventory,
			&oBSA.OtherCurrentAssets,
			&oBSA.TotalCurrentAssets,
			&oBSA.PropertyPlantEquipmentNet,
			&oBSA.Goodwill,
			&oBSA.IntangibleAssets,
			&oBSA.GoodwillAndIntangibleAssets,
			&oBSA.LongTermInvestments,
			&oBSA.TaxAssets,
			&oBSA.OtherNonCurrentAssets,
			&oBSA.TotalNonCurrentAssets,
			&oBSA.OtherAssets,
			&oBSA.TotalAssets,
			&oBSA.AccountPayables,
			&oBSA.ShortTermDebt,
			&oBSA.TaxPayables,
			&oBSA.DeferredRevenue,
			&oBSA.OtherCurrentLiabilities,
			&oBSA.TotalCurrentLiabilities,
			&oBSA.LongTermDebt,
			&oBSA.DeferredRevenueNonCurrent,
			&oBSA.DeferredTaxLiabilitiesNonCurrent,
			&oBSA.OtherNonCurrentLiabilities,
			&oBSA.TotalNonCurrentLiabilities,
			&oBSA.OtherLiabilities,
			&oBSA.CapitalLeaseObligations,
			&oBSA.TotalLiabilities,
			&oBSA.PreferredStock,
			&oBSA.CommonStock,
			&oBSA.RetainedEarnings,
			&oBSA.AccumulatedOtherComprehensiveIncomeLoss,
			&oBSA.OthertotalStockholdersEquity,
			&oBSA.TotalStockholdersEquity,
			&oBSA.TotalLiabilitiesAndStockholdersEquity,
			&oBSA.MinorityInterest,
			&oBSA.TotalEquity,
			&oBSA.TotalLiabilitiesAndTotalEquity,
			&oBSA.TotalInvestments,
			&oBSA.TotalDebt,
			&oBSA.NetDebt,
			&oBSA.Link,
			&oBSA.FinalLink)
		oListBSA = append(oListBSA, oBSA)
	}

	//OK
	return oListBSA, nil
}

func Si_Find_Quarter(symbol string, limit int) ([]models.BalanceSheet_Quarter, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
	symbol,
	reportedCurrency,
	cik,
	fillingDate,
	acceptedDate,
	calendarYear,
	period,
	cashAndCashEquivalents,
	shortTermInvestments,
	cashAndShortTermInvestments,
	netReceivables,
	inventory,
	otherCurrentAssets,
	totalCurrentAssets,
	propertyPlantEquipmentNet,
	goodwill,
	intangibleAssets,
	goodwillAndIntangibleAssets,
	longTermInvestments,
	taxAssets,
	otherNonCurrentAssets,
	totalNonCurrentAssets,
	otherAssets,
	totalAssets,
	accountPayables,
	shortTermDebt,
	taxPayables,
	deferredRevenue,
	otherCurrentLiabilities,
	totalCurrentLiabilities,
	longTermDebt,
	deferredRevenueNonCurrent,
	deferredTaxLiabilitiesNonCurrent,
	otherNonCurrentLiabilities,
	totalNonCurrentLiabilities,
	otherLiabilities,
	capitalLeaseObligations,
	totalLiabilities,
	preferredStock,
	commonStock,
	retainedEarnings,
	accumulatedOtherComprehensiveIncomeLoss,
	othertotalStockholdersEquity,
	totalStockholdersEquity,
	totalLiabilitiesAndStockholdersEquity,
	minorityInterest,
	totalEquity,
	totalLiabilitiesAndTotalEquity,
	totalInvestments,
	totalDebt,
	netDebt,
	link,
	finalLink FROM BalanceSheet_Quarter WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListBSQ := []models.BalanceSheet_Quarter{}

	if error_show != nil {

		return oListBSQ, error_show
	}

	for rows.Next() {
		oBSQ := models.BalanceSheet_Quarter{}
		rows.Scan(&oBSQ.Date,
			&oBSQ.Symbol,
			&oBSQ.ReportedCurrency,
			&oBSQ.Cik,
			&oBSQ.FillingDate,
			&oBSQ.AcceptedDate,
			&oBSQ.CalendarYear,
			&oBSQ.Period,
			&oBSQ.CashAndCashEquivalents,
			&oBSQ.ShortTermInvestments,
			&oBSQ.CashAndShortTermInvestments,
			&oBSQ.NetReceivables,
			&oBSQ.Inventory,
			&oBSQ.OtherCurrentAssets,
			&oBSQ.TotalCurrentAssets,
			&oBSQ.PropertyPlantEquipmentNet,
			&oBSQ.Goodwill,
			&oBSQ.IntangibleAssets,
			&oBSQ.GoodwillAndIntangibleAssets,
			&oBSQ.LongTermInvestments,
			&oBSQ.TaxAssets,
			&oBSQ.OtherNonCurrentAssets,
			&oBSQ.TotalNonCurrentAssets,
			&oBSQ.OtherAssets,
			&oBSQ.TotalAssets,
			&oBSQ.AccountPayables,
			&oBSQ.ShortTermDebt,
			&oBSQ.TaxPayables,
			&oBSQ.DeferredRevenue,
			&oBSQ.OtherCurrentLiabilities,
			&oBSQ.TotalCurrentLiabilities,
			&oBSQ.LongTermDebt,
			&oBSQ.DeferredRevenueNonCurrent,
			&oBSQ.DeferredTaxLiabilitiesNonCurrent,
			&oBSQ.OtherNonCurrentLiabilities,
			&oBSQ.TotalNonCurrentLiabilities,
			&oBSQ.OtherLiabilities,
			&oBSQ.CapitalLeaseObligations,
			&oBSQ.TotalLiabilities,
			&oBSQ.PreferredStock,
			&oBSQ.CommonStock,
			&oBSQ.RetainedEarnings,
			&oBSQ.AccumulatedOtherComprehensiveIncomeLoss,
			&oBSQ.OthertotalStockholdersEquity,
			&oBSQ.TotalStockholdersEquity,
			&oBSQ.TotalLiabilitiesAndStockholdersEquity,
			&oBSQ.MinorityInterest,
			&oBSQ.TotalEquity,
			&oBSQ.TotalLiabilitiesAndTotalEquity,
			&oBSQ.TotalInvestments,
			&oBSQ.TotalDebt,
			&oBSQ.NetDebt,
			&oBSQ.Link,
			&oBSQ.FinalLink)
		oListBSQ = append(oListBSQ, oBSQ)
	}

	//OK
	return oListBSQ, nil
}

func Si_Find_AnnualGrowth(symbol string, limit int) ([]models.BalanceSheet_AnnualGrowth, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
    symbol,
    period,
    growthCashAndCashEquivalents,
    growthShortTermInvestments,
    growthCashAndShortTermInvestments,
    growthNetReceivables,
    growthInventory,
    growthOtherCurrentAssets,
    growthTotalCurrentAssets,
    growthPropertyPlantEquipmentNet,
    growthGoodwill,
    growthIntangibleAssets,
    growthGoodwillAndIntangibleAssets,
    growthLongTermInvestments,
    growthTaxAssets,
    growthOtherNonCurrentAssets,
    growthTotalNonCurrentAssets,
    growthOtherAssets,
    growthTotalAssets,
    growthAccountPayables,
    growthShortTermDebt,
    growthTaxPayables,
    growthDeferredRevenue,
    growthOtherCurrentLiabilities,
    growthTotalCurrentLiabilities,
    growthLongTermDebt,
    growthDeferredRevenueNonCurrent,
    growthDeferrredTaxLiabilitiesNonCurrent,
    growthOtherNonCurrentLiabilities,
    growthTotalNonCurrentLiabilities,
    growthOtherLiabilities,
    growthTotalLiabilities,
    growthCommonStock,
    growthRetainedEarnings,
    growthAccumulatedOtherComprehensiveIncomeLoss,
    growthOthertotalStockholdersEquity,
    growthTotalStockholdersEquity,
    growthTotalLiabilitiesAndStockholdersEquity,
    growthTotalInvestments,
    growthTotalDebt,
    growthNetDebt FROM BalanceSheet_AnnualGrowth WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListBSAG := []models.BalanceSheet_AnnualGrowth{}

	if error_show != nil {

		return oListBSAG, error_show
	}

	for rows.Next() {
		oBSAG := models.BalanceSheet_AnnualGrowth{}
		rows.Scan(&oBSAG.Date,
			&oBSAG.Symbol,
			&oBSAG.Period,
			&oBSAG.GrowthCashAndCashEquivalents,
			&oBSAG.GrowthShortTermInvestments,
			&oBSAG.GrowthCashAndShortTermInvestments,
			&oBSAG.GrowthNetReceivables,
			&oBSAG.GrowthInventory,
			&oBSAG.GrowthOtherCurrentAssets,
			&oBSAG.GrowthTotalCurrentAssets,
			&oBSAG.GrowthPropertyPlantEquipmentNet,
			&oBSAG.GrowthGoodwill,
			&oBSAG.GrowthIntangibleAssets,
			&oBSAG.GrowthGoodwillAndIntangibleAssets,
			&oBSAG.GrowthLongTermInvestments,
			&oBSAG.GrowthTaxAssets,
			&oBSAG.GrowthOtherNonCurrentAssets,
			&oBSAG.GrowthTotalNonCurrentAssets,
			&oBSAG.GrowthOtherAssets,
			&oBSAG.GrowthTotalAssets,
			&oBSAG.GrowthAccountPayables,
			&oBSAG.GrowthShortTermDebt,
			&oBSAG.GrowthTaxPayables,
			&oBSAG.GrowthDeferredRevenue,
			&oBSAG.GrowthOtherCurrentLiabilities,
			&oBSAG.GrowthTotalCurrentLiabilities,
			&oBSAG.GrowthLongTermDebt,
			&oBSAG.GrowthDeferredRevenueNonCurrent,
			&oBSAG.GrowthDeferrredTaxLiabilitiesNonCurrent,
			&oBSAG.GrowthOtherNonCurrentLiabilities,
			&oBSAG.GrowthTotalNonCurrentLiabilities,
			&oBSAG.GrowthOtherLiabilities,
			&oBSAG.GrowthTotalLiabilities,
			&oBSAG.GrowthCommonStock,
			&oBSAG.GrowthRetainedEarnings,
			&oBSAG.GrowthAccumulatedOtherComprehensiveIncomeLoss,
			&oBSAG.GrowthOthertotalStockholdersEquity,
			&oBSAG.GrowthTotalStockholdersEquity,
			&oBSAG.GrowthTotalLiabilitiesAndStockholdersEquity,
			&oBSAG.GrowthTotalInvestments,
			&oBSAG.GrowthTotalDebt,
			&oBSAG.GrowthNetDebt)
		oListBSAG = append(oListBSAG, oBSAG)
	}

	//OK
	return oListBSAG, nil
}

func Si_Find_QuarterGrowth(symbol string, limit int) ([]models.BalanceSheet_QuarterGrowth, error) {

	db := models.SingleStoreCN
	q := `SELECT date,
    symbol,
    period,
    growthCashAndCashEquivalents,
    growthShortTermInvestments,
    growthCashAndShortTermInvestments,
    growthNetReceivables,
    growthInventory,
    growthOtherCurrentAssets,
    growthTotalCurrentAssets,
    growthPropertyPlantEquipmentNet,
    growthGoodwill,
    growthIntangibleAssets,
    growthGoodwillAndIntangibleAssets,
    growthLongTermInvestments,
    growthTaxAssets,
    growthOtherNonCurrentAssets,
    growthTotalNonCurrentAssets,
    growthOtherAssets,
    growthTotalAssets,
    growthAccountPayables,
    growthShortTermDebt,
    growthTaxPayables,
    growthDeferredRevenue,
    growthOtherCurrentLiabilities,
    growthTotalCurrentLiabilities,
    growthLongTermDebt,
    growthDeferredRevenueNonCurrent,
    growthDeferrredTaxLiabilitiesNonCurrent,
    growthOtherNonCurrentLiabilities,
    growthTotalNonCurrentLiabilities,
    growthOtherLiabilities,
    growthTotalLiabilities,
    growthCommonStock,
    growthRetainedEarnings,
    growthAccumulatedOtherComprehensiveIncomeLoss,
    growthOthertotalStockholdersEquity,
    growthTotalStockholdersEquity,
    growthTotalLiabilitiesAndStockholdersEquity,
    growthTotalInvestments,
    growthTotalDebt,
    growthNetDebt FROM BalanceSheet_QuarterGrowth WHERE symbol=? LIMIT ?`
	rows, error_show := db.Query(q, symbol, limit)

	oListBSQG := []models.BalanceSheet_QuarterGrowth{}

	if error_show != nil {

		return oListBSQG, error_show
	}

	for rows.Next() {
		oBSQG := models.BalanceSheet_QuarterGrowth{}
		rows.Scan(&oBSQG.Date,
			&oBSQG.Symbol,
			&oBSQG.Period,
			&oBSQG.GrowthCashAndCashEquivalents,
			&oBSQG.GrowthShortTermInvestments,
			&oBSQG.GrowthCashAndShortTermInvestments,
			&oBSQG.GrowthNetReceivables,
			&oBSQG.GrowthInventory,
			&oBSQG.GrowthOtherCurrentAssets,
			&oBSQG.GrowthTotalCurrentAssets,
			&oBSQG.GrowthPropertyPlantEquipmentNet,
			&oBSQG.GrowthGoodwill,
			&oBSQG.GrowthIntangibleAssets,
			&oBSQG.GrowthGoodwillAndIntangibleAssets,
			&oBSQG.GrowthLongTermInvestments,
			&oBSQG.GrowthTaxAssets,
			&oBSQG.GrowthOtherNonCurrentAssets,
			&oBSQG.GrowthTotalNonCurrentAssets,
			&oBSQG.GrowthOtherAssets,
			&oBSQG.GrowthTotalAssets,
			&oBSQG.GrowthAccountPayables,
			&oBSQG.GrowthShortTermDebt,
			&oBSQG.GrowthTaxPayables,
			&oBSQG.GrowthDeferredRevenue,
			&oBSQG.GrowthOtherCurrentLiabilities,
			&oBSQG.GrowthTotalCurrentLiabilities,
			&oBSQG.GrowthLongTermDebt,
			&oBSQG.GrowthDeferredRevenueNonCurrent,
			&oBSQG.GrowthDeferrredTaxLiabilitiesNonCurrent,
			&oBSQG.GrowthOtherNonCurrentLiabilities,
			&oBSQG.GrowthTotalNonCurrentLiabilities,
			&oBSQG.GrowthOtherLiabilities,
			&oBSQG.GrowthTotalLiabilities,
			&oBSQG.GrowthCommonStock,
			&oBSQG.GrowthRetainedEarnings,
			&oBSQG.GrowthAccumulatedOtherComprehensiveIncomeLoss,
			&oBSQG.GrowthOthertotalStockholdersEquity,
			&oBSQG.GrowthTotalStockholdersEquity,
			&oBSQG.GrowthTotalLiabilitiesAndStockholdersEquity,
			&oBSQG.GrowthTotalInvestments,
			&oBSQG.GrowthTotalDebt,
			&oBSQG.GrowthNetDebt)
		oListBSQG = append(oListBSQG, oBSQG)
	}

	//OK
	return oListBSQG, nil
}
