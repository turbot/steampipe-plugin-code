package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&jwt{})
}

type jwt struct{}

func (*jwt) Type() string {
	return "jwt"
}

func (*jwt) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)eyJ[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*`),
	}
}

func (*jwt) Verify(secret string) (VerifiedValue, error) {
	return UNVERIFIED, nil
}
