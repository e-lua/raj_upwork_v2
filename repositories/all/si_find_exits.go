package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find(symbol string) (string, error) {

	//Variable to get data
	var symbol_output string

	db := models.SingleStoreCN
	q := `SELECT symbol FROM CompanyProfile2 WHERE symbol = ?`
	error_show := db.QueryRow(q, symbol).Scan(&symbol_output)
	if error_show != nil {
		return symbol_output, error_show
	}

	return symbol_output, nil
}
