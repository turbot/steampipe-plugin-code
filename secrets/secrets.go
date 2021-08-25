package secrets

import (
	"regexp"

	"github.com/turbot/go-kit/types"
)

var matcherList []SecretMatcher

type SecretMatcher interface {
	Type() string
	DenyList() []*regexp.Regexp
	// VERIFIED_TRUE if secret works, VERIFIED_FALSE if it does not and UNVERIFIED if there is no test.
	Verify(string) (VerifiedResult, error)
}

func RegisterMatcher(sm SecretMatcher) {
	matcherList = append(matcherList, sm)
}

func Matchers() []SecretMatcher {
	return matcherList
}

// VERIFIED enumerates the values for verified value for secret.
type VerifiedResult *string

var (
	// when tested and creds do not work
	VERIFIED_FALSE VerifiedResult = types.String("verified false")
	// when tested and creds do work
	VERIFIED_TRUE VerifiedResult = types.String("verified true")
	// when not tested
	UNVERIFIED VerifiedResult = types.String("unverified")
)
