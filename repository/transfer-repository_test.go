package repository

import (
	"desafio-tecnico/model"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create Transfer
// Valid Transfer - New transfer with valid values
func TestCreateTransfer_ValidTransfer(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{
		Account_origin_id: 1, Account_destination_id: 2, Amount: 1, Created_at: 1}

	err := CreateTransfer(validTransfer)

	assertions.Empty(err)
}

// Create Transfer
// Invalid Transfer - New transfer with invalid values
func TestCreateTransfer_InvalidTransfer(t *testing.T) {

	assertions := require.New(t)

	invalidTransfer := []model.Transfer{
		{Account_origin_id: 1, Account_destination_id: 1, Amount: 1, Created_at: 1},
		{Account_origin_id: 1, Account_destination_id: 2, Amount: 0, Created_at: 1},
		{Account_origin_id: 1, Account_destination_id: 2, Amount: 1, Created_at: 0},
		{},
	}

	for _, invalidTrf := range invalidTransfer {
		err := CreateTransfer(invalidTrf)
		assertions.Equal(err.Error(), "Error while try to create the transfer!")
	}
}

// Read Transfers
// Valid Account - Read transfers for a valid account
func TestReadTransfers_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	accountId := 1

	transfers, err := ReadTransfers(accountId)

	assertions.NotEmpty(transfers)
	assertions.Empty(err)
}

// Read Transfers
// Invalid Account - Try to read transfers for an invalid account
func TestReadTransfers_InvalidAccount(t *testing.T) {

	assertions := require.New(t)

	invalidAccountId := 999

	transfers, err := ReadTransfers(invalidAccountId)

	assertions.Empty(transfers)
	assertions.Equal(err.Error(), "This account doesn't have transfers!")
}
