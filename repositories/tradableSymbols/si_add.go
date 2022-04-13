package all

import (
	"log"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Add(ts []models.TradableSymbols) error {

	/*-------------------DATA: TradableSymbols---------------*/
	vals_TS := []interface{}{}
	sqlStr_TS := `INSERT INTO TradableSymbols(
		symbol,
		name,
		price,
		exchange,
		exchangeShortName) VALUES`
	counter_TS := 0
	for _, val := range ts {

		//Insert data in the query
		sqlStr_TS += "(?,?,?,?,?),"
		//Assign the data to the query
		vals_TS = append(vals_TS, val.Symbol,
			val.Name,
			val.Price,
			val.Exchange,
			val.ExchangeShortName)
		//Sum counter
		counter_TS = counter_TS + 1
	}
	//Deleting the last nil value
	sqlStr_TS = sqlStr_TS[0 : len(sqlStr_TS)-1]
	/*---------------------------------------------------------------*/

	//BEGIN
	tx, error_tx := models.SingleStoreCN.Begin()
	if error_tx != nil {
		tx.Rollback()
		return error_tx
	}

	//TradableSymbols
	stmt_TS, _ := tx.Prepare(sqlStr_TS)
	if _, err := stmt_TS.Exec(vals_TS...); err != nil {
		return err
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit()
	if err_commit != nil {
		tx.Rollback()
		return err_commit
	}

	log.Print("LOAD INCOME STATEMENT ANNUAL....Done")

	return nil
}
