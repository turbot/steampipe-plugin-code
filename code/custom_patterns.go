package code

import (
	"regexp"

	"github.com/turbot/steampipe-plugin-code/secrets"
)

type customPatterns struct{}

func (*customPatterns) Type() string {
	return "custom_pattern"
}

func (*customPatterns) DenyList() []*regexp.Regexp {
	return customRegexList
}

func (*customPatterns) Verify(secret string) (secrets.VerifiedResult, error) {
	return secrets.UNVERIFIED, nil
}
