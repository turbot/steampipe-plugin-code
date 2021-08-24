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

// func (*mailchimpAccessKey) Verify(secret string) (VerifiedResult, error) {
// 	return UNVERIFIED, nil
// }

func (*mailchimpAccessKey) Verify(secret string) (VerifiedResult, error) {
	datacenter_number := strings.Split(secret, "-us")
	verify_url := fmt.Sprintf("https://us%s.api.mailchimp.com/3.0/", datacenter_number)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", verify_url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("any_user:"+secret))))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return VERIFIED_TRUE, nil
	}

	return VERIFIED_FALSE, nil
}
