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
			fmt.Printf("TestSlackTokenVerify verification result: %s", *result)
		}

	}
}

// go test -v -run TestStripVerifyOk /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestStripVerifyOk(t *testing.T) {
	fmt.Println()
	stripeApiKey := "sk_test_tR3PYbcVNZZ796tH88S4VQ2u"

	for _, sm := range Matchers() {
		if sm.Type() == "stripe_api_key" {
			result, err := sm.Verify(stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestStripVerifyOk verification result: %s", *result)
		}

	}
}

// go test -v -run TestStripVerifyFail /Users/lalitbhardwaj/Turbot/prod/steampipe-plugin-code/secrets
func TestStripVerifyFail(t *testing.T) {
	fmt.Println()
	stripeApiKey := "sk_test_tR3PYbcVNZZ796tH88S4VQ2b"

	for _, sm := range Matchers() {
		if sm.Type() == "stripe_api_key" {
			result, err := sm.Verify(stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("TestStripVerifyFail verification result: %s", *result)
		}
	}
}
