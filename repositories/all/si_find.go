package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_IndustriesAndSector() ([]models.IndustryAndSector, error) {

	//Variable to get data
	var ouput []models.IndustryAndSector

	db := models.SingleStoreCN
	q := `SELECT industry,sector FROM CompanyProfile GROUP BY industry,sector`
	rows, error_show := db.Query(q)

	if error_show != nil {

		return ouput, error_show
	}

	for rows.Next() {
		oIS := models.IndustryAndSector{}
		rows.Scan(&oIS.Industries, &oIS.Sectors)
		ouput = append(ouput, oIS)
	}

	return ouput, nil
}
