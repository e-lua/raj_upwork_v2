package all

import (
	"log"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Add(cp []models.CompanyProfile) error {

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

		//Insert data in the query
		sqlStr_CP += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		//Assign the data to the query
		vals_CP = append(vals_CP, val.Symbol,
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

	db := models.SingleStoreCN

	//TradableSymbols
	stmt_CP, _ := db.Prepare(sqlStr_CP)
	if _, err := stmt_CP.Exec(vals_CP...); err != nil {
		return err
	}
	log.Print("LOAD COMPANY PROFILE....Done")

	return nil
}
