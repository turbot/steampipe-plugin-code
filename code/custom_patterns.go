package code

import (
	"regexp"

	"github.com/turbot/steampipe-plugin-code/secrets"
)

type customPattern struct{}

func (*customPattern) Type() string {
	return "custom_pattern"
}

func (*customPattern) DenyList() []*regexp.Regexp {
	return customRegexList
}

func (*customPattern) Verify(secret string) (secrets.VerifiedResult, error) {
	return secrets.UNVERIFIED, nil
}
