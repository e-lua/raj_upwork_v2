package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_Industries() ([]string, error) {

	//Variable to get data
	var ouput []string

	db := models.SingleStoreCN
	q := `SELECT industry FROM CompanyProfile  WHERE industry<>'' GROUP BY industry`
	rows, error_show := db.Query(q)

	if error_show != nil {

		return ouput, error_show
	}

	for rows.Next() {
		var oIS string
		rows.Scan(&oIS)
		ouput = append(ouput, oIS)
	}

	return ouput, nil
}

func Si_Find_Sectors() ([]string, error) {

	//Variable to get data
	var ouput []string

	db := models.SingleStoreCN
	q := `SELECT sector FROM CompanyProfile  WHERE industry<>'' GROUP BY sector`
	rows, error_show := db.Query(q)

	if error_show != nil {

		return ouput, error_show
	}

	for rows.Next() {
		var oIS string
		rows.Scan(&oIS)
		ouput = append(ouput, oIS)
	}

	return ouput, nil
}
