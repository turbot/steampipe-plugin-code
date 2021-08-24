package secrets

import "regexp"

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/turbot/go-kit/types"
// )

func init() {
	RegisterMatcher(&stripeAccessKey{})
}

type stripeAccessKey struct{}

func (*stripeAccessKey) Type() string {
	return "stripe_access_key"
}

func (*stripeAccessKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)(?:r|s)k_live_[0-9a-zA-Z]{24}`),
	}
}

func (*stripeAccessKey) Verify(secret string) (VerifiedValue, error) {
	return UNVERIFIED, nil
}

// func (*stripeAccessKey) Verify(secret string) (*bool, error) {
// 	verify_url := "https://slack.com/api/auth.test"

// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", verify_url, nil)
// 	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(secret))))

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode == http.StatusOK {
// 		return types.Bool(true), nil
// 	}
// 	if strings.HasPrefix(secret, "rk_live") {
// 		return nil, nil
// 	}

// 	return types.Bool(false), nil
// }
