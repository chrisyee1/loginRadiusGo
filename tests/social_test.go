package test

import (
	"fmt"
	"os"
	"testing"
	"time"
	"loginradius"
)

func setupSocialTests(t *testing.T) (string, string, string) {
	PresetLoginRadiusTestEnv()
	fbSession, err := loginradius.GetAccessTokenViaFacebook(os.Getenv("FACEBOOKTOKEN"))
	facebook := fbSession.AccessToken
	if err != nil || facebook == "" {
		t.Errorf("Error retrieving facebook token")
		fmt.Println(err)
	}
	twSession, err := loginradius.GetAccessTokenViaTwitter(os.Getenv("TWITTERTOKEN"), os.Getenv("TWITTERSECRET"))
	twitter := twSession.AccessToken
	if err != nil || twitter == "" {
		t.Errorf("Error retrieving twitter token")
		fmt.Println(err)
	}
	vkSession, err := loginradius.GetAccessTokenViaVkontakte(os.Getenv("VKONTAKTETOKEN"))
	vkontakte := vkSession.AccessToken
	if err != nil || vkontakte == "" {
		t.Errorf("Error retrieving vkontakte token")
		fmt.Println(err)
	}
	return vkontakte, facebook, twitter
}

func TestPostSocialMessageAPI(t *testing.T) {
	fmt.Println("Starting test TestPostSocialMessageAPI")
	_, _, twitter := setupSocialTests(t)
	session, err := loginradius.PostSocialMessageAPI(twitter, os.Getenv("SOCIALMESSAGE"), "hello", "I am messaging you as a test")
	if err != nil || session.IsPosted != true {
		t.Errorf("Error sending message to user")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostSocialStatusPosting(t *testing.T) {
	fmt.Println("Starting test TestPostSocialStatusPosting")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	_, _, twitter := setupSocialTests(t)
	_, err := loginradius.PostSocialStatusPosting(twitter, timestamp, timestamp, "Test Image", timestamp, "Test Caption", "Test Desc")
	if err != nil {
		t.Errorf("Error sending message to provider")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialAccessToken(t *testing.T) {
	fmt.Println("Starting test TestGetSocialAccessToken")
	_, _, twitter := setupSocialTests(t)
	_, err := loginradius.GetSocialAccessToken(twitter)
	if err != nil {
		t.Errorf("Error setting up token for other APIs")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialTokenValidate(t *testing.T) {
	fmt.Println("Starting test TestGetSocialTokenValidate")
	_, _, twitter := setupSocialTests(t)
	_, err := loginradius.GetSocialTokenValidate(twitter)
	if err != nil {
		t.Errorf("Error validating token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialAccessTokenInvalidate(t *testing.T) {
	fmt.Println("Starting test TestGetSocialAccessTokenInvalidate")
	_, _, twitter := setupSocialTests(t)
	_, err := loginradius.GetSocialAccessTokenInvalidate(twitter)
	if err != nil {
		t.Errorf("Error invalidating token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialAlbum(t *testing.T) {
	fmt.Println("Starting test TestGetSocialAlbum")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialAlbum(facebook)
	if err != nil {
		t.Errorf("Error getting album")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialAudio(t *testing.T) {
	fmt.Println("Starting test TestGetSocialAudio")
	vkontakte, _, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialAudio(vkontakte)
	if err != nil {
		t.Errorf("Error getting audio")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialCheckin(t *testing.T) {
	fmt.Println("Starting test TestGetSocialCheckin")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialCheckin(facebook)
	if err != nil {
		t.Errorf("Error getting checkin")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialCompany(t *testing.T) {
	fmt.Println("Starting test TestGetSocialCompany")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialCompany(facebook)
	if err != nil {
		t.Errorf("Error getting company")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialContact(t *testing.T) {
	fmt.Println("Starting test TestGetSocialContact")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialContact(facebook)
	if err != nil {
		t.Errorf("Error getting contacts")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialEvent(t *testing.T) {
	fmt.Println("Starting test TestGetSocialEvent")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialEvent(facebook)
	if err != nil {
		t.Errorf("Error getting events")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialFollowing(t *testing.T) {
	fmt.Println("Starting test TestGetSocialFollowing")
	_, _, twitter := setupSocialTests(t)
	_, err := loginradius.GetSocialFollowing(twitter)
	if err != nil {
		t.Errorf("Error getting following")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialGroup(t *testing.T) {
	fmt.Println("Starting test TestGetSocialGroup")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialGroup(facebook)
	if err != nil {
		t.Errorf("Error getting groups")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialLike(t *testing.T) {
	fmt.Println("Starting test TestGetSocialLike")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialLike(facebook)
	if err != nil {
		t.Errorf("Error getting likes")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialMention(t *testing.T) {
	fmt.Println("Starting test TestGetSocialMention")
	_, _, twitter := setupSocialTests(t)
	_, err := loginradius.GetSocialMention(twitter)
	if err != nil {
		t.Errorf("Error getting mention")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialMessageAPI(t *testing.T) {
	fmt.Println("Starting test TestPostSocialMessageAPI")
	_, _, twitter := setupSocialTests(t)
	session, err := loginradius.PostSocialMessageAPI(twitter, os.Getenv("SOCIALMESSAGE"), "hello", "I am messaging you as a test")
	if err != nil || session.IsPosted != true {
		t.Errorf("Error sending message to user")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialPage(t *testing.T) {
	fmt.Println("Starting test TestGetSocialPage")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialPage(facebook, os.Getenv("SOCIALPAGEID"))
	if err != nil {
		t.Errorf("Error getting page")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialPhoto(t *testing.T) {
	fmt.Println("Starting test TestGetSocialPhoto")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialPhoto(facebook, "")
	if err != nil {
		t.Errorf("Error getting photos")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialPost(t *testing.T) {
	fmt.Println("Starting test TestGetSocialPost")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialPost(facebook)
	if err != nil {
		t.Errorf("Error getting mention")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialStatusFetch(t *testing.T) {
	fmt.Println("Starting test TestGetSocialStatusFetch")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialStatusFetch(facebook)
	if err != nil {
		t.Errorf("Error getting mention")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialStatusPosting(t *testing.T) {
	fmt.Println("Starting test TestGetSocialStatusPosting")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	_, _, twitter := setupSocialTests(t)
	_, err := loginradius.GetSocialStatusPost(twitter, "Test Title", "Test URL", "Test Image", timestamp, timestamp, timestamp)
	if err != nil {
		t.Errorf("Error sending message to provider")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialUserProfile(t *testing.T) {
	fmt.Println("Starting test TestGetSocialUserProfile")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialUserProfile(facebook)
	if err != nil {
		t.Errorf("Error getting profile")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSocialVideo(t *testing.T) {
	fmt.Println("Starting test TestGetSocialVideo")
	_, facebook, _ := setupSocialTests(t)
	_, err := loginradius.GetSocialVideo(facebook, "")
	if err != nil {
		t.Errorf("Error getting mention")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
