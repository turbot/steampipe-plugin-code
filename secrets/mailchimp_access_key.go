package secrets

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func init() {
	RegisterMatcher(&mailchimpAccessKey{})
}

type mailchimpAccessKey struct{}

func (*mailchimpAccessKey) Type() string {
	return "mailchimp_access_key"
}

func (*mailchimpAccessKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("(?m)[0-9a-z]{32}-us[0-9]{1,2}"),
	}
}

// func (*mailchimpAccessKey) Authenticate(secret string, src string) (AuthenticatedResult, error) {
// 	return NOT_IMPLEMENTED, nil
// }

func (*mailchimpAccessKey) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	datacenter_number := strings.Split(secret, "-us")
	test_url := fmt.Sprintf("https://us%s.api.mailchimp.com/3.0/", datacenter_number[len(datacenter_number)-1])

	client := &http.Client{}
	req, _ := http.NewRequest("GET", test_url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("any_user:"+secret))))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return AUTHENTICATED, nil
	}

	return UNAUTHENTICATED, nil
}
