package secrets

import (
	"fmt"
	"testing"
)

func TestStripVerify(t *testing.T) {
	// go test -v -run TestStripVerify /Users/lalitbhardwaj/Documents/Learning_Code/Go_Basics/azure_ad_mine
	stripeApiKey := "abc"
	// fmt.Println("String Outtput: \n", string(body))

	for _, sm := range Matchers() {
		if sm.Type() == "stripe_api_key" {
			result, err := sm.Verify(stripeApiKey)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("Verification result %v", *result)
		}

	}
}
