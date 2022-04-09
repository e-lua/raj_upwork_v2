package all

import (
	"strconv"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_Annual(symbol string, limit int) ([]models.IncomeStatement_Annual, error) {

	db := models.SingleStoreCN
	q := `SELECT date,symbol,reportedCurrency,cik,fillingDate,acceptedDate,calendarYear,period,revenue,costOfRevenue,grossProfit,grossProfitRatio,researchAndDevelopmentExpenses,generalAndAdministrativeExpenses,sellingAndMarketingExpenses,sellingGeneralAndAdministrativeExpenses,otherExpenses,operatingExpenses,costAndExpenses,interestIncome,interestExpense,depreciationAndAmortization,ebitda,ebitdaratio,operatingIncome,operatingIncomeRatio,totalOtherIncomeExpensesNet,incomeBeforeTax,incomeBeforeTaxRatio,incomeTaxExpense,netIncome,netIncomeRatio,eps,epsdiluted,weightedAverageShsOut,weightedAverageShsOutDil,link,finalLink FROM IncomeStatement_Annual WHERE symbol=? LIMIT ` + strconv.Itoa(limit)
	rows, error_show := db.Query(q, symbol)

	oListISA := []models.IncomeStatement_Annual{}

	if error_show != nil {

		return oListISA, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oISA := models.IncomeStatement_Annual{}
		rows.Scan(&oISA.Date,
			&oISA.Symbol,
			&oISA.ReportedCurrency,
			&oISA.Cik,
			&oISA.FillingDate,
			&oISA.AcceptedDate,
			&oISA.CalendarYear,
			&oISA.Period,
			&oISA.Revenue,
			&oISA.CostOfRevenue,
			&oISA.GrossProfit,
			&oISA.GrossProfitRatio,
			&oISA.ResearchAndDevelopmentExpenses,
			&oISA.GeneralAndAdministrativeExpenses,
			&oISA.SellingAndMarketingExpenses,
			&oISA.SellingGeneralAndAdministrativeExpenses,
			&oISA.OtherExpenses,
			&oISA.OperatingExpenses,
			&oISA.CostAndExpenses,
			&oISA.InterestIncome,
			&oISA.InterestExpense,
			&oISA.DepreciationAndAmortization,
			&oISA.Ebitda,
			&oISA.Ebitdaratio,
			&oISA.OperatingIncome,
			&oISA.OperatingIncomeRatio,
			&oISA.TotalOtherIncomeExpensesNet,
			&oISA.IncomeBeforeTax,
			&oISA.IncomeBeforeTaxRatio,
			&oISA.IncomeTaxExpense,
			&oISA.NetIncome,
			&oISA.NetIncomeRatio,
			&oISA.Eps,
			&oISA.Epsdiluted,
			&oISA.WeightedAverageShsOut,
			&oISA.WeightedAverageShsOutDil,
			&oISA.Link,
			&oISA.FinalLink)
		oListISA = append(oListISA, oISA)
	}

	//OK
	return oListISA, nil
}

func Si_Find_Quarter(symbol string, limit int) ([]models.IncomeStatement_Quarter, error) {

	db := models.SingleStoreCN
	q := `SELECT date,symbol,reportedCurrency,cik,fillingDate,acceptedDate,calendarYear,period,revenue,costOfRevenue,grossProfit,grossProfitRatio,researchAndDevelopmentExpenses,generalAndAdministrativeExpenses,sellingAndMarketingExpenses,sellingGeneralAndAdministrativeExpenses,otherExpenses,operatingExpenses,costAndExpenses,interestIncome,interestExpense,depreciationAndAmortization,ebitda,ebitdaratio,operatingIncome,operatingIncomeRatio,totalOtherIncomeExpensesNet,incomeBeforeTax,incomeBeforeTaxRatio,incomeTaxExpense,netIncome,netIncomeRatio,eps,epsdiluted,weightedAverageShsOut,weightedAverageShsOutDil,link,finalLink FROM IncomeStatement_Quarter WHERE symbol = ? LIMIT ` + strconv.Itoa(limit)
	rows, error_show := db.Query(q, symbol)

	oListISQ := []models.IncomeStatement_Quarter{}

	if error_show != nil {

		return oListISQ, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oISQ := models.IncomeStatement_Quarter{}
		rows.Scan(&oISQ.Date,
			&oISQ.Symbol,
			&oISQ.ReportedCurrency,
			&oISQ.Cik,
			&oISQ.FillingDate,
			&oISQ.AcceptedDate,
			&oISQ.CalendarYear,
			&oISQ.Period,
			&oISQ.Revenue,
			&oISQ.CostOfRevenue,
			&oISQ.GrossProfit,
			&oISQ.GrossProfitRatio,
			&oISQ.ResearchAndDevelopmentExpenses,
			&oISQ.GeneralAndAdministrativeExpenses,
			&oISQ.SellingAndMarketingExpenses,
			&oISQ.SellingGeneralAndAdministrativeExpenses,
			&oISQ.OtherExpenses,
			&oISQ.OperatingExpenses,
			&oISQ.CostAndExpenses,
			&oISQ.InterestIncome,
			&oISQ.InterestExpense,
			&oISQ.DepreciationAndAmortization,
			&oISQ.Ebitda,
			&oISQ.Ebitdaratio,
			&oISQ.OperatingIncome,
			&oISQ.OperatingIncomeRatio,
			&oISQ.TotalOtherIncomeExpensesNet,
			&oISQ.IncomeBeforeTax,
			&oISQ.IncomeBeforeTaxRatio,
			&oISQ.IncomeTaxExpense,
			&oISQ.NetIncome,
			&oISQ.NetIncomeRatio,
			&oISQ.Eps,
			&oISQ.Epsdiluted,
			&oISQ.WeightedAverageShsOut,
			&oISQ.WeightedAverageShsOutDil,
			&oISQ.Link,
			&oISQ.FinalLink)
		oListISQ = append(oListISQ, oISQ)
	}

	//OK
	return oListISQ, nil
}

func Si_Find_AnnualGrowth(symbol string, limit int) ([]models.IncomeStatement_AnnualGrowth, error) {

	db := models.SingleStoreCN
	q := `SELECT date,symbol,period,growthRevenue,growthCostOfRevenue,growthGrossProfit,growthGrossProfitRatio,growthResearchAndDevelopmentExpenses,growthGeneralAndAdministrativeExpenses,growthSellingAndMarketingExpenses,growthOtherExpenses,growthOperatingExpenses,growthCostAndExpenses,growthInterestExpense,growthDepreciationAndAmortization,growthEBITDA,growthEBITDARatio,growthOperatingIncome,growthOperatingIncomeRatio,growthTotalOtherIncomeExpensesNet,growthIncomeBeforeTax,growthIncomeBeforeTaxRatio,growthIncomeTaxExpense,growthNetIncome,growthNetIncomeRatio,growthEPS,growthEPSDiluted,growthWeightedAverageShsOut,growthWeightedAverageShsOutDil FROM IncomeStatement_AnnualGrowth WHERE symbol = ? LIMIT ` + strconv.Itoa(limit)
	rows, error_show := db.Query(q, symbol)

	oListISAG := []models.IncomeStatement_AnnualGrowth{}

	if error_show != nil {

		return oListISAG, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oISAG := models.IncomeStatement_AnnualGrowth{}
		rows.Scan(&oISAG.Date,
			&oISAG.Symbol,
			&oISAG.Period,
			&oISAG.GrowthRevenue,
			&oISAG.GrowthCostOfRevenue,
			&oISAG.GrowthGrossProfit,
			&oISAG.GrowthGrossProfitRatio,
			&oISAG.GrowthResearchAndDevelopmentExpenses,
			&oISAG.GrowthGeneralAndAdministrativeExpenses,
			&oISAG.GrowthSellingAndMarketingExpenses,
			&oISAG.GrowthOtherExpenses,
			&oISAG.GrowthOperatingExpenses,
			&oISAG.GrowthCostAndExpenses,
			&oISAG.GrowthInterestExpense,
			&oISAG.GrowthDepreciationAndAmortization,
			&oISAG.GrowthEBITDA,
			&oISAG.GrowthEBITDARatio,
			&oISAG.GrowthOperatingIncome,
			&oISAG.GrowthOperatingIncomeRatio,
			&oISAG.GrowthTotalOtherIncomeExpensesNet,
			&oISAG.GrowthIncomeBeforeTax,
			&oISAG.GrowthIncomeBeforeTaxRatio,
			&oISAG.GrowthIncomeTaxExpense,
			&oISAG.GrowthNetIncome,
			&oISAG.GrowthNetIncomeRatio,
			&oISAG.GrowthEPS,
			&oISAG.GrowthEPSDiluted,
			&oISAG.GrowthWeightedAverageShsOut,
			&oISAG.GrowthWeightedAverageShsOutDil)
		oListISAG = append(oListISAG, oISAG)
	}

	//OK
	return oListISAG, nil
}

func Si_Find_QuarterGrowth(symbol string, limit int) ([]models.IncomeStatement_QuarterGrowth, error) {

	db := models.SingleStoreCN
	q := `SELECT date,symbol,period,growthRevenue,growthCostOfRevenue,growthGrossProfit,growthGrossProfitRatio,growthResearchAndDevelopmentExpenses,growthGeneralAndAdministrativeExpenses,growthSellingAndMarketingExpenses,growthOtherExpenses,growthOperatingExpenses,growthCostAndExpenses,growthInterestExpense,growthDepreciationAndAmortization,growthEBITDA,growthEBITDARatio,growthOperatingIncome,growthOperatingIncomeRatio,growthTotalOtherIncomeExpensesNet,growthIncomeBeforeTax,growthIncomeBeforeTaxRatio,growthIncomeTaxExpense,growthNetIncome,growthNetIncomeRatio,growthEPS,growthEPSDiluted,growthWeightedAverageShsOut,growthWeightedAverageShsOutDil FROM IncomeStatement_QuarterGrowth WHERE symbol = ? LIMIT ` + strconv.Itoa(limit)
	rows, error_show := db.Query(q, symbol)

	oListISQG := []models.IncomeStatement_QuarterGrowth{}

	if error_show != nil {

		return oListISQG, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oISQG := models.IncomeStatement_QuarterGrowth{}
		rows.Scan(&oISQG.Date,
			&oISQG.Symbol,
			&oISQG.Period,
			&oISQG.GrowthRevenue,
			&oISQG.GrowthCostOfRevenue,
			&oISQG.GrowthGrossProfit,
			&oISQG.GrowthGrossProfitRatio,
			&oISQG.GrowthResearchAndDevelopmentExpenses,
			&oISQG.GrowthGeneralAndAdministrativeExpenses,
			&oISQG.GrowthSellingAndMarketingExpenses,
			&oISQG.GrowthOtherExpenses,
			&oISQG.GrowthOperatingExpenses,
			&oISQG.GrowthCostAndExpenses,
			&oISQG.GrowthInterestExpense,
			&oISQG.GrowthDepreciationAndAmortization,
			&oISQG.GrowthEBITDA,
			&oISQG.GrowthEBITDARatio,
			&oISQG.GrowthOperatingIncome,
			&oISQG.GrowthOperatingIncomeRatio,
			&oISQG.GrowthTotalOtherIncomeExpensesNet,
			&oISQG.GrowthIncomeBeforeTax,
			&oISQG.GrowthIncomeBeforeTaxRatio,
			&oISQG.GrowthIncomeTaxExpense,
			&oISQG.GrowthNetIncome,
			&oISQG.GrowthNetIncomeRatio,
			&oISQG.GrowthEPS,
			&oISQG.GrowthEPSDiluted,
			&oISQG.GrowthWeightedAverageShsOut,
			&oISQG.GrowthWeightedAverageShsOutDil)
		oListISQG = append(oListISQG, oISQG)
	}

	//OK
	return oListISQG, nil
}
