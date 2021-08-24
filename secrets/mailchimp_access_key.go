package secrets

import (
	"regexp"
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

func (*mailchimpAccessKey) Verify(secret string) (VerifiedValue, error) {
	return UNVERIFIED, nil
}
