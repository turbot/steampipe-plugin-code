package secrets

import "regexp"

func init() {
	RegisterMatcher(&oktaToken{})
}

type oktaToken struct{}

func (*oktaToken) Type() string {
	return "okta_token"
}

func (*oktaToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)00[a-zA-Z0-9\-\_]{40}`), // https://devforum.okta.com/t/api-token-length/5519
	}
}

func (*oktaToken) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
