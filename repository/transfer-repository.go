// Repository package contains repository operations
// for account and transfer models
package repository

import (
	"desafio-tecnico/model"
	"errors"
)

// Create Transfer
// Receives a transfer and insert into the database
// If the operation is successful, returns nil
// If the operation fails, returns an error
func CreateTransfer(transfer model.Transfer) error {

	db := StartDB()

	sqlStatement := `INSERT INTO desafiotecnicoprincipal.transfers
		(account_origin_id, account_destination_id, amount, created_at)
		VALUES ($1, $2, $3, $4) RETURNING account_origin_id;`

	var accOriginId int

	err := db.QueryRow(sqlStatement, transfer.Account_origin_id, transfer.Account_destination_id,
		transfer.Amount, transfer.Created_at).Scan(&accOriginId)
	if err != nil || transfer.Account_origin_id != accOriginId {
		return errors.New("Error while try to create the transfer!")
	}

	return nil
}

// Read Transfers
// Receives an account id and searches for all the transfers of that account
// If the operation is successful, returns the transfers and nil
// If the operation fails, returns nil and and an error
func ReadTransfers(accountId int) ([]model.Transfer, error) {

	db := StartDB()

	sqlStatement := `SELECT id, account_origin_id, account_destination_id, amount, created_at
		FROM desafiotecnicoprincipal.transfers WHERE account_origin_id=$1;`

	rows, err := db.Query(sqlStatement, accountId)
	if err != nil {
		return nil, errors.New("Error while try to read transfers!")
	}
	defer rows.Close()

	var transfer model.Transfer
	var transfers []model.Transfer

	for rows.Next() {
		err := rows.Scan(&transfer.ID, &transfer.Account_origin_id, &transfer.Account_destination_id,
			&transfer.Amount, &transfer.Created_at)
		if err != nil {
			return nil, errors.New("Error while try to scan a row of transfers!")
		}

		transfers = append(transfers, transfer)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.New("Unexpected error in accounts row!")
	}

	if len(transfers) == 0 {
		return nil, errors.New("This account doesn't have transfers!")
	}

	return transfers, nil
}
