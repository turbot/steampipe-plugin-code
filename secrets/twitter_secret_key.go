package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&twitterSecretKey{})
}

type twitterSecretKey struct{}

func (*twitterSecretKey) Type() string {
	return "twitter_secret_key"
}

func (*twitterSecretKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?im)twitter(.{0,20})?['\"][0-9a-z]{35,44}`),
	}
}

func (*twitterSecretKey) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
