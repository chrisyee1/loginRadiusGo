package test

import (
	"fmt"
	"testing"
	"loginradius"
)

func TestPostWebhookSubscribe(t *testing.T) {
	fmt.Println("Starting test TestPostWebhookSubscribe")
	PresetLoginRadiusTestEnv()
	webhook := WebhookTest{"https://www.google.ca", "Register"}
	resp, err := loginradius.PostWebhookSubscribe(webhook)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error setting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetWebhookTest(t *testing.T) {
	fmt.Println("Starting test TestGetWebhookTest")
	PresetLoginRadiusTestEnv()
	_, err := loginradius.GetWebhookTest()
	if err != nil {
		t.Errorf("Error getting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetWebhookSubscribedURLs(t *testing.T) {
	fmt.Println("Starting test TestGetWebhookSubscribedURLs")
	PresetLoginRadiusTestEnv()
	_, err := loginradius.GetWebhookSubscribedURLs("Register")
	if err != nil {
		t.Errorf("Error getting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteWebhookUnsubscribe(t *testing.T) {
	fmt.Println("Starting test TestDeleteWebhookUnsubscribe")
	PresetLoginRadiusTestEnv()
	webhook := WebhookTest{"https://www.google.ca", "Register"}
	_, err := loginradius.DeleteWebhookUnsubscribe(webhook)
	if err != nil {
		t.Errorf("Error getting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
