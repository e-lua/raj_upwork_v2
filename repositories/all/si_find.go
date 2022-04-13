package all

import (
	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Find_IndustriesAndSector() (models.IndustryAndSector, error) {

	//Variable to get data
	var ouput models.IndustryAndSector

	db := models.SingleStoreCN
	q := `SELECT json_agg(industry) AS industries,json_agg(sector) AS sectors FROM CompanyProfile GROUP BY industry,sector`
	error_show := db.QueryRow(q).Scan(&ouput.Industries, &ouput.Sectors)
	if error_show != nil {
		return ouput, error_show
	}

	return ouput, nil
}
