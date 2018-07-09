package test

import (
	"fmt"
	"testing"
	"loginradius"
)

func TestGetOneTouchLoginByEmail(t *testing.T) {
	fmt.Println("Starting test TestGetOneTouchLoginByEmail")
	_, _, _, email, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	session, err := loginradius.GetOneTouchLoginByEmail(email, "", createClientGUID(), "", "", "")
	if err != nil && session.IsPosted != true {
		t.Errorf("Error sending one touch email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetOneTouchLoginByPhone(t *testing.T) {
	fmt.Println("Starting test TestGetOneTouchLoginByPhone")
	phoneID, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	fmt.Println(phoneID)
	_, err := loginradius.GetOneTouchLoginByPhone(phoneID, "", "")
	if err != nil {
		t.Errorf("Error sending one touch sms")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
