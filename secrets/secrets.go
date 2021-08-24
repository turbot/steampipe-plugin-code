package secrets

import (
	"regexp"

	"github.com/turbot/go-kit/types"
)

var matcherList []SecretMatcher

type SecretMatcher interface {
	Type() string
	DenyList() []*regexp.Regexp
	// True if secret works, false if it does not and nil if there is no test.
	Verify(string) (VerifiedValue, error)
}

func RegisterMatcher(sm SecretMatcher) {
	matcherList = append(matcherList, sm)
}

func Matchers() []SecretMatcher {
	return matcherList
}

// VERIFIED enumerates the values for verified value for secret.
type VerifiedValue *string

var (
	// when tested and creds do not work
	VERIFIEDFALSE VerifiedValue = types.String("verified false")
	// when tested and creds do work
	VERIFIEDTRUE VerifiedValue = types.String("verified true")
	// when not tested
	UNVERIFIED VerifiedValue = types.String("unverified")
)
