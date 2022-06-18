package security

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Generate Token
// Valid Account Id - Generate token with valid account id
func TestGenerateToken_ValidAccountId(t *testing.T) {

	assertions := require.New(t)

	validAccountId := 1

	token, expirationTime, err := GenerateToken(validAccountId)

	assertions.NotEqual(token, "")
	assertions.NotEqual(expirationTime, 0)
	assertions.Empty(err)
}

// Generate Token
// Invalid Account Id - Generate token with invalid account id
func TestGenerateToken_InvalidAccountId(t *testing.T) {

	assertions := require.New(t)

	invalidAccountId := 0

	token, expirationTime, err := GenerateToken(invalidAccountId)

	assertions.Equal(token, "")
	assertions.Equal(expirationTime, 0)
	assertions.Equal(err.Error(), "Error while try to create token!")
}

// Validate Token
// Valid Token - Validates a valid token
func TestValidateToken_ValidToken(t *testing.T) {

	assertions := require.New(t)

	validAccountId := 1

	token, expirationTime, err := GenerateToken(validAccountId)

	assertions.NotEqual(token, "")
	assertions.NotEqual(expirationTime, 0)
	assertions.Empty(err)

	accountId, err := ValidateToken(token)

	assertions.Equal(validAccountId, accountId)
	assertions.Empty(err)
}

// Validate Token
// Invalid Token - Validates a invalid token
func TestValidateToken_InvalidToken(t *testing.T) {

	assertions := require.New(t)

	invalidAccountId := 0

	token, expirationTime, err := GenerateToken(invalidAccountId)

	assertions.Equal(token, "")
	assertions.Equal(expirationTime, 0)
	assertions.Equal(err.Error(), "Error while try to create token!")

	accountId, err := ValidateToken(token)

	assertions.Equal(accountId, 0)
	assertions.Equal(err.Error(), "Invalid token!")
}
