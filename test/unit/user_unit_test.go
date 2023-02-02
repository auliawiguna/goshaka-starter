package unit_test

import (
	"goshaka/app/models"
	repositories_v1 "goshaka/app/repositories"
	"goshaka/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errConnectDb = database.Connect()

// Test code can connect to database
func TestCanConnectToDatabase(t *testing.T) {
	assert.Equal(t, errConnectDb, nil, "No error")
}

// Test code can get user from database
func TestGetUserByEmail(t *testing.T) {

	var user models.User
	var isFound bool
	user, isFound = repositories_v1.FindByEmail("aulia@goshaka.id")

	assert.Equal(t, isFound, true, "No error")
	assert.Equal(t, user.Email, "aulia@goshaka.id", "Email is found")
}

// Test code can get user from database
func TestGetUserById(t *testing.T) {

	var user models.User
	user = repositories_v1.FindById(1)

	assert.Equal(t, user.Email, "aulia@goshaka.id", "Email is found")
}
