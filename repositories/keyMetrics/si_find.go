package all

import (
	"strconv"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_Annual(symbol string, limit int) ([]models.KeyMetrics_Annual, error) {

	db := models.SingleStoreCN
	q := `SELECT symbol,date,period,revenuePerShare,netIncomePerShare,operatingCashFlowPerShare,freeCashFlowPerShare,cashPerShare,bookValuePerShare,tangibleBookValuePerShare,shareholdersEquityPerShare,interestDebtPerShare,marketCap,enterpriseValue,peRatio,priceToSalesRatio,pocfratio,pfcfRatio,pbRatio,ptbRatio,evToSales,enterpriseValueOverEBITDA,evToOperatingCashFlow,evToFreeCashFlow,earningsYield,freeCashFlowYield,debtToEquity,debtToAssets,netDebtToEBITDA,currentRatio,interestCoverage,incomeQuality,dividendYield,payoutRatio,salesGeneralAndAdministrativeToRevenue,researchAndDdevelopementToRevenue,intangiblesToTotalAssets,capexToOperatingCashFlow,capexToRevenue,capexToDepreciation,stockBasedCompensationToRevenue,grahamNumber,roic,returnOnTangibleAssets,grahamNetNet,workingCapital,tangibleAssetValue,netCurrentAssetValue,investedCapital,averageReceivables,averagePayables,averageInventory,daysSalesOutstanding,daysPayablesOutstanding,daysOfInventoryOnHand,receivablesTurnover,payablesTurnover,inventoryTurnover,roe,capexPerShare FROM KeyMetrics_Annual WHERE symbol=? LIMIT ` + strconv.Itoa(limit)
	rows, error_show := db.Query(q, symbol)

	oListKMA := []models.KeyMetrics_Annual{}

	if error_show != nil {

		return oListKMA, error_show
	}

	for rows.Next() {
		oKMA := models.KeyMetrics_Annual{}
		rows.Scan(&oKMA.Symbol,
			&oKMA.Date,
			&oKMA.Period,
			&oKMA.RevenuePerShare,
			&oKMA.NetIncomePerShare,
			&oKMA.OperatingCashFlowPerShare,
			&oKMA.FreeCashFlowPerShare,
			&oKMA.CashPerShare,
			&oKMA.BookValuePerShare,
			&oKMA.TangibleBookValuePerShare,
			&oKMA.ShareholdersEquityPerShare,
			&oKMA.InterestDebtPerShare,
			&oKMA.MarketCap,
			&oKMA.EnterpriseValue,
			&oKMA.PeRatio,
			&oKMA.PriceToSalesRatio,
			&oKMA.Pocfratio,
			&oKMA.PfcfRatio,
			&oKMA.PbRatio,
			&oKMA.PtbRatio,
			&oKMA.EvToSales,
			&oKMA.EnterpriseValueOverEBITDA,
			&oKMA.EvToOperatingCashFlow,
			&oKMA.EvToFreeCashFlow,
			&oKMA.EarningsYield,
			&oKMA.FreeCashFlowYield,
			&oKMA.DebtToEquity,
			&oKMA.DebtToAssets,
			&oKMA.NetDebtToEBITDA,
			&oKMA.CurrentRatio,
			&oKMA.InterestCoverage,
			&oKMA.IncomeQuality,
			&oKMA.DividendYield,
			&oKMA.PayoutRatio,
			&oKMA.SalesGeneralAndAdministrativeToRevenue,
			&oKMA.ResearchAndDdevelopementToRevenue,
			&oKMA.IntangiblesToTotalAssets,
			&oKMA.CapexToOperatingCashFlow,
			&oKMA.CapexToRevenue,
			&oKMA.CapexToDepreciation,
			&oKMA.StockBasedCompensationToRevenue,
			&oKMA.GrahamNumber,
			&oKMA.Roic,
			&oKMA.ReturnOnTangibleAssets,
			&oKMA.GrahamNetNet,
			&oKMA.WorkingCapital,
			&oKMA.TangibleAssetValue,
			&oKMA.NetCurrentAssetValue,
			&oKMA.InvestedCapital,
			&oKMA.AverageReceivables,
			&oKMA.AveragePayables,
			&oKMA.AverageInventory,
			&oKMA.DaysSalesOutstanding,
			&oKMA.DaysPayablesOutstanding,
			&oKMA.DaysOfInventoryOnHand,
			&oKMA.ReceivablesTurnover,
			&oKMA.PayablesTurnover,
			&oKMA.InventoryTurnover,
			&oKMA.Roe,
			&oKMA.CapexPerShare)
		oListKMA = append(oListKMA, oKMA)
	}

	//OK
	return oListKMA, nil
}

func Si_Find_Quarter(symbol string, limit int) ([]models.KeyMetrics_Quarter, error) {

	db := models.SingleStoreCN
	q := "SELECT symbol,date,period,revenuePerShare,netIncomePerShare,operatingCashFlowPerShare,freeCashFlowPerShare,cashPerShare,bookValuePerShare,tangibleBookValuePerShare,shareholdersEquityPerShare,interestDebtPerShare,marketCap,enterpriseValue,peRatio,priceToSalesRatio,pocfratio,pfcfRatio,pbRatio,ptbRatio,evToSales,enterpriseValueOverEBITDA,evToOperatingCashFlow,evToFreeCashFlow,earningsYield,freeCashFlowYield,debtToEquity,debtToAssets,netDebtToEBITDA,currentRatio,interestCoverage,incomeQuality,dividendYield,payoutRatio,salesGeneralAndAdministrativeToRevenue,researchAndDdevelopementToRevenue,intangiblesToTotalAssets,capexToOperatingCashFlow,capexToRevenue,capexToDepreciation,stockBasedCompensationToRevenue,grahamNumber,roic,returnOnTangibleAssets,grahamNetNet,workingCapital,tangibleAssetValue,netCurrentAssetValue,investedCapital,averageReceivables,averagePayables,averageInventory,daysSalesOutstanding,daysPayablesOutstanding,daysOfInventoryOnHand,receivablesTurnover,payablesTurnover,inventoryTurnover,roe,capexPerShare FROM KeyMetrics_Annual WHERE symbol = ? LIMIT ?"
	rows, error_show := db.Query(q, symbol, limit)

	oListKMQ := []models.KeyMetrics_Quarter{}

	if error_show != nil {

		return oListKMQ, error_show
	}

	for rows.Next() {
		oKMQ := models.KeyMetrics_Quarter{}
		rows.Scan(&oKMQ.Symbol,
			&oKMQ.Date,
			&oKMQ.Period,
			&oKMQ.RevenuePerShare,
			&oKMQ.NetIncomePerShare,
			&oKMQ.OperatingCashFlowPerShare,
			&oKMQ.FreeCashFlowPerShare,
			&oKMQ.CashPerShare,
			&oKMQ.BookValuePerShare,
			&oKMQ.TangibleBookValuePerShare,
			&oKMQ.ShareholdersEquityPerShare,
			&oKMQ.InterestDebtPerShare,
			&oKMQ.MarketCap,
			&oKMQ.EnterpriseValue,
			&oKMQ.PeRatio,
			&oKMQ.PriceToSalesRatio,
			&oKMQ.Pocfratio,
			&oKMQ.PfcfRatio,
			&oKMQ.PbRatio,
			&oKMQ.PtbRatio,
			&oKMQ.EvToSales,
			&oKMQ.EnterpriseValueOverEBITDA,
			&oKMQ.EvToOperatingCashFlow,
			&oKMQ.EvToFreeCashFlow,
			&oKMQ.EarningsYield,
			&oKMQ.FreeCashFlowYield,
			&oKMQ.DebtToEquity,
			&oKMQ.DebtToAssets,
			&oKMQ.NetDebtToEBITDA,
			&oKMQ.CurrentRatio,
			&oKMQ.InterestCoverage,
			&oKMQ.IncomeQuality,
			&oKMQ.DividendYield,
			&oKMQ.PayoutRatio,
			&oKMQ.SalesGeneralAndAdministrativeToRevenue,
			&oKMQ.ResearchAndDdevelopementToRevenue,
			&oKMQ.IntangiblesToTotalAssets,
			&oKMQ.CapexToOperatingCashFlow,
			&oKMQ.CapexToRevenue,
			&oKMQ.CapexToDepreciation,
			&oKMQ.StockBasedCompensationToRevenue,
			&oKMQ.GrahamNumber,
			&oKMQ.Roic,
			&oKMQ.ReturnOnTangibleAssets,
			&oKMQ.GrahamNetNet,
			&oKMQ.WorkingCapital,
			&oKMQ.TangibleAssetValue,
			&oKMQ.NetCurrentAssetValue,
			&oKMQ.InvestedCapital,
			&oKMQ.AverageReceivables,
			&oKMQ.AveragePayables,
			&oKMQ.AverageInventory,
			&oKMQ.DaysSalesOutstanding,
			&oKMQ.DaysPayablesOutstanding,
			&oKMQ.DaysOfInventoryOnHand,
			&oKMQ.ReceivablesTurnover,
			&oKMQ.PayablesTurnover,
			&oKMQ.InventoryTurnover,
			&oKMQ.Roe,
			&oKMQ.CapexPerShare)
		oListKMQ = append(oListKMQ, oKMQ)
	}

	//OK
	return oListKMQ, nil
}

func Si_Find_companyTTM(symbol string, limit int) ([]models.KeyMetrics_CompanyTTM, error) {

	db := models.SingleStoreCN
	q := "SELECT symbol,revenuePerShareTTM,netIncomePerShareTTM,operatingCashFlowPerShareTTM,freeCashFlowPerShareTTM,cashPerShareTTM,bookValuePerShareTTM,tangibleBookValuePerShareTTM,shareholdersEquityPerShareTTM,interestDebtPerShareTTM,marketCapTTM,enterpriseValueTTM,peRatioTTM,priceToSalesRatioTTM,pocfratioTTM,pfcfRatioTTM,pbRatioTTM,ptbRatioTTM,evToSalesTTM,enterpriseValueOverEBITDATTM,evToOperatingCashFlowTTM,evToFreeCashFlowTTM,earningsYieldTTM,freeCashFlowYieldTTM,debtToEquityTTM,debtToAssetsTTM,netDebtToEBITDATTM,currentRatioTTM,interestCoverageTTM,incomeQualityTTM,dividendYieldTTM,dividendYieldPercentageTTM,payoutRatioTTM,salesGeneralAndAdministrativeToRevenueTTM,researchAndDevelopementToRevenueTTM,intangiblesToTotalAssetsTTM,capexToOperatingCashFlowTTM,capexToRevenueTTM,capexToDepreciationTTM,stockBasedCompensationToRevenueTTM,grahamNumberTTM,roicTTM,returnOnTangibleAssetsTTM,grahamNetNetTTM,workingCapitalTTM,tangibleAssetValueTTM,netCurrentAssetValueTTM,investedCapitalTTM,averageReceivablesTTM,averagePayablesTTM,averageInventoryTTM,daysSalesOutstandingTTM,daysPayablesOutstandingTTM,daysOfInventoryOnHandTTM,receivablesTurnoverTTM,payablesTurnoverTTM,inventoryTurnoverTTM,roeTTM,capexPerShareTTM,dividendPerShareTTM,debtToMarketCapTTM FROM KeyMetrics_CompanyTTM WHERE symbol = " + symbol + "LIMIT ?"
	rows, error_show := db.Query(q, limit)

	oListKMCTTM := []models.KeyMetrics_CompanyTTM{}

	if error_show != nil {

		return oListKMCTTM, error_show
	}

	for rows.Next() {
		oKMCTTM := models.KeyMetrics_CompanyTTM{}
		rows.Scan(&oKMCTTM.Symbol,
			&oKMCTTM.RevenuePerShareTTM,
			&oKMCTTM.NetIncomePerShareTTM,
			&oKMCTTM.OperatingCashFlowPerShareTTM,
			&oKMCTTM.FreeCashFlowPerShareTTM,
			&oKMCTTM.CashPerShareTTM,
			&oKMCTTM.BookValuePerShareTTM,
			&oKMCTTM.TangibleBookValuePerShareTTM,
			&oKMCTTM.ShareholdersEquityPerShareTTM,
			&oKMCTTM.InterestDebtPerShareTTM,
			&oKMCTTM.MarketCapTTM,
			&oKMCTTM.EnterpriseValueTTM,
			&oKMCTTM.PeRatioTTM,
			&oKMCTTM.PriceToSalesRatioTTM,
			&oKMCTTM.PocfratioTTM,
			&oKMCTTM.PfcfRatioTTM,
			&oKMCTTM.PbRatioTTM,
			&oKMCTTM.PtbRatioTTM,
			&oKMCTTM.EvToSalesTTM,
			&oKMCTTM.EnterpriseValueOverEBITDATTM,
			&oKMCTTM.EvToOperatingCashFlowTTM,
			&oKMCTTM.EvToFreeCashFlowTTM,
			&oKMCTTM.EarningsYieldTTM,
			&oKMCTTM.FreeCashFlowYieldTTM,
			&oKMCTTM.DebtToEquityTTM,
			&oKMCTTM.DebtToAssetsTTM,
			&oKMCTTM.NetDebtToEBITDATTM,
			&oKMCTTM.CurrentRatioTTM,
			&oKMCTTM.InterestCoverageTTM,
			&oKMCTTM.IncomeQualityTTM,
			&oKMCTTM.DividendYieldTTM,
			&oKMCTTM.DividendYieldPercentageTTM,
			&oKMCTTM.PayoutRatioTTM,
			&oKMCTTM.SalesGeneralAndAdministrativeToRevenueTTM,
			&oKMCTTM.ResearchAndDevelopementToRevenueTTM,
			&oKMCTTM.IntangiblesToTotalAssetsTTM,
			&oKMCTTM.CapexToOperatingCashFlowTTM,
			&oKMCTTM.CapexToRevenueTTM,
			&oKMCTTM.CapexToDepreciationTTM,
			&oKMCTTM.StockBasedCompensationToRevenueTTM,
			&oKMCTTM.GrahamNumberTTM,
			&oKMCTTM.RoicTTM,
			&oKMCTTM.ReturnOnTangibleAssetsTTM,
			&oKMCTTM.GrahamNetNetTTM,
			&oKMCTTM.WorkingCapitalTTM,
			&oKMCTTM.TangibleAssetValueTTM,
			&oKMCTTM.NetCurrentAssetValueTTM,
			&oKMCTTM.InvestedCapitalTTM,
			&oKMCTTM.AverageReceivablesTTM,
			&oKMCTTM.AveragePayablesTTM,
			&oKMCTTM.AverageInventoryTTM,
			&oKMCTTM.DaysSalesOutstandingTTM,
			&oKMCTTM.DaysPayablesOutstandingTTM,
			&oKMCTTM.DaysOfInventoryOnHandTTM,
			&oKMCTTM.ReceivablesTurnoverTTM,
			&oKMCTTM.PayablesTurnoverTTM,
			&oKMCTTM.InventoryTurnoverTTM,
			&oKMCTTM.RoeTTM,
			&oKMCTTM.CapexPerShareTTM,
			&oKMCTTM.DividendPerShareTTM,
			&oKMCTTM.DebtToMarketCapTTM)
		oListKMCTTM = append(oListKMCTTM, oKMCTTM)
	}

	//OK
	return oListKMCTTM, nil
}
