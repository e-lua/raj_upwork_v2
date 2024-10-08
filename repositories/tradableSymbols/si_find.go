package all

import (
	"strconv"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find() ([]models.TradableSymbols, error) {

	db := models.SingleStoreCN
	q := `SELECT symbol FROM TradableSymbols`
	rows, error_show := db.Query(q)

	oListTS := []models.TradableSymbols{}

	if error_show != nil {

		return oListTS, error_show
	}

	for rows.Next() {
		oTS := models.TradableSymbols{}
		rows.Scan(&oTS.Symbol)
		oListTS = append(oListTS, oTS)
	}

	//OK
	return oListTS, nil
}

func Si_Find_WithLimit(index int) ([]models.TradableSymbols, error) {

	db := models.SingleStoreCN
	q := `SELECT symbol FROM TradableSymbols ORDER BY symbol ASC LIMIT ` + strconv.Itoa(10*index) + "," + strconv.Itoa(10)
	rows, error_show := db.Query(q)

	oListTS := []models.TradableSymbols{}

	if error_show != nil {

		return oListTS, error_show
	}

	for rows.Next() {
		oTS := models.TradableSymbols{}
		rows.Scan(&oTS.Symbol)
		oListTS = append(oListTS, oTS)
	}

	//OK
	return oListTS, nil
}
