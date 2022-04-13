package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_Profile(symbol string) (models.CompanyProfile, error) {

	db := models.SingleStoreCN

	var profile models.CompanyProfile

	q := `SELECT symbol,price,beta,volAvg,mktCap,lastDiv,rangeCompany,changesCompany,companyName,currency,cik,isin,cusip,exchange,exchangeShortName,industry,website,description,ceo,sector,country,fullTimeEmployees,phone,address,city,state,zip,dcfDiff,dcf,image,ipoDate,defaultImage,isEtf,isActivelyTrading,isAdr,isFund FROM CompanyProfile WHERE symbol = ?`
	error_show := db.QueryRow(q, symbol).Scan(&profile.Symbol, &profile.Price, &profile.Beta, &profile.VolAvg, &profile.MktCap, &profile.LastDiv, &profile.RangeCompany, &profile.ChangesCompany, &profile.CompanyName, &profile.Currency, &profile.Cik, &profile.Isin, &profile.Cusip, &profile.Exchange, &profile.ExchangeShortName, &profile.Industry, &profile.Website, &profile.Description, &profile.Ceo, &profile.Sector, &profile.Country, &profile.FullTimeEmployees, &profile.Phone, &profile.Address, &profile.City, &profile.State, &profile.Zip, &profile.DcfDiff, &profile.Dcf, &profile.Image, &profile.IpoDate, &profile.DefaultImage, &profile.IsEtf, &profile.IsActivelyTrading, &profile.IsAdr, &profile.IsFund)

	if error_show != nil {
		return profile, error_show
	}

	//OK
	return profile, nil
}
