package service

import (
	"desafio-tecnico/model"
	"desafio-tecnico/security"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create Transfer
// Valid Transfer - New transfer with valid values
func TestCreateTransfer_ValidTransfer(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{Account_origin_id: 61, Account_destination_id: 63, Amount: 1}

	token, _, err := security.GenerateToken(validTransfer.Account_origin_id)

	err = CreateTransfer(token, validTransfer.Account_destination_id, validTransfer.Amount)

	assertions.Empty(err)
}

// Create Transfer
// Invalid Token - New transfer with invalid token
func TestCreateTransfer_InvalidToken(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{Account_origin_id: 61, Account_destination_id: 63, Amount: 1}

	err := CreateTransfer("", validTransfer.Account_destination_id, validTransfer.Amount)

	assertions.Equal(err.Error(), "Invalid token!")
}

// Create Transfer
// Invalid Transfer Values - New transfer with invalid transfer values
func TestCreateTransfer_InvalidTransferValues(t *testing.T) {

	assertions := require.New(t)

	invalidTransfer := []model.Transfer{
		{Account_origin_id: 61, Account_destination_id: 61, Amount: 1},
		{Account_origin_id: 61, Account_destination_id: 63, Amount: 0},
		{Account_origin_id: 61, Account_destination_id: 999, Amount: 1},
	}

	token, _, _ := security.GenerateToken(invalidTransfer[0].Account_origin_id)
	err := CreateTransfer(token, invalidTransfer[0].Account_destination_id, invalidTransfer[0].Amount)

	assertions.Equal(err.Error(), "The destination account is the same as origin account!")

	token, _, _ = security.GenerateToken(invalidTransfer[1].Account_origin_id)
	err = CreateTransfer(token, invalidTransfer[1].Account_destination_id, invalidTransfer[1].Amount)

	assertions.Equal(err.Error(), "Invalid amount to transfer!")

	token, _, _ = security.GenerateToken(invalidTransfer[2].Account_origin_id)
	err = CreateTransfer(token, invalidTransfer[2].Account_destination_id, invalidTransfer[2].Amount)

	assertions.Equal(err.Error(), "Account id not found!")
}

// Create Transfer
// Invalid Account Balance - New transfer with invalid account balance
func TestCreateTransfer_InvalidAccountBalance(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{Account_origin_id: 72, Account_destination_id: 63, Amount: 1}

	token, _, err := security.GenerateToken(validTransfer.Account_origin_id)

	err = CreateTransfer(token, validTransfer.Account_destination_id, validTransfer.Amount)

	assertions.Equal(err.Error(), "Invalid account balance! You have 0.00")
}

// Create Transfer
// Invalid Amount To Transfer - New transfer with invalid amount to transfer
func TestCreateTransfer_InvalidAmountToTransfer(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{Account_origin_id: 73, Account_destination_id: 72, Amount: 10}

	token, _, err := security.GenerateToken(validTransfer.Account_origin_id)

	err = CreateTransfer(token, validTransfer.Account_destination_id, validTransfer.Amount)

	assertions.Equal(err.Error(), "Invalid amount to transfer! You have 0.01 and wants to transfer 0.10")
}

// Read Transfers
// Valid Account - Try to read transfers from a valid account
func TestReadTransfers_ValidAccountId(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{Account_origin_id: 73}

	token, _, err := security.GenerateToken(validTransfer.Account_origin_id)

	transfers, err := ReadTransfers(token)

	assertions.NotEmpty(transfers)
	assertions.Empty(err)
}

// Read Transfers
// Invalid Token - Try to read transfers with an invalid token
func TestReadTransfers_InvalidToken(t *testing.T) {

	assertions := require.New(t)

	transfers, err := ReadTransfers("")

	assertions.Empty(transfers)
	assertions.Equal(err.Error(), "Invalid token!")
}

// Read Transfers
// Invalid Account Id - Try to read transfers with an invalid account id
func TestReadTransfers_InvalidAccountId(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{Account_origin_id: 999}

	token, _, err := security.GenerateToken(validTransfer.Account_origin_id)

	transfers, err := ReadTransfers(token)

	assertions.Empty(transfers)
	assertions.Equal(err.Error(), "Account id not found!")
}

// Read Transfers
// Transfers Not Registred - Try to read transfers from an account without transfers
func TestReadTransfers_TransfersNotRegistred(t *testing.T) {

	assertions := require.New(t)

	validTransfer := model.Transfer{Account_origin_id: 72}

	token, _, err := security.GenerateToken(validTransfer.Account_origin_id)

	transfers, err := ReadTransfers(token)

	assertions.Empty(transfers)
	assertions.Equal(err.Error(), "This account doesn't have transfers!")
}
