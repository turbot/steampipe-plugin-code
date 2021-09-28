package secrets

import (
	"fmt"
	"log"
	"testing"
)

// Run all tests
// go test -v -run $*^  /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets

// go test -v -run TestSlackTokenAuth /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestSlackTokenAuthenticate(t *testing.T) {
	fmt.Println()
	slack_api_token := "xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65"

	for _, sm := range Matchers() {
		if sm.Type() == "slack_api_token" {
			result, err := sm.Authenticate(slack_api_token, slack_api_token)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestSlackTokenAuth authentication result: [%s]", *result)
		}

	}
}

// go test -v -run TestStripeAuthOk /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestStripeAuthOk(t *testing.T) {
	fmt.Println()
	stripeApiKey := "sk_test_tR3PYbcVNZZ796tH88S4VQ2u"

	for _, sm := range Matchers() {
		if sm.Type() == "stripe_api_key" {
			result, err := sm.Authenticate(stripeApiKey, stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestStripeAuthOk authentication result: [%s]", *result)
		}

	}
}

// go test -v -run TestStripeAuthFail /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestStripeAuthFail(t *testing.T) {
	fmt.Println()
	stripeApiKey := "sk_live_tR3PYbcVNZZ796tH88S4VQ2u"

	for _, sm := range Matchers() {
		if sm.Type() == "stripe_api_key" {
			result, err := sm.Authenticate(stripeApiKey, stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestStripeAuthFail authentication result: [%s]", *result)
		}
	}
}

// go test -v -run TestMailChimpAuthFalse /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestMailChimpAuthFalse(t *testing.T) {
	fmt.Println()
	mailchimpAccessKey := "a11b1d1baf01fd5556666f434g9b123a-us5"

	for _, sm := range Matchers() {
		if sm.Type() == "mailchimp_access_key" {
			result, err := sm.Authenticate(mailchimpAccessKey, mailchimpAccessKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestMailChimp authentication result: [%s]", *result)
		}
	}
}

// go test -v -run TestAWSAuthFalse /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestAWSAuthFalse(t *testing.T) {
	fmt.Println()
	AWSAccessKey := "AKIAABCDABCDABCDRFB"
	src := "AKIAQGDRKHTKEHWDVRFB   HZjvc32t8fOFrYYD2RyGFUlPXeZHHZne+u0K/Waa"

	for _, sm := range Matchers() {
		if sm.Type() == "aws_access_key_id" {
			result, err := sm.Authenticate(AWSAccessKey, src)
			if err != nil {
				t.Fatal(err)
			}
			log.Printf("TestAWSAuthTrue authentication result: [%s]", *result)
		}
	}
}
