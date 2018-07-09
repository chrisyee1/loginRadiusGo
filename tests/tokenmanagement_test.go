package test

import (
	"fmt"
	"os"
	"testing"
	"loginradius"
)

func TestGetAccessTokenViaFacebook(t *testing.T) {
	fmt.Println("Starting test TestGetAccessTokenViaFacebook")
	PresetLoginRadiusTestEnv()
	session, err := loginradius.GetAccessTokenViaFacebook(os.Getenv("FACEBOOKTOKEN"))
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error retrieving facebook token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAccessTokenViaTwitter(t *testing.T) {
	fmt.Println("Starting test TestGetAccessTokenViaTwitter")
	PresetLoginRadiusTestEnv()
	session, err := loginradius.GetAccessTokenViaTwitter(os.Getenv("TWITTERTOKEN"), os.Getenv("TWITTERSECRET"))
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error retrieving twitter token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAccessTokenViaVkontakte(t *testing.T) {
	fmt.Println("Starting test TestGetAccessTokenViaVkontakte")
	PresetLoginRadiusTestEnv()
	session, err := loginradius.GetAccessTokenViaVkontakte(os.Getenv("VKONTAKTETOKEN"))
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error retrieving facebook token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetRefreshUserProfile(t *testing.T) {
	fmt.Println("Starting test TestGetRefreshUserProfile")
	PresetLoginRadiusTestEnv()
	session, err := loginradius.GetAccessTokenViaFacebook(os.Getenv("FACEBOOKTOKEN"))
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error retrieving facebook token")
		fmt.Println(err)
	}
	_, err2 := loginradius.GetRefreshUserProfile(session.AccessToken)
	if err2 != nil {
		t.Errorf("Error refreshing profile")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestGetRefreshToken(t *testing.T) {
	fmt.Println("Starting test TestGetRefreshToken")
	PresetLoginRadiusTestEnv()
	session, err := loginradius.GetAccessTokenViaTwitter(os.Getenv("TWITTERTOKEN"), os.Getenv("TWITTERSECRET"))
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error retrieving twitter token")
		fmt.Println(err)
	}
	_, err2 := loginradius.GetRefreshToken(session.AccessToken)
	if err2 != nil {
		t.Errorf("Error refreshing token")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}
