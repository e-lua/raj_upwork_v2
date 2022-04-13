package all

import (
	"log"
	"time"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Add(symbol string, cp []models.CompanyProfile, isa []models.IncomeStatement_Annual, isq []models.IncomeStatement_Quarter, isag []models.IncomeStatement_AnnualGrowth, isqg []models.IncomeStatement_QuarterGrowth, bsa []models.BalanceSheet_Annual, bsq []models.BalanceSheet_Quarter, bsag []models.BalanceSheet_AnnualGrowth, bsqg []models.BalanceSheet_QuarterGrowth, cfa []models.CashFlow_Annual, cfq []models.CashFlow_Quarter, cfag []models.CashFlow_AnnualGrowth, cfqg []models.CashFlow_QuarterGrowth, fra []models.FinancialRatio_Annual, frq []models.FinancialRatio_Quarter, frattm []models.FinancialRatio_AnnualTTM, kmcttm []models.KeyMetrics_CompanyTTM, kma []models.KeyMetrics_Annual, kmq []models.KeyMetrics_Quarter) error {

	counter_BalanceSheet_Annual := 0
	counter_BalanceSheet_AnnualGrowth := 0
	counter_BalanceSheet_Quarter := 0
	counter_BalanceSheet_QuarterGrowth := 0
	counter_CashFlow_Annual := 0
	counter_CashFlow_AnnualGrowth := 0
	counter_CashFlow_Quarter := 0
	counter_CashFlow_QuarterGrowth := 0
	counter_FinancialRatio_Annual := 0
	counter_FinancialRatio_AnnualTTM := 0
	counter_FinancialRatio_Quarter := 0
	counter_IncomeStatement_Annual := 0
	counter_IncomeStatement_AnnualGrowth := 0
	counter_IncomeStatement_Quarter := 0
	counter_IncomeStatement_QuarterGrowth := 0
	counter_CompanyProfile := 0
	counter_KeyMetrics_Annual := 0
	counter_KeyMetrics_Quarter := 0
	counter_KeyMetrics_CompanyTTM := 0

	/*-------------------DATA: IncomeStatement Annual---------------*/
	vals_ISA := []interface{}{}
	sqlStr_ISA := `INSERT INTO IncomeStatement_Annual(
		id,
		date,
		symbol,
		reportedCurrency,
		cik,
		fillingDate,
		acceptedDate,
		calendarYear,
		period,
		revenue,
		costOfRevenue,
		grossProfit,
		grossProfitRatio,
		researchAndDevelopmentExpenses,
		generalAndAdministrativeExpenses,
		sellingAndMarketingExpenses,
		sellingGeneralAndAdministrativeExpenses,
		otherExpenses,
		operatingExpenses,
		costAndExpenses,
		interestIncome,
		interestExpense,
		depreciationAndAmortization,
		ebitda,
		ebitdaratio,
		operatingIncome,
		operatingIncomeRatio,
		totalOtherIncomeExpensesNet,
		incomeBeforeTax,
		incomeBeforeTaxRatio,
		incomeTaxExpense,
		netIncome,
		netIncomeRatio,
		eps,
		epsdiluted,
		weightedAverageShsOut,
		weightedAverageShsOutDil,
		link,
		finalLink) VALUES`
	counter_ISA := 0
	for _, val := range isa {
		counter_IncomeStatement_Annual = counter_IncomeStatement_Annual + 1
		//Insert data in the query
		sqlStr_ISA += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_ISA = append(vals_ISA, time.Now().UnixMilli()+int64(counter_ISA), val.Date,
			symbol,
			val.ReportedCurrency,
			val.Cik,
			val.FillingDate,
			val.AcceptedDate,
			val.CalendarYear,
			val.Period,
			val.Revenue,
			val.CostOfRevenue,
			val.GrossProfit,
			val.GrossProfitRatio,
			val.ResearchAndDevelopmentExpenses,
			val.GeneralAndAdministrativeExpenses,
			val.SellingAndMarketingExpenses,
			val.SellingGeneralAndAdministrativeExpenses,
			val.OtherExpenses,
			val.OperatingExpenses,
			val.CostAndExpenses,
			val.InterestIncome,
			val.InterestExpense,
			val.DepreciationAndAmortization,
			val.Ebitda,
			val.Ebitdaratio,
			val.OperatingIncome,
			val.OperatingIncomeRatio,
			val.TotalOtherIncomeExpensesNet,
			val.IncomeBeforeTax,
			val.IncomeBeforeTaxRatio,
			val.IncomeTaxExpense,
			val.NetIncome,
			val.NetIncomeRatio,
			val.Eps,
			val.Epsdiluted,
			val.WeightedAverageShsOut,
			val.WeightedAverageShsOutDil,
			val.Link,
			val.FinalLink)
		//Sum counter
		counter_ISA = counter_ISA + 1
	}
	//Deleting the last nil value
	sqlStr_ISA = sqlStr_ISA[0 : len(sqlStr_ISA)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: IncomeStatement Quarter---------------*/
	vals_ISQ := []interface{}{}
	sqlStr_ISQ := `INSERT INTO IncomeStatement_Quarter(
		id,
		date,
		symbol,
		reportedCurrency,
		cik,
		fillingDate,
		acceptedDate,
		calendarYear,
		period,
		revenue,
		costOfRevenue,
		grossProfit,
		grossProfitRatio,
		researchAndDevelopmentExpenses,
		generalAndAdministrativeExpenses,
		sellingAndMarketingExpenses,
		sellingGeneralAndAdministrativeExpenses,
		otherExpenses,
		operatingExpenses,
		costAndExpenses,
		interestIncome,
		interestExpense,
		depreciationAndAmortization,
		ebitda,
		ebitdaratio,
		operatingIncome,
		operatingIncomeRatio,
		totalOtherIncomeExpensesNet,
		incomeBeforeTax,
		incomeBeforeTaxRatio,
		incomeTaxExpense,
		netIncome,
		netIncomeRatio,
		eps,
		epsdiluted,
		weightedAverageShsOut,
		weightedAverageShsOutDil,
		link,
		finalLink) VALUES`
	counter_ISQ := 0
	for _, val := range isq {
		counter_IncomeStatement_Quarter = counter_IncomeStatement_Quarter + 1
		//Insert data in the query
		sqlStr_ISQ += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_ISQ = append(vals_ISQ, time.Now().UnixMilli()+int64(counter_ISQ), val.Date,
			symbol,
			val.ReportedCurrency,
			val.Cik,
			val.FillingDate,
			val.AcceptedDate,
			val.CalendarYear,
			val.Period,
			val.Revenue,
			val.CostOfRevenue,
			val.GrossProfit,
			val.GrossProfitRatio,
			val.ResearchAndDevelopmentExpenses,
			val.GeneralAndAdministrativeExpenses,
			val.SellingAndMarketingExpenses,
			val.SellingGeneralAndAdministrativeExpenses,
			val.OtherExpenses,
			val.OperatingExpenses,
			val.CostAndExpenses,
			val.InterestIncome,
			val.InterestExpense,
			val.DepreciationAndAmortization,
			val.Ebitda,
			val.Ebitdaratio,
			val.OperatingIncome,
			val.OperatingIncomeRatio,
			val.TotalOtherIncomeExpensesNet,
			val.IncomeBeforeTax,
			val.IncomeBeforeTaxRatio,
			val.IncomeTaxExpense,
			val.NetIncome,
			val.NetIncomeRatio,
			val.Eps,
			val.Epsdiluted,
			val.WeightedAverageShsOut,
			val.WeightedAverageShsOutDil,
			val.Link,
			val.FinalLink)
		//Sum counter
		counter_ISQ = counter_ISQ + 1
	}
	//Deleting the last nil value
	sqlStr_ISQ = sqlStr_ISQ[0 : len(sqlStr_ISQ)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: IncomeStatement Annual Growth---------------*/
	vals_ISAG := []interface{}{}
	sqlStr_ISAG := `INSERT INTO IncomeStatement_AnnualGrowth (
		id,
		date,
		symbol,
		period,
		growthRevenue,
		growthCostOfRevenue,
		growthGrossProfit,
		growthGrossProfitRatio,
		growthResearchAndDevelopmentExpenses,
		growthGeneralAndAdministrativeExpenses,
		growthSellingAndMarketingExpenses,
		growthOtherExpenses,
		growthOperatingExpenses,
		growthCostAndExpenses,
		growthInterestExpense,
		growthDepreciationAndAmortization,
		growthEBITDA,
		growthEBITDARatio,
		growthOperatingIncome,
		growthOperatingIncomeRatio,
		growthTotalOtherIncomeExpensesNet,
		growthIncomeBeforeTax,
		growthIncomeBeforeTaxRatio,
		growthIncomeTaxExpense,
		growthNetIncome,
		growthNetIncomeRatio,
		growthEPS,
		growthEPSDiluted,
		growthWeightedAverageShsOut,
		growthWeightedAverageShsOutDil) VALUES`
	counter_ISAG := 0
	for _, val := range isag {
		counter_IncomeStatement_AnnualGrowth = counter_IncomeStatement_AnnualGrowth + 1
		//Insert data in the query
		sqlStr_ISAG += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_ISAG = append(vals_ISAG, time.Now().UnixMilli()+int64(counter_ISAG),
			val.Date,
			symbol,
			val.Period,
			val.GrowthRevenue,
			val.GrowthCostOfRevenue,
			val.GrowthGrossProfit,
			val.GrowthGrossProfitRatio,
			val.GrowthResearchAndDevelopmentExpenses,
			val.GrowthGeneralAndAdministrativeExpenses,
			val.GrowthSellingAndMarketingExpenses,
			val.GrowthOtherExpenses,
			val.GrowthOperatingExpenses,
			val.GrowthCostAndExpenses,
			val.GrowthInterestExpense,
			val.GrowthDepreciationAndAmortization,
			val.GrowthEBITDA,
			val.GrowthEBITDARatio,
			val.GrowthOperatingIncome,
			val.GrowthOperatingIncomeRatio,
			val.GrowthTotalOtherIncomeExpensesNet,
			val.GrowthIncomeBeforeTax,
			val.GrowthIncomeBeforeTaxRatio,
			val.GrowthIncomeTaxExpense,
			val.GrowthNetIncome,
			val.GrowthNetIncomeRatio,
			val.GrowthEPS,
			val.GrowthEPSDiluted,
			val.GrowthWeightedAverageShsOut,
			val.GrowthWeightedAverageShsOutDil)
		//Sum counter
		counter_ISAG = counter_ISAG + 1
	}
	//Deleting the last nil value
	sqlStr_ISAG = sqlStr_ISAG[0 : len(sqlStr_ISAG)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: IncomeStatement Quarter Growth---------------*/
	vals_ISQG := []interface{}{}
	sqlStr_ISQG := `INSERT INTO IncomeStatement_QuarterGrowth (
		id,
		date,
		symbol,
		period,
		growthRevenue,
		growthCostOfRevenue,
		growthGrossProfit,
		growthGrossProfitRatio,
		growthResearchAndDevelopmentExpenses,
		growthGeneralAndAdministrativeExpenses,
		growthSellingAndMarketingExpenses,
		growthOtherExpenses,
		growthOperatingExpenses,
		growthCostAndExpenses,
		growthInterestExpense,
		growthDepreciationAndAmortization,
		growthEBITDA,
		growthEBITDARatio,
		growthOperatingIncome,
		growthOperatingIncomeRatio,
		growthTotalOtherIncomeExpensesNet,
		growthIncomeBeforeTax,
		growthIncomeBeforeTaxRatio,
		growthIncomeTaxExpense,
		growthNetIncome,
		growthNetIncomeRatio,
		growthEPS,
		growthEPSDiluted,
		growthWeightedAverageShsOut,
		growthWeightedAverageShsOutDil) VALUES`
	counter_ISQG := 0
	for _, val := range isag {
		counter_IncomeStatement_QuarterGrowth = counter_IncomeStatement_QuarterGrowth + 1
		//Insert data in the query
		sqlStr_ISQG += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_ISQG = append(vals_ISQG, time.Now().UnixMilli()+int64(counter_ISQG),
			val.Date,
			symbol,
			val.Period,
			val.GrowthRevenue,
			val.GrowthCostOfRevenue,
			val.GrowthGrossProfit,
			val.GrowthGrossProfitRatio,
			val.GrowthResearchAndDevelopmentExpenses,
			val.GrowthGeneralAndAdministrativeExpenses,
			val.GrowthSellingAndMarketingExpenses,
			val.GrowthOtherExpenses,
			val.GrowthOperatingExpenses,
			val.GrowthCostAndExpenses,
			val.GrowthInterestExpense,
			val.GrowthDepreciationAndAmortization,
			val.GrowthEBITDA,
			val.GrowthEBITDARatio,
			val.GrowthOperatingIncome,
			val.GrowthOperatingIncomeRatio,
			val.GrowthTotalOtherIncomeExpensesNet,
			val.GrowthIncomeBeforeTax,
			val.GrowthIncomeBeforeTaxRatio,
			val.GrowthIncomeTaxExpense,
			val.GrowthNetIncome,
			val.GrowthNetIncomeRatio,
			val.GrowthEPS,
			val.GrowthEPSDiluted,
			val.GrowthWeightedAverageShsOut,
			val.GrowthWeightedAverageShsOutDil)
		//Sum counter
		counter_ISQG = counter_ISQG + 1
	}
	//Deleting the last nil value
	sqlStr_ISQG = sqlStr_ISQG[0 : len(sqlStr_ISQG)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: BalanceSheet Annual---------------*/
	vals_BSA := []interface{}{}
	sqlStr_BSA := `INSERT INTO BalanceSheet_Annual(
		id,
		date,
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
		finalLink
	) VALUES`
	counter_BSA := 0
	for _, val := range bsa {
		counter_BalanceSheet_Annual += counter_BalanceSheet_Annual + 1
		//Insert data in the query
		sqlStr_BSA += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_BSA = append(vals_BSA, time.Now().UnixMilli()+int64(counter_BSA),
			val.Date,
			symbol,
			val.ReportedCurrency,
			val.Cik,
			val.FillingDate,
			val.AcceptedDate,
			val.CalendarYear,
			val.Period,
			val.CashAndCashEquivalents,
			val.ShortTermInvestments,
			val.CashAndShortTermInvestments,
			val.NetReceivables,
			val.Inventory,
			val.OtherCurrentAssets,
			val.TotalCurrentAssets,
			val.PropertyPlantEquipmentNet,
			val.Goodwill,
			val.IntangibleAssets,
			val.GoodwillAndIntangibleAssets,
			val.LongTermInvestments,
			val.TaxAssets,
			val.OtherNonCurrentAssets,
			val.TotalNonCurrentAssets,
			val.OtherAssets,
			val.TotalAssets,
			val.AccountPayables,
			val.ShortTermDebt,
			val.TaxPayables,
			val.DeferredRevenue,
			val.OtherCurrentLiabilities,
			val.TotalCurrentLiabilities,
			val.LongTermDebt,
			val.DeferredRevenueNonCurrent,
			val.DeferredTaxLiabilitiesNonCurrent,
			val.OtherNonCurrentLiabilities,
			val.TotalNonCurrentLiabilities,
			val.OtherLiabilities,
			val.CapitalLeaseObligations,
			val.TotalLiabilities,
			val.PreferredStock,
			val.CommonStock,
			val.RetainedEarnings,
			val.AccumulatedOtherComprehensiveIncomeLoss,
			val.OthertotalStockholdersEquity,
			val.TotalStockholdersEquity,
			val.TotalLiabilitiesAndStockholdersEquity,
			val.MinorityInterest,
			val.TotalEquity,
			val.TotalLiabilitiesAndTotalEquity,
			val.TotalInvestments,
			val.TotalDebt,
			val.NetDebt,
			val.Link,
			val.FinalLink)
		//Sum counter
		counter_BSA = counter_BSA + 1
	}
	//Deleting the last nil value
	sqlStr_BSA = sqlStr_BSA[0 : len(sqlStr_BSA)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: BalanceSheet Quarter---------------*/
	vals_BSQ := []interface{}{}
	sqlStr_BSQ := `INSERT INTO BalanceSheet_Quarter(
		id,
		date,
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
		finalLink
	) VALUES`
	counter_BSQ := 0
	for _, val := range bsq {
		counter_BalanceSheet_Quarter = counter_BalanceSheet_Quarter + 1
		//Insert data in the query
		sqlStr_BSQ += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_BSQ = append(vals_BSQ, time.Now().UnixMilli()+int64(counter_BSQ),
			val.Date,
			symbol,
			val.ReportedCurrency,
			val.Cik,
			val.FillingDate,
			val.AcceptedDate,
			val.CalendarYear,
			val.Period,
			val.CashAndCashEquivalents,
			val.ShortTermInvestments,
			val.CashAndShortTermInvestments,
			val.NetReceivables,
			val.Inventory,
			val.OtherCurrentAssets,
			val.TotalCurrentAssets,
			val.PropertyPlantEquipmentNet,
			val.Goodwill,
			val.IntangibleAssets,
			val.GoodwillAndIntangibleAssets,
			val.LongTermInvestments,
			val.TaxAssets,
			val.OtherNonCurrentAssets,
			val.TotalNonCurrentAssets,
			val.OtherAssets,
			val.TotalAssets,
			val.AccountPayables,
			val.ShortTermDebt,
			val.TaxPayables,
			val.DeferredRevenue,
			val.OtherCurrentLiabilities,
			val.TotalCurrentLiabilities,
			val.LongTermDebt,
			val.DeferredRevenueNonCurrent,
			val.DeferredTaxLiabilitiesNonCurrent,
			val.OtherNonCurrentLiabilities,
			val.TotalNonCurrentLiabilities,
			val.OtherLiabilities,
			val.CapitalLeaseObligations,
			val.TotalLiabilities,
			val.PreferredStock,
			val.CommonStock,
			val.RetainedEarnings,
			val.AccumulatedOtherComprehensiveIncomeLoss,
			val.OthertotalStockholdersEquity,
			val.TotalStockholdersEquity,
			val.TotalLiabilitiesAndStockholdersEquity,
			val.MinorityInterest,
			val.TotalEquity,
			val.TotalLiabilitiesAndTotalEquity,
			val.TotalInvestments,
			val.TotalDebt,
			val.NetDebt,
			val.Link,
			val.FinalLink)
		//Sum counter
		counter_BSQ = counter_BSQ + 1
	}
	//Deleting the last nil value
	sqlStr_BSQ = sqlStr_BSQ[0 : len(sqlStr_BSQ)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: BalanceSheet Annual Growth---------------*/
	vals_BSAG := []interface{}{}
	sqlStr_BSAG := `INSERT INTO BalanceSheet_AnnualGrowth(
		id,
    date,
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
    growthNetDebt
	) VALUES`
	counter_BSAG := 0
	for _, val := range bsag {
		counter_BalanceSheet_AnnualGrowth = counter_BalanceSheet_AnnualGrowth + 1
		//Insert data in the query
		sqlStr_BSAG += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_BSAG = append(vals_BSAG, time.Now().UnixMilli()+int64(counter_BSAG),
			val.Date,
			symbol,
			val.Period,
			val.GrowthCashAndCashEquivalents,
			val.GrowthShortTermInvestments,
			val.GrowthCashAndShortTermInvestments,
			val.GrowthNetReceivables,
			val.GrowthInventory,
			val.GrowthOtherCurrentAssets,
			val.GrowthTotalCurrentAssets,
			val.GrowthPropertyPlantEquipmentNet,
			val.GrowthGoodwill,
			val.GrowthIntangibleAssets,
			val.GrowthGoodwillAndIntangibleAssets,
			val.GrowthLongTermInvestments,
			val.GrowthTaxAssets,
			val.GrowthOtherNonCurrentAssets,
			val.GrowthTotalNonCurrentAssets,
			val.GrowthOtherAssets,
			val.GrowthTotalAssets,
			val.GrowthAccountPayables,
			val.GrowthShortTermDebt,
			val.GrowthTaxPayables,
			val.GrowthDeferredRevenue,
			val.GrowthOtherCurrentLiabilities,
			val.GrowthTotalCurrentLiabilities,
			val.GrowthLongTermDebt,
			val.GrowthDeferredRevenueNonCurrent,
			val.GrowthDeferrredTaxLiabilitiesNonCurrent,
			val.GrowthOtherNonCurrentLiabilities,
			val.GrowthTotalNonCurrentLiabilities,
			val.GrowthOtherLiabilities,
			val.GrowthTotalLiabilities,
			val.GrowthCommonStock,
			val.GrowthRetainedEarnings,
			val.GrowthAccumulatedOtherComprehensiveIncomeLoss,
			val.GrowthOthertotalStockholdersEquity,
			val.GrowthTotalStockholdersEquity,
			val.GrowthTotalLiabilitiesAndStockholdersEquity,
			val.GrowthTotalInvestments,
			val.GrowthTotalDebt,
			val.GrowthNetDebt)
		//Sum counter
		counter_BSAG = counter_BSAG + 1
	}
	//Deleting the last nil value
	sqlStr_BSAG = sqlStr_BSAG[0 : len(sqlStr_BSAG)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: BalanceSheet Quarter Growth---------------*/
	vals_BSQG := []interface{}{}
	sqlStr_BSQG := `INSERT INTO BalanceSheet_QuarterGrowth(
		id,
    date,
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
    growthNetDebt
	) VALUES`
	counter_BSQG := 0
	for _, val := range bsqg {
		counter_BalanceSheet_QuarterGrowth = counter_BalanceSheet_QuarterGrowth + 1
		//Insert data in the query
		sqlStr_BSQG += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_BSQG = append(vals_BSQG, time.Now().UnixMilli()+int64(counter_BSQG),
			val.Date,
			symbol,
			val.Period,
			val.GrowthCashAndCashEquivalents,
			val.GrowthShortTermInvestments,
			val.GrowthCashAndShortTermInvestments,
			val.GrowthNetReceivables,
			val.GrowthInventory,
			val.GrowthOtherCurrentAssets,
			val.GrowthTotalCurrentAssets,
			val.GrowthPropertyPlantEquipmentNet,
			val.GrowthGoodwill,
			val.GrowthIntangibleAssets,
			val.GrowthGoodwillAndIntangibleAssets,
			val.GrowthLongTermInvestments,
			val.GrowthTaxAssets,
			val.GrowthOtherNonCurrentAssets,
			val.GrowthTotalNonCurrentAssets,
			val.GrowthOtherAssets,
			val.GrowthTotalAssets,
			val.GrowthAccountPayables,
			val.GrowthShortTermDebt,
			val.GrowthTaxPayables,
			val.GrowthDeferredRevenue,
			val.GrowthOtherCurrentLiabilities,
			val.GrowthTotalCurrentLiabilities,
			val.GrowthLongTermDebt,
			val.GrowthDeferredRevenueNonCurrent,
			val.GrowthDeferrredTaxLiabilitiesNonCurrent,
			val.GrowthOtherNonCurrentLiabilities,
			val.GrowthTotalNonCurrentLiabilities,
			val.GrowthOtherLiabilities,
			val.GrowthTotalLiabilities,
			val.GrowthCommonStock,
			val.GrowthRetainedEarnings,
			val.GrowthAccumulatedOtherComprehensiveIncomeLoss,
			val.GrowthOthertotalStockholdersEquity,
			val.GrowthTotalStockholdersEquity,
			val.GrowthTotalLiabilitiesAndStockholdersEquity,
			val.GrowthTotalInvestments,
			val.GrowthTotalDebt,
			val.GrowthNetDebt)
		//Sum counter
		counter_BSQG = counter_BSQG + 1
	}
	//Deleting the last nil value
	sqlStr_BSQG = sqlStr_BSQG[0 : len(sqlStr_BSQG)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: CashFlow Annual---------------*/
	vals_CFA := []interface{}{}
	sqlStr_CFA := `INSERT INTO CashFlow_Annual(
		id,
		date,
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
		finalLink
	) VALUES`
	counter_CFA := 0
	for _, val := range cfa {
		counter_CashFlow_Annual = counter_CashFlow_Annual + 1
		//Insert data in the query
		sqlStr_CFA += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_CFA = append(vals_CFA, time.Now().UnixMilli()+int64(counter_CFA),
			val.Date,
			symbol,
			val.ReportedCurrency,
			val.Cik,
			val.FillingDate,
			val.AcceptedDate,
			val.CalendarYear,
			val.Period,
			val.NetIncome,
			val.DepreciationAndAmortization,
			val.DeferredIncomeTax,
			val.StockBasedCompensation,
			val.ChangeInWorkingCapital,
			val.AccountsReceivables,
			val.Inventory,
			val.AccountsPayables,
			val.OtherWorkingCapital,
			val.OtherNonCashItems,
			val.NetCashProvidedByOperatingActivities,
			val.InvestmentsInPropertyPlantAndEquipment,
			val.AcquisitionsNet,
			val.PurchasesOfInvestments,
			val.SalesMaturitiesOfInvestments,
			val.OtherInvestingActivites,
			val.NetCashUsedForInvestingActivites,
			val.DebtRepayment,
			val.CommonStockIssued,
			val.CommonStockRepurchased,
			val.DividendsPaid,
			val.OtherFinancingActivites,
			val.NetCashUsedProvidedByFinancingActivities,
			val.EffectOfForexChangesOnCash,
			val.NetChangeInCash,
			val.CashAtEndOfPeriod,
			val.CashAtBeginningOfPeriod,
			val.OperatingCashFlow,
			val.CapitalExpenditure,
			val.FreeCashFlow,
			val.Link,
			val.FinalLink)
		//Sum counter
		counter_CFA = counter_CFA + 1
	}
	//Deleting the last nil value
	sqlStr_CFA = sqlStr_CFA[0 : len(sqlStr_CFA)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: CashFlow Quarter---------------*/
	vals_CFQ := []interface{}{}
	sqlStr_CFQ := `INSERT INTO CashFlow_Quarter(
		id,
		date,
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
		finalLink
	) VALUES`
	counter_CFQ := 0
	for _, val := range cfq {
		counter_CashFlow_Quarter = counter_CashFlow_Quarter + 1
		//Insert data in the query
		sqlStr_CFQ += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_CFQ = append(vals_CFQ, time.Now().UnixMilli()+int64(counter_CFQ),
			val.Date,
			symbol,
			val.ReportedCurrency,
			val.Cik,
			val.FillingDate,
			val.AcceptedDate,
			val.CalendarYear,
			val.Period,
			val.NetIncome,
			val.DepreciationAndAmortization,
			val.DeferredIncomeTax,
			val.StockBasedCompensation,
			val.ChangeInWorkingCapital,
			val.AccountsReceivables,
			val.Inventory,
			val.AccountsPayables,
			val.OtherWorkingCapital,
			val.OtherNonCashItems,
			val.NetCashProvidedByOperatingActivities,
			val.InvestmentsInPropertyPlantAndEquipment,
			val.AcquisitionsNet,
			val.PurchasesOfInvestments,
			val.SalesMaturitiesOfInvestments,
			val.OtherInvestingActivites,
			val.NetCashUsedForInvestingActivites,
			val.DebtRepayment,
			val.CommonStockIssued,
			val.CommonStockRepurchased,
			val.DividendsPaid,
			val.OtherFinancingActivites,
			val.NetCashUsedProvidedByFinancingActivities,
			val.EffectOfForexChangesOnCash,
			val.NetChangeInCash,
			val.CashAtEndOfPeriod,
			val.CashAtBeginningOfPeriod,
			val.OperatingCashFlow,
			val.CapitalExpenditure,
			val.FreeCashFlow,
			val.Link,
			val.FinalLink)
		//Sum counter
		counter_CFQ = counter_CFQ + 1
	}
	//Deleting the last nil value
	sqlStr_CFQ = sqlStr_CFQ[0 : len(sqlStr_CFQ)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: CashFlow Annual Growth---------------*/
	vals_CFAG := []interface{}{}
	sqlStr_CFAG := `INSERT INTO CashFlow_AnnualGrowth (
		id,
    date,
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
    growthFreeCashFlow
	) VALUES`
	counter_CFAG := 0
	for _, val := range cfag {
		counter_CashFlow_AnnualGrowth = counter_CashFlow_AnnualGrowth + 1
		//Insert data in the query
		sqlStr_CFAG += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_CFAG = append(vals_CFAG, time.Now().UnixMilli()+int64(counter_CFAG),
			val.Date,
			symbol,
			val.Period,
			val.GrowthNetIncome,
			val.GrowthDepreciationAndAmortization,
			val.GrowthDeferredIncomeTax,
			val.GrowthStockBasedCompensation,
			val.GrowthChangeInWorkingCapital,
			val.GrowthAccountsReceivables,
			val.GrowthInventory,
			val.GrowthAccountsPayables,
			val.GrowthOtherWorkingCapital,
			val.GrowthOtherNonCashItems,
			val.GrowthNetCashProvidedByOperatingActivites,
			val.GrowthInvestmentsInPropertyPlantAndEquipment,
			val.GrowthAcquisitionsNet,
			val.GrowthPurchasesOfInvestments,
			val.GrowthSalesMaturitiesOfInvestments,
			val.GrowthOtherInvestingActivites,
			val.GrowthNetCashUsedForInvestingActivites,
			val.GrowthDebtRepayment,
			val.GrowthCommonStockIssued,
			val.GrowthCommonStockRepurchased,
			val.GrowthDividendsPaid,
			val.GrowthOtherFinancingActivites,
			val.GrowthNetCashUsedProvidedByFinancingActivities,
			val.GrowthEffectOfForexChangesOnCash,
			val.GrowthNetChangeInCash,
			val.GrowthCashAtEndOfPeriod,
			val.GrowthCashAtBeginningOfPeriod,
			val.GrowthOperatingCashFlow,
			val.GrowthCapitalExpenditure,
			val.GrowthFreeCashFlow)
		//Sum counter
		counter_CFAG = counter_CFAG + 1
	}
	//Deleting the last nil value
	sqlStr_CFAG = sqlStr_CFAG[0 : len(sqlStr_CFAG)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: CashFlow Quarter Growth---------------*/
	vals_CFQG := []interface{}{}
	sqlStr_CFQG := `INSERT INTO CashFlow_QuarterGrowth (
		id,
    date,
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
    growthFreeCashFlow
	) VALUES`
	counter_CFQG := 0
	for _, val := range cfqg {
		counter_CashFlow_QuarterGrowth = counter_CashFlow_QuarterGrowth + 1
		//Insert data in the query
		sqlStr_CFQG += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_CFQG = append(vals_CFQG, time.Now().UnixMilli()+int64(counter_CFQG),
			val.Date,
			symbol,
			val.Period,
			val.GrowthNetIncome,
			val.GrowthDepreciationAndAmortization,
			val.GrowthDeferredIncomeTax,
			val.GrowthStockBasedCompensation,
			val.GrowthChangeInWorkingCapital,
			val.GrowthAccountsReceivables,
			val.GrowthInventory,
			val.GrowthAccountsPayables,
			val.GrowthOtherWorkingCapital,
			val.GrowthOtherNonCashItems,
			val.GrowthNetCashProvidedByOperatingActivites,
			val.GrowthInvestmentsInPropertyPlantAndEquipment,
			val.GrowthAcquisitionsNet,
			val.GrowthPurchasesOfInvestments,
			val.GrowthSalesMaturitiesOfInvestments,
			val.GrowthOtherInvestingActivites,
			val.GrowthNetCashUsedForInvestingActivites,
			val.GrowthDebtRepayment,
			val.GrowthCommonStockIssued,
			val.GrowthCommonStockRepurchased,
			val.GrowthDividendsPaid,
			val.GrowthOtherFinancingActivites,
			val.GrowthNetCashUsedProvidedByFinancingActivities,
			val.GrowthEffectOfForexChangesOnCash,
			val.GrowthNetChangeInCash,
			val.GrowthCashAtEndOfPeriod,
			val.GrowthCashAtBeginningOfPeriod,
			val.GrowthOperatingCashFlow,
			val.GrowthCapitalExpenditure,
			val.GrowthFreeCashFlow)
		//Sum counter
		counter_CFQG = counter_CFQG + 1
	}
	//Deleting the last nil value
	sqlStr_CFQG = sqlStr_CFQG[0 : len(sqlStr_CFQG)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: FinancialRatio Annual---------------*/
	vals_FRA := []interface{}{}
	sqlStr_FRA := `INSERT INTO FinancialRatio_Annual(
		id,
    date,
    symbol,
    period,
    currentRatio,
    quickRatio,
    cashRatio,
    daysOfSalesOutstanding,
    daysOfInventoryOutstanding,
    operatingCycle,
    daysOfPayablesOutstanding,
    cashConversionCycle,
    grossProfitMargin,
    operatingProfitMargin,
    pretaxProfitMargin,
    netProfitMargin,
    effectiveTaxRate,
    returnOnAssets,
    returnOnEquity,
    returnOnCapitalEmployed,
    netIncomePerEBT,
    ebtPerEbit,
    ebitPerRevenue,
    debtRatio,
    debtEquityRatio,
    longTermDebtToCapitalization,
    totalDebtToCapitalization,
    interestCoverage,
    cashFlowToDebtRatio,
    companyEquityMultiplier,
    receivablesTurnover,
    payablesTurnover,
    inventoryTurnover,
    fixedAssetTurnover,
    assetTurnover,
    operatingCashFlowPerShare,
    freeCashFlowPerShare,
    cashPerShare,
    payoutRatio,
    operatingCashFlowSalesRatio,
    freeCashFlowOperatingCashFlowRatio,
    cashFlowCoverageRatios,
    shortTermCoverageRatios,
    capitalExpenditureCoverageRatio,
    dividendPaidAndCapexCoverageRatio,
    dividendPayoutRatio,
    priceBookValueRatio,
    priceToBookRatio,
    priceEarningsRatio,
    priceToFreeCashFlowsRatio,
    priceToOperatingCashFlowsRatio,
    priceCashFlowRatio,
    priceEarningsToGrowthRatio,
    priceSalesRatio,
    dividendYield,
    enterpriseValueMultiple,
    priceFairValue
	) VALUES`
	counter_FRA := 0
	for _, val := range fra {
		counter_FinancialRatio_Annual = counter_FinancialRatio_Annual + 1
		//Insert data in the query
		sqlStr_FRA += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_FRA = append(vals_FRA, time.Now().UnixMilli()+int64(counter_FRA),
			val.Date,
			symbol,
			val.Period,
			val.CurrentRatio,
			val.QuickRatio,
			val.CashRatio,
			val.DaysOfSalesOutstanding,
			val.DaysOfInventoryOutstanding,
			val.OperatingCycle,
			val.DaysOfPayablesOutstanding,
			val.CashConversionCycle,
			val.GrossProfitMargin,
			val.OperatingProfitMargin,
			val.PretaxProfitMargin,
			val.NetProfitMargin,
			val.EffectiveTaxRate,
			val.ReturnOnAssets,
			val.ReturnOnEquity,
			val.ReturnOnCapitalEmployed,
			val.NetIncomePerEBT,
			val.EbtPerEbit,
			val.EbitPerRevenue,
			val.DebtRatio,
			val.DebtEquityRatio,
			val.LongTermDebtToCapitalization,
			val.TotalDebtToCapitalization,
			val.InterestCoverage,
			val.CashFlowToDebtRatio,
			val.CompanyEquityMultiplier,
			val.ReceivablesTurnover,
			val.PayablesTurnover,
			val.InventoryTurnover,
			val.FixedAssetTurnover,
			val.AssetTurnover,
			val.OperatingCashFlowPerShare,
			val.FreeCashFlowPerShare,
			val.CashPerShare,
			val.PayoutRatio,
			val.OperatingCashFlowSalesRatio,
			val.FreeCashFlowOperatingCashFlowRatio,
			val.CashFlowCoverageRatios,
			val.ShortTermCoverageRatios,
			val.CapitalExpenditureCoverageRatio,
			val.DividendPaidAndCapexCoverageRatio,
			val.DividendPayoutRatio,
			val.PriceBookValueRatio,
			val.PriceToBookRatio,
			val.PriceEarningsRatio,
			val.PriceToFreeCashFlowsRatio,
			val.PriceToOperatingCashFlowsRatio,
			val.PriceCashFlowRatio,
			val.PriceEarningsToGrowthRatio,
			val.PriceSalesRatio,
			val.DividendYield,
			val.EnterpriseValueMultiple,
			val.PriceFairValue)
		//Sum counter
		counter_FRA = counter_FRA + 1
	}
	//Deleting the last nil value
	sqlStr_FRA = sqlStr_FRA[0 : len(sqlStr_FRA)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: FinancialRatio Quarter---------------*/
	vals_FRQ := []interface{}{}
	sqlStr_FRQ := `INSERT INTO FinancialRatio_Quarter (
		id,
    date,
    symbol,
    period,
    currentRatio,
    quickRatio,
    cashRatio,
    daysOfSalesOutstanding,
    daysOfInventoryOutstanding,
    operatingCycle,
    daysOfPayablesOutstanding,
    cashConversionCycle,
    grossProfitMargin,
    operatingProfitMargin,
    pretaxProfitMargin,
    netProfitMargin,
    effectiveTaxRate,
    returnOnAssets,
    returnOnEquity,
    returnOnCapitalEmployed,
    netIncomePerEBT,
    ebtPerEbit,
    ebitPerRevenue,
    debtRatio,
    debtEquityRatio,
    longTermDebtToCapitalization,
    totalDebtToCapitalization,
    interestCoverage,
    cashFlowToDebtRatio,
    companyEquityMultiplier,
    receivablesTurnover,
    payablesTurnover,
    inventoryTurnover,
    fixedAssetTurnover,
    assetTurnover,
    operatingCashFlowPerShare,
    freeCashFlowPerShare,
    cashPerShare,
    payoutRatio,
    operatingCashFlowSalesRatio,
    freeCashFlowOperatingCashFlowRatio,
    cashFlowCoverageRatios,
    shortTermCoverageRatios,
    capitalExpenditureCoverageRatio,
    dividendPaidAndCapexCoverageRatio,
    dividendPayoutRatio,
    priceBookValueRatio,
    priceToBookRatio,
    priceEarningsRatio,
    priceToFreeCashFlowsRatio,
    priceToOperatingCashFlowsRatio,
    priceCashFlowRatio,
    priceEarningsToGrowthRatio,
    priceSalesRatio,
    dividendYield,
    enterpriseValueMultiple,
    priceFairValue
	) VALUES`
	counter_FRQ := 0
	for _, val := range frq {
		counter_FinancialRatio_Quarter = counter_FinancialRatio_Quarter + 1
		//Insert data in the query
		sqlStr_FRQ += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_FRQ = append(vals_FRQ, time.Now().UnixMilli()+int64(counter_FRQ),
			val.Date,
			symbol,
			val.Period,
			val.CurrentRatio,
			val.QuickRatio,
			val.CashRatio,
			val.DaysOfSalesOutstanding,
			val.DaysOfInventoryOutstanding,
			val.OperatingCycle,
			val.DaysOfPayablesOutstanding,
			val.CashConversionCycle,
			val.GrossProfitMargin,
			val.OperatingProfitMargin,
			val.PretaxProfitMargin,
			val.NetProfitMargin,
			val.EffectiveTaxRate,
			val.ReturnOnAssets,
			val.ReturnOnEquity,
			val.ReturnOnCapitalEmployed,
			val.NetIncomePerEBT,
			val.EbtPerEbit,
			val.EbitPerRevenue,
			val.DebtRatio,
			val.DebtEquityRatio,
			val.LongTermDebtToCapitalization,
			val.TotalDebtToCapitalization,
			val.InterestCoverage,
			val.CashFlowToDebtRatio,
			val.CompanyEquityMultiplier,
			val.ReceivablesTurnover,
			val.PayablesTurnover,
			val.InventoryTurnover,
			val.FixedAssetTurnover,
			val.AssetTurnover,
			val.OperatingCashFlowPerShare,
			val.FreeCashFlowPerShare,
			val.CashPerShare,
			val.PayoutRatio,
			val.OperatingCashFlowSalesRatio,
			val.FreeCashFlowOperatingCashFlowRatio,
			val.CashFlowCoverageRatios,
			val.ShortTermCoverageRatios,
			val.CapitalExpenditureCoverageRatio,
			val.DividendPaidAndCapexCoverageRatio,
			val.DividendPayoutRatio,
			val.PriceBookValueRatio,
			val.PriceToBookRatio,
			val.PriceEarningsRatio,
			val.PriceToFreeCashFlowsRatio,
			val.PriceToOperatingCashFlowsRatio,
			val.PriceCashFlowRatio,
			val.PriceEarningsToGrowthRatio,
			val.PriceSalesRatio,
			val.DividendYield,
			val.EnterpriseValueMultiple,
			val.PriceFairValue)
		//Sum counter
		counter_FRQ = counter_FRQ + 1
	}
	//Deleting the last nil value
	sqlStr_FRQ = sqlStr_FRQ[0 : len(sqlStr_FRQ)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: FinancialRatio AnnualTTM---------------*/
	vals_FRATTM := []interface{}{}
	sqlStr_FRATTM := `INSERT INTO FinancialRatio_AnnualTTM (
		id,
    symbol,
    dividendYielTTM,
    dividendYielPercentageTTM,
    peRatioTTM,
    pegRatioTTM,
    payoutRatioTTM,
    currentRatioTTM,
    quickRatioTTM,
    cashRatioTTM,
    daysOfSalesOutstandingTTM,
    daysOfInventoryOutstandingTTM,
    operatingCycleTTM,
    daysOfPayablesOutstandingTTM,
    cashConversionCycleTTM,
    grossProfitMarginTTM,
    operatingProfitMarginTTM,
    pretaxProfitMarginTTM,
    netProfitMarginTTM,
    effectiveTaxRateTTM,
    returnOnAssetsTTM,
    returnOnEquityTTM,
    returnOnCapitalEmployedTTM,
    netIncomePerEBTTTM,
    ebtPerEbitTTM,
    ebitPerRevenueTTM,
    debtRatioTTM,
    debtEquityRatioTTM,
    longTermDebtToCapitalizationTTM,
    totalDebtToCapitalizationTTM,
    interestCoverageTTM,
    cashFlowToDebtRatioTTM,
    companyEquityMultiplierTTM,
    receivablesTurnoverTTM,
    payablesTurnoverTTM,
    inventoryTurnoverTTM,
    fixedAssetTurnoverTTM,
    assetTurnoverTTM,
    operatingCashFlowPerShareTTM,
    freeCashFlowPerShareTTM,
    cashPerShareTTM,
    operatingCashFlowSalesRatioTTM,
    freeCashFlowOperatingCashFlowRatioTTM,
    cashFlowCoverageRatiosTTM,
    shortTermCoverageRatiosTTM,
    capitalExpenditureCoverageRatioTTM,
    dividendPaidAndCapexCoverageRatioTTM,
    priceBookValueRatioTTM,
    priceToBookRatioTTM,
    priceToSalesRatioTTM,
    priceEarningsRatioTTM,
    priceToFreeCashFlowsRatioTTM,
    priceToOperatingCashFlowsRatioTTM,
    priceCashFlowRatioTTM,
    priceEarningsToGrowthRatioTTM,
    priceSalesRatioTTM,
    dividendYieldTTM,
    enterpriseValueMultipleTTM,
    priceFairValueTTM,
    dividendPerShareTTM
	) VALUES`
	counter_FRATTM := 0
	for _, val := range frattm {
		counter_FinancialRatio_AnnualTTM = counter_FinancialRatio_AnnualTTM + 1
		//Insert data in the query
		sqlStr_FRATTM += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_FRATTM = append(vals_FRATTM, time.Now().UnixMilli()+int64(counter_FRATTM),
			symbol,
			val.DividendYielTTM,
			val.DividendYielPercentageTTM,
			val.PeRatioTTM,
			val.PegRatioTTM,
			val.PayoutRatioTTM,
			val.CurrentRatioTTM,
			val.QuickRatioTTM,
			val.CashRatioTTM,
			val.DaysOfSalesOutstandingTTM,
			val.DaysOfInventoryOutstandingTTM,
			val.OperatingCycleTTM,
			val.DaysOfPayablesOutstandingTTM,
			val.CashConversionCycleTTM,
			val.GrossProfitMarginTTM,
			val.OperatingProfitMarginTTM,
			val.PretaxProfitMarginTTM,
			val.NetProfitMarginTTM,
			val.EffectiveTaxRateTTM,
			val.ReturnOnAssetsTTM,
			val.ReturnOnEquityTTM,
			val.ReturnOnCapitalEmployedTTM,
			val.NetIncomePerEBTTTM,
			val.EbtPerEbitTTM,
			val.EbitPerRevenueTTM,
			val.DebtRatioTTM,
			val.DebtEquityRatioTTM,
			val.LongTermDebtToCapitalizationTTM,
			val.TotalDebtToCapitalizationTTM,
			val.InterestCoverageTTM,
			val.CashFlowToDebtRatioTTM,
			val.CompanyEquityMultiplierTTM,
			val.ReceivablesTurnoverTTM,
			val.PayablesTurnoverTTM,
			val.InventoryTurnoverTTM,
			val.FixedAssetTurnoverTTM,
			val.AssetTurnoverTTM,
			val.OperatingCashFlowPerShareTTM,
			val.FreeCashFlowPerShareTTM,
			val.CashPerShareTTM,
			val.OperatingCashFlowSalesRatioTTM,
			val.FreeCashFlowOperatingCashFlowRatioTTM,
			val.CashFlowCoverageRatiosTTM,
			val.ShortTermCoverageRatiosTTM,
			val.CapitalExpenditureCoverageRatioTTM,
			val.DividendPaidAndCapexCoverageRatioTTM,
			val.PriceBookValueRatioTTM,
			val.PriceToBookRatioTTM,
			val.PriceToSalesRatioTTM,
			val.PriceEarningsRatioTTM,
			val.PriceToFreeCashFlowsRatioTTM,
			val.PriceToOperatingCashFlowsRatioTTM,
			val.PriceCashFlowRatioTTM,
			val.PriceEarningsToGrowthRatioTTM,
			val.PriceSalesRatioTTM,
			val.DividendYieldTTM,
			val.EnterpriseValueMultipleTTM,
			val.PriceFairValueTTM,
			val.DividendPerShareTTM)
		//Sum counter
		counter_FRATTM = counter_FRATTM + 1
	}
	//Deleting the last nil value
	sqlStr_FRATTM = sqlStr_FRATTM[0 : len(sqlStr_FRATTM)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: KeyMetrics Annual---------------*/
	vals_KMA := []interface{}{}
	sqlStr_KMA := `INSERT INTO KeyMetrics_Annual(
		id,
   symbol,
   date,
   period,
   revenuePerShare,
   netIncomePerShare,
   operatingCashFlowPerShare,
   freeCashFlowPerShare,
   cashPerShare,
   bookValuePerShare,
   tangibleBookValuePerShare,
   shareholdersEquityPerShare,
   interestDebtPerShare,
   marketCap,
   enterpriseValue,
   peRatio,
   priceToSalesRatio,
   pocfratio,
   pfcfRatio,
   pbRatio,
   ptbRatio,
   evToSales,
   enterpriseValueOverEBITDA,
   evToOperatingCashFlow,
   evToFreeCashFlow,
   earningsYield,
   freeCashFlowYield,
   debtToEquity,
   debtToAssets,
   netDebtToEBITDA,
   currentRatio,
   interestCoverage,
   incomeQuality,
   dividendYield,
   payoutRatio,
   salesGeneralAndAdministrativeToRevenue,
   researchAndDdevelopementToRevenue,
   intangiblesToTotalAssets,
   capexToOperatingCashFlow,
   capexToRevenue,
   capexToDepreciation,
   stockBasedCompensationToRevenue,
   grahamNumber,
   roic,
   returnOnTangibleAssets,
   grahamNetNet,
   workingCapital,
   tangibleAssetValue,
   netCurrentAssetValue,
   investedCapital,
   averageReceivables,
   averagePayables,
   averageInventory,
   daysSalesOutstanding,
   daysPayablesOutstanding,
   daysOfInventoryOnHand,
   receivablesTurnover,
   payablesTurnover,
   inventoryTurnover,
   roe,
   capexPerShare
	) VALUES`
	counter_KMA := 0
	for _, val := range kma {
		counter_KeyMetrics_Annual = counter_KeyMetrics_Annual + 1
		//Insert data in the query
		sqlStr_KMA += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_KMA = append(vals_KMA, time.Now().UnixMilli()+int64(counter_KMA),
			symbol,
			val.Date,
			val.Period,
			val.RevenuePerShare,
			val.NetIncomePerShare,
			val.OperatingCashFlowPerShare,
			val.FreeCashFlowPerShare,
			val.CashPerShare,
			val.BookValuePerShare,
			val.TangibleBookValuePerShare,
			val.ShareholdersEquityPerShare,
			val.InterestDebtPerShare,
			val.MarketCap,
			val.EnterpriseValue,
			val.PeRatio,
			val.PriceToSalesRatio,
			val.Pocfratio,
			val.PfcfRatio,
			val.PbRatio,
			val.PtbRatio,
			val.EvToSales,
			val.EnterpriseValueOverEBITDA,
			val.EvToOperatingCashFlow,
			val.EvToFreeCashFlow,
			val.EarningsYield,
			val.FreeCashFlowYield,
			val.DebtToEquity,
			val.DebtToAssets,
			val.NetDebtToEBITDA,
			val.CurrentRatio,
			val.InterestCoverage,
			val.IncomeQuality,
			val.DividendYield,
			val.PayoutRatio,
			val.SalesGeneralAndAdministrativeToRevenue,
			val.ResearchAndDdevelopementToRevenue,
			val.IntangiblesToTotalAssets,
			val.CapexToOperatingCashFlow,
			val.CapexToRevenue,
			val.CapexToDepreciation,
			val.StockBasedCompensationToRevenue,
			val.GrahamNumber,
			val.Roic,
			val.ReturnOnTangibleAssets,
			val.GrahamNetNet,
			val.WorkingCapital,
			val.TangibleAssetValue,
			val.NetCurrentAssetValue,
			val.InvestedCapital,
			val.AverageReceivables,
			val.AveragePayables,
			val.AverageInventory,
			val.DaysSalesOutstanding,
			val.DaysPayablesOutstanding,
			val.DaysOfInventoryOnHand,
			val.ReceivablesTurnover,
			val.PayablesTurnover,
			val.InventoryTurnover,
			val.Roe,
			val.CapexPerShare)
		//Sum counter
		counter_KMA = counter_KMA + 1
	}
	//Deleting the last nil value
	sqlStr_KMA = sqlStr_KMA[0 : len(sqlStr_KMA)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: KeyMetrics Quarter---------------*/
	vals_KMQ := []interface{}{}
	sqlStr_KMQ := `INSERT INTO KeyMetrics_Quarter(
		id,
		symbol,
		date,
		period,
		revenuePerShare,
		netIncomePerShare,
		operatingCashFlowPerShare,
		freeCashFlowPerShare,
		cashPerShare,
		bookValuePerShare,
		tangibleBookValuePerShare,
		shareholdersEquityPerShare,
		interestDebtPerShare,
		marketCap,
		enterpriseValue,
		peRatio,
		priceToSalesRatio,
		pocfratio,
		pfcfRatio,
		pbRatio,
		ptbRatio,
		evToSales,
		enterpriseValueOverEBITDA,
		evToOperatingCashFlow,
		evToFreeCashFlow,
		earningsYield,
		freeCashFlowYield,
		debtToEquity,
		debtToAssets,
		netDebtToEBITDA,
		currentRatio,
		interestCoverage,
		incomeQuality,
		dividendYield,
		payoutRatio,
		salesGeneralAndAdministrativeToRevenue,
		researchAndDdevelopementToRevenue,
		intangiblesToTotalAssets,
		capexToOperatingCashFlow,
		capexToRevenue,
		capexToDepreciation,
		stockBasedCompensationToRevenue,
		grahamNumber,
		roic,
		returnOnTangibleAssets,
		grahamNetNet,
		workingCapital,
		tangibleAssetValue,
		netCurrentAssetValue,
		investedCapital,
		averageReceivables,
		averagePayables,
		averageInventory,
		daysSalesOutstanding,
		daysPayablesOutstanding,
		daysOfInventoryOnHand,
		receivablesTurnover,
		payablesTurnover,
		inventoryTurnover,
		roe,
		capexPerShare
	) VALUES`
	counter_KMQ := 0
	for _, val := range kmq {
		counter_KeyMetrics_Quarter = counter_KeyMetrics_Quarter + 1
		//Insert data in the query
		sqlStr_KMQ += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_KMQ = append(vals_KMQ, time.Now().UnixMilli()+int64(counter_KMQ),
			symbol,
			val.Date,
			val.Period,
			val.RevenuePerShare,
			val.NetIncomePerShare,
			val.OperatingCashFlowPerShare,
			val.FreeCashFlowPerShare,
			val.CashPerShare,
			val.BookValuePerShare,
			val.TangibleBookValuePerShare,
			val.ShareholdersEquityPerShare,
			val.InterestDebtPerShare,
			val.MarketCap,
			val.EnterpriseValue,
			val.PeRatio,
			val.PriceToSalesRatio,
			val.Pocfratio,
			val.PfcfRatio,
			val.PbRatio,
			val.PtbRatio,
			val.EvToSales,
			val.EnterpriseValueOverEBITDA,
			val.EvToOperatingCashFlow,
			val.EvToFreeCashFlow,
			val.EarningsYield,
			val.FreeCashFlowYield,
			val.DebtToEquity,
			val.DebtToAssets,
			val.NetDebtToEBITDA,
			val.CurrentRatio,
			val.InterestCoverage,
			val.IncomeQuality,
			val.DividendYield,
			val.PayoutRatio,
			val.SalesGeneralAndAdministrativeToRevenue,
			val.ResearchAndDdevelopementToRevenue,
			val.IntangiblesToTotalAssets,
			val.CapexToOperatingCashFlow,
			val.CapexToRevenue,
			val.CapexToDepreciation,
			val.StockBasedCompensationToRevenue,
			val.GrahamNumber,
			val.Roic,
			val.ReturnOnTangibleAssets,
			val.GrahamNetNet,
			val.WorkingCapital,
			val.TangibleAssetValue,
			val.NetCurrentAssetValue,
			val.InvestedCapital,
			val.AverageReceivables,
			val.AveragePayables,
			val.AverageInventory,
			val.DaysSalesOutstanding,
			val.DaysPayablesOutstanding,
			val.DaysOfInventoryOnHand,
			val.ReceivablesTurnover,
			val.PayablesTurnover,
			val.InventoryTurnover,
			val.Roe,
			val.CapexPerShare)
		//Sum counter
		counter_KMQ = counter_KMQ + 1
	}
	//Deleting the last nil value
	sqlStr_KMQ = sqlStr_KMQ[0 : len(sqlStr_KMQ)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: KeyMetrics Company TTM---------------*/
	vals_KMCTTM := []interface{}{}
	sqlStr_KMCTTM := `INSERT INTO KeyMetrics_CompanyTTM(
		id,
   symbol,
   revenuePerShareTTM,
   netIncomePerShareTTM,
   operatingCashFlowPerShareTTM,
   freeCashFlowPerShareTTM,
   cashPerShareTTM,
   bookValuePerShareTTM,
   tangibleBookValuePerShareTTM,
   shareholdersEquityPerShareTTM,
   interestDebtPerShareTTM,
   marketCapTTM,
   enterpriseValueTTM,
   peRatioTTM,
   priceToSalesRatioTTM,
   pocfratioTTM,
   pfcfRatioTTM,
   pbRatioTTM,
   ptbRatioTTM,
   evToSalesTTM,
   enterpriseValueOverEBITDATTM,
   evToOperatingCashFlowTTM,
   evToFreeCashFlowTTM,
   earningsYieldTTM,
   freeCashFlowYieldTTM,
   debtToEquityTTM,
   debtToAssetsTTM,
   netDebtToEBITDATTM,
   currentRatioTTM,
   interestCoverageTTM,
   incomeQualityTTM,
   dividendYieldTTM,
   dividendYieldPercentageTTM,
   payoutRatioTTM,
   salesGeneralAndAdministrativeToRevenueTTM,
   researchAndDevelopementToRevenueTTM,
   intangiblesToTotalAssetsTTM,
   capexToOperatingCashFlowTTM,
   capexToRevenueTTM,
   capexToDepreciationTTM,
   stockBasedCompensationToRevenueTTM,
   grahamNumberTTM,
   roicTTM,
   returnOnTangibleAssetsTTM,
   grahamNetNetTTM,
   workingCapitalTTM,
   tangibleAssetValueTTM,
   netCurrentAssetValueTTM,
   investedCapitalTTM,
   averageReceivablesTTM,
   averagePayablesTTM,
   averageInventoryTTM,
   daysSalesOutstandingTTM,
   daysPayablesOutstandingTTM,
   daysOfInventoryOnHandTTM,
   receivablesTurnoverTTM,
   payablesTurnoverTTM,
   inventoryTurnoverTTM,
   roeTTM,
   capexPerShareTTM,
   dividendPerShareTTM,
   debtToMarketCapTTM
	) VALUES`
	counter_KMCTTM := 0
	for _, val := range kmcttm {
		counter_KeyMetrics_CompanyTTM = counter_KeyMetrics_CompanyTTM + 1
		//Insert data in the query
		sqlStr_KMCTTM += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_KMCTTM = append(vals_KMCTTM, time.Now().UnixMilli()+int64(counter_KMCTTM),
			symbol,
			val.RevenuePerShareTTM,
			val.NetIncomePerShareTTM,
			val.OperatingCashFlowPerShareTTM,
			val.FreeCashFlowPerShareTTM,
			val.CashPerShareTTM,
			val.BookValuePerShareTTM,
			val.TangibleBookValuePerShareTTM,
			val.ShareholdersEquityPerShareTTM,
			val.InterestDebtPerShareTTM,
			val.MarketCapTTM,
			val.EnterpriseValueTTM,
			val.PeRatioTTM,
			val.PriceToSalesRatioTTM,
			val.PocfratioTTM,
			val.PfcfRatioTTM,
			val.PbRatioTTM,
			val.PtbRatioTTM,
			val.EvToSalesTTM,
			val.EnterpriseValueOverEBITDATTM,
			val.EvToOperatingCashFlowTTM,
			val.EvToFreeCashFlowTTM,
			val.EarningsYieldTTM,
			val.FreeCashFlowYieldTTM,
			val.DebtToEquityTTM,
			val.DebtToAssetsTTM,
			val.NetDebtToEBITDATTM,
			val.CurrentRatioTTM,
			val.InterestCoverageTTM,
			val.IncomeQualityTTM,
			val.DividendYieldTTM,
			val.DividendYieldPercentageTTM,
			val.PayoutRatioTTM,
			val.SalesGeneralAndAdministrativeToRevenueTTM,
			val.ResearchAndDevelopementToRevenueTTM,
			val.IntangiblesToTotalAssetsTTM,
			val.CapexToOperatingCashFlowTTM,
			val.CapexToRevenueTTM,
			val.CapexToDepreciationTTM,
			val.StockBasedCompensationToRevenueTTM,
			val.GrahamNumberTTM,
			val.RoicTTM,
			val.ReturnOnTangibleAssetsTTM,
			val.GrahamNetNetTTM,
			val.WorkingCapitalTTM,
			val.TangibleAssetValueTTM,
			val.NetCurrentAssetValueTTM,
			val.InvestedCapitalTTM,
			val.AverageReceivablesTTM,
			val.AveragePayablesTTM,
			val.AverageInventoryTTM,
			val.DaysSalesOutstandingTTM,
			val.DaysPayablesOutstandingTTM,
			val.DaysOfInventoryOnHandTTM,
			val.ReceivablesTurnoverTTM,
			val.PayablesTurnoverTTM,
			val.InventoryTurnoverTTM,
			val.RoeTTM,
			val.CapexPerShareTTM,
			val.DividendPerShareTTM,
			val.DebtToMarketCapTTM)
		//Sum counter
		counter_KMCTTM = counter_KMCTTM + 1
	}
	//Deleting the last nil value
	sqlStr_KMCTTM = sqlStr_KMCTTM[0 : len(sqlStr_KMCTTM)-1]
	/*---------------------------------------------------------------*/

	/*-------------------DATA: CompanyProfile---------------*/
	vals_CP := []interface{}{}
	sqlStr_CP := `INSERT INTO CompanyProfile(
	symbol,
	price,
	beta,
	volAvg,
	mktCap,
	lastDiv,
	rangeCompany,
	changesCompany,
	companyName,
	currency,
	cik,
	isin,
	cusip,
	exchange,
	exchangeShortName,
	industry,
	website,
	description,
	ceo,
	sector,
	country,
	fullTimeEmployees,
	phone,
	address,
	city,
	state,
	zip,
	dcfDiff,
	dcf,
	image,
	ipoDate,
	defaultImage,
	isEtf,
	isActivelyTrading,
	isAdr,
	isFund) VALUES`
	counter_CP := 0
	for _, val := range cp {
		counter_CompanyProfile = counter_CompanyProfile + 1
		//Insert data in the query
		sqlStr_CP += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_CP = append(vals_CP, symbol,
			val.Price,
			val.Beta,
			val.VolAvg,
			val.MktCap,
			val.LastDiv,
			val.RangeCompany,
			val.ChangesCompany,
			val.CompanyName,
			val.Currency,
			val.Cik,
			val.Isin,
			val.Cusip,
			val.Exchange,
			val.ExchangeShortName,
			val.Industry,
			val.Website,
			val.Description,
			val.Ceo,
			val.Sector,
			val.Country,
			val.FullTimeEmployees,
			val.Phone,
			val.Address,
			val.City,
			val.State,
			val.Zip,
			val.DcfDiff,
			val.Dcf,
			val.Image,
			val.IpoDate,
			val.DefaultImage,
			val.IsEtf,
			val.IsActivelyTrading,
			val.IsAdr,
			val.IsFund)
		//Sum counter
		counter_CP = counter_CP + 1
	}
	//Deleting the last nil value
	sqlStr_CP = sqlStr_CP[0 : len(sqlStr_CP)-1]
	/*---------------------------------------------------------------*/

	//BEGIN
	tx, error_tx := models.SingleStoreCN.Begin()
	if error_tx != nil {
		tx.Rollback()
		return error_tx
	}

	//TradableSymbols
	stmt_CP, _ := tx.Prepare(sqlStr_CP)
	if _, err := stmt_CP.Exec(vals_CP...); err != nil {
		return err
	}
	log.Print("LOAD COMPANY PROFILE....Done")

	if counter_IncomeStatement_Annual > 0 {
		//Income Statement Annual
		stmt_ISA, _ := tx.Prepare(sqlStr_ISA)
		if _, err := stmt_ISA.Exec(vals_ISA...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD INCOME STATEMENT ANNUAL....Done")
	}

	if counter_IncomeStatement_Quarter > 0 {
		//Income Statement Quarter
		stmt_ISQ, _ := tx.Prepare(sqlStr_ISQ)
		if _, err := stmt_ISQ.Exec(vals_ISQ...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD INCOME STATEMENT QUARTER....Done")
	}

	if counter_IncomeStatement_AnnualGrowth > 0 {
		//Income Statement Annual Growth
		stmt_ISAG, _ := tx.Prepare(sqlStr_ISAG)
		if _, err := stmt_ISAG.Exec(vals_ISAG...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD INCOME STATEMENT ANNUAL GROWTH....Done")
	}

	if counter_IncomeStatement_QuarterGrowth > 0 {
		//Income Statement Quarter Growth
		stmt_ISQG, _ := tx.Prepare(sqlStr_ISQG)
		if _, err := stmt_ISQG.Exec(vals_ISQG...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD INCOME STATEMENT QUARTER GROWTH....Done")
	}

	if counter_BalanceSheet_Annual > 0 {
		//Balance Sheet Annual
		stmt_BSA, _ := tx.Prepare(sqlStr_BSA)
		if _, err := stmt_BSA.Exec(vals_BSA...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD BALANCE SHEET ANNUAL....Done")
	}

	if counter_BalanceSheet_Quarter > 0 {
		//Balance Sheet Quarter
		stmt_BSQ, _ := tx.Prepare(sqlStr_BSQ)
		if _, err := stmt_BSQ.Exec(vals_BSQ...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD BALANCE SHEET QUARTER....Done")
	}

	if counter_IncomeStatement_AnnualGrowth > 0 {
		//Balance Sheet Annual Growth
		stmt_BSAG, _ := tx.Prepare(sqlStr_BSAG)
		if _, err := stmt_BSAG.Exec(vals_BSAG...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD BALANCE SHEET ANNUAL GROWTH....Done")
	}

	if counter_BalanceSheet_QuarterGrowth > 0 {
		//Balance Sheet Quarter Growth
		stmt_BSQG, _ := tx.Prepare(sqlStr_BSQG)
		if _, err := stmt_BSQG.Exec(vals_BSQG...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD BALANCE SHEET QUARTER GROWTH....Done")
	}

	if counter_CashFlow_Annual > 0 {
		//CashFlow Annual
		stmt_CFA, _ := tx.Prepare(sqlStr_CFA)
		if _, err := stmt_CFA.Exec(vals_CFA...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD CASHFLOW ANNUAL....Done")
	}

	if counter_CashFlow_Quarter > 0 {
		//CashFlow Quarter
		stmt_CFQ, _ := tx.Prepare(sqlStr_CFQ)
		if _, err := stmt_CFQ.Exec(vals_CFQ...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD CASHFLOW QUARTER....Done")
	}

	if counter_CashFlow_AnnualGrowth > 0 {
		//CashFlow Annual Growth
		stmt_CFAG, _ := tx.Prepare(sqlStr_CFAG)
		if _, err := stmt_CFAG.Exec(vals_CFAG...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD CASHFLOW ANNUAL GROWTH....Done")
	}

	if counter_CashFlow_QuarterGrowth > 0 {
		//CashFlow Quarter Growth
		stmt_CFQG, _ := tx.Prepare(sqlStr_CFQG)
		if _, err := stmt_CFQG.Exec(vals_CFQG...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD CASHFLOW QUARTER GROWTH....Done")
	}

	if counter_FinancialRatio_Annual > 0 {
		//Financial Ratio Annual
		stmt_FRA, _ := tx.Prepare(sqlStr_FRA)
		if _, err := stmt_FRA.Exec(vals_FRA...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD FINANCIAL RATIO ANNUAL....Done")
	}

	if counter_FinancialRatio_Quarter > 0 {
		//Financial Ratio Quarter
		stmt_FRQ, _ := tx.Prepare(sqlStr_FRQ)
		if _, err := stmt_FRQ.Exec(vals_FRQ...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD FINANCIAL RATIO QUARTER....Done")
	}

	if counter_FinancialRatio_AnnualTTM > 0 {
		//Financial Ratio Annual TTM
		stmt_FRATTM, _ := tx.Prepare(sqlStr_FRATTM)
		if _, err := stmt_FRATTM.Exec(vals_FRATTM...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD FINANCIAL RATIO ANNUAL TTM....Done")
	}

	if counter_KeyMetrics_Annual > 0 {
		//Key Metrics Annual
		stmt_KMA, _ := tx.Prepare(sqlStr_KMA)
		if _, err := stmt_KMA.Exec(vals_KMA...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD KEY METRICS ANNUAL....Done")
	}

	if counter_KeyMetrics_Quarter > 0 {
		//Key Metrics Quarter
		stmt_KMQ, _ := tx.Prepare(sqlStr_KMQ)
		if _, err := stmt_KMQ.Exec(vals_KMQ...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD KEY METRICS QUARTER....Done")
	}

	if counter_KeyMetrics_CompanyTTM > 0 {
		//Key Metrics Company TTM
		stmt_KMCTTM, _ := tx.Prepare(sqlStr_KMCTTM)
		if _, err := stmt_KMCTTM.Exec(vals_KMCTTM...); err != nil {
			tx.Rollback()
			return err
		}
		log.Print("LOAD KEY METRICS COMPANY TTM....Done")
	}
	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit()
	if err_commit != nil {
		tx.Rollback()
		return err_commit
	}

	log.Print("-------->SUCCESSFUL LOADING DATA PROCESS")

	return nil
}
