package secrets

import (
	"fmt"
	"log"
	"testing"
)

// Run all test
// go test -v -run $*^  /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets

// go test -v -run TestSlackTokenVerify /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestSlackTokenVerify(t *testing.T) {
	fmt.Println()
	slack_api_token := "xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65"

	for _, sm := range Matchers() {
		if sm.Type() == "slack_api_token" {
			result, err := sm.Verify(slack_api_token, slack_api_token)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestSlackTokenVerify verification result: [%s]", *result)
		}

	}
}

// go test -v -run TestStripeVerifyOk /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestStripeVerifyOk(t *testing.T) {
	fmt.Println()
	stripeApiKey := "sk_test_tR3PYbcVNZZ796tH88S4VQ2u"

	for _, sm := range Matchers() {
		if sm.Type() == "stripe_api_key" {
			result, err := sm.Verify(stripeApiKey, stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestStripeVerifyOk verification result: [%s]", *result)
		}

	}
}

// go test -v -run TestStripeVerifyFail /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestStripeVerifyFail(t *testing.T) {
	fmt.Println()
	stripeApiKey := "sk_live_tR3PYbcVNZZ796tH88S4VQ2u"

	for _, sm := range Matchers() {
		if sm.Type() == "stripe_api_key" {
			result, err := sm.Verify(stripeApiKey, stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestStripeVerifyFail verification result: [%s]", *result)
		}
	}
}

// go test -v -run TestMailChimpVerfifiedFalse /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestMailChimpVerfifiedFalse(t *testing.T) {
	fmt.Println()
	mailchimpAccessKey := "a11b1d1baf01fd5556666f434g9b123a-us5"

	for _, sm := range Matchers() {
		if sm.Type() == "mailchimp_access_key" {
			result, err := sm.Verify(mailchimpAccessKey, mailchimpAccessKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestMailChimp verification result: [%s]", *result)
		}
	}
}

// go test -v -run TestAWSVerfifiedFalse /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestAWSVerfifiedFalse(t *testing.T) {
	fmt.Println()
	AWSAccessKey := "AKIAABCDABCDABCDRFB"
	src := "AKIAQGDRKHTKEHWDVRFB   HZjvc32t8fOFrYYD2RyGFUlPXeZHHZne+u0K/Waa"

	for _, sm := range Matchers() {
		if sm.Type() == "aws_access_key_id" {
			result, err := sm.Verify(AWSAccessKey, src)
			if err != nil {
				t.Fatal(err)
			}
			log.Printf("TestAWSVerfifiedTrue verification result: [%s]", *result)
		}
	}
}
