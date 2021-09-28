package secrets

import (
	"regexp"

	"github.com/turbot/go-kit/types"
)

var matcherList []SecretMatcher

type SecretMatcher interface {
	Type() string
	DenyList() []*regexp.Regexp
	Authenticate(secret string, src string) (AuthenticatedResult, error)
}

func RegisterMatcher(sm SecretMatcher) {
	matcherList = append(matcherList, sm)
}

func Matchers() []SecretMatcher {
	return matcherList
}

type AuthenticatedResult *string

var (
	// Tested and creds are working
	AUTHENTICATED AuthenticatedResult = types.String("authenticated")
	// Tested and creds are not working
	UNAUTHENTICATED AuthenticatedResult = types.String("unauthenticated")
	// Not tested
	NOT_IMPLEMENTED AuthenticatedResult = types.String("not_implemented")
	// Tested but inconclusive
	UNKNOWN AuthenticatedResult = types.String("unknown")
)
