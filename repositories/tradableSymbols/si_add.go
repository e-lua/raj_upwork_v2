package all

import (
	"log"
	"time"

	models "github.com/Aphofisis/raj_upwork_v2/models"
)

func Si_Add(ts []models.TradableSymbols) (error, interface{}, interface{}) {

	/*-------------------DATA: TradableSymbols---------------*/
	vals_TS := []interface{}{}
	sqlStr_TS := `INSERT INTO TradableSymbols(id,
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
		vals_TS = append(vals_TS, time.Now().UnixMilli()+int64(counter_TS), val.Symbol,
			val.Name,
			val.Price,
			val.Exchange,
			val.ExchangeShortName)

		if counter_TS > 3 {
			break
		}
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
		return error_tx, sqlStr_TS, vals_TS
	}

	//TradableSymbols
	stmt_TS, _ := tx.Prepare(sqlStr_TS)
	if _, err := stmt_TS.Exec(vals_TS...); err != nil {
		return err, sqlStr_TS, vals_TS
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit()
	if err_commit != nil {
		tx.Rollback()
		return err_commit, sqlStr_TS, vals_TS
	}

	log.Print("LOAD INCOME STATEMENT ANNUAL....Done")

	return nil, sqlStr_TS, vals_TS
}
