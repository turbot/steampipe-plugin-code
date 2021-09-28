package secrets

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func init() {
	RegisterMatcher(&stripeApiKey{})
}

type stripeApiKey struct{}

func (*stripeApiKey) Type() string {
	return "stripe_api_key"
}

func (*stripeApiKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)(?:r|s)k_live_[0-9a-zA-Z]{24}`),
	}
}

func (*stripeApiKey) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	test_url := "https://api.stripe.com/v1/charges"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", test_url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(secret))))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return AUTHENTICATED, nil
	}

	// Restricted keys may be limited to certain endpoints
	if strings.HasPrefix(secret, "rk_live") {
		return NOT_IMPLEMENTED, nil
	}

	return UNAUTHENTICATED, nil
}
