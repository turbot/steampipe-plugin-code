package secrets

import (
	"fmt"
	"testing"
)

func TestSlackTokenVerify(t *testing.T) {
	// go test -v -run TestSlackTokenVerify /Users/lalitbhardwaj/Documents/Learning_Code/Go_Basics/azure_ad_mine
	slack_api_token := "xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65"

	for _, sm := range Matchers() {
		if sm.Type() == "slack_api_token" {
			result, err := sm.Verify(slack_api_token)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("Verification result: %v\n", *result)
		}

	}
}
