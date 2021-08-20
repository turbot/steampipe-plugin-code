package secrets

import (
	"regexp"
)

var matcherList []SecretMatcher

type SecretMatcher interface {
	Type() string
	DenyList() []*regexp.Regexp
	// True if secret works, false if it does not and nil if there is no test.
	Verify(string) (*bool, error)
}

func RegisterMatcher(sm SecretMatcher) {
	matcherList = append(matcherList, sm)
}

func Matchers() []SecretMatcher {
	return matcherList
}
