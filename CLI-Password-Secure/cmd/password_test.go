package password_test

import (
	"testing"
	"time"

	password "github.com/beto20/CLI-Password-Encrypt/cmd"
)

func TestRecievePassword_thenOk(t *testing.T) {
	// Initialize 
	pssw := password.PasswordStruct{}
	
	// Data mock
	inputMock := password.PasswordInput{
		Key:      "key",
		Password: "password",
	}

	var mock [1]password.PasswordStruct
	mock[0] = password.PasswordStruct{
		Id:                1,
		Key:               "key",
		PasswordEncrypted: "password000",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	// Method tested
	pssw.RecievePassword(inputMock)

	// Assertions
	res1 := len(password.Array)
	expected1 := len(mock)

	if res1 != expected1 {
		t.Errorf("res %v, expected %v", res1, expected1)
	}

	res2 := password.Array[0].Key
	expected2 := mock[0].Key

	if res2 != expected2 {
		t.Errorf("res  %v, expected %v", res2, expected2)
	}
}
