package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find() ([]models.TradableSymbols, error) {

	db := models.SingleStoreCN
	q := "SELECT symbol,name,price,exchange,exchangeShortName FROM TradableSymbols"
	rows, error_show := db.Query(q)

	oListTS := []models.TradableSymbols{}

	if error_show != nil {

		return oListTS, error_show
	}

	for rows.Next() {
		oTS := models.TradableSymbols{}
		rows.Scan(&oTS.Symbol,
			&oTS.Name,
			&oTS.Price,
			&oTS.Exchange,
			&oTS.ExchangeShortName,
		)
		oListTS = append(oListTS, oTS)
	}

	//OK
	return oListTS, nil
}
