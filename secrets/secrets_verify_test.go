package secrets

import (
	"fmt"
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
			result, err := sm.Verify(slack_api_token)
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
			result, err := sm.Verify(stripeApiKey)
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
			result, err := sm.Verify(stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestStripeVerifyFail verification result: [%s]", *result)
		}
	}
}

// fmt.Println("String Output: \n", string(body))
// fmt.Println("resp.Status: \n", resp.Status)
// fmt.Println("resp.StatusCode: \n", resp.StatusCode)
