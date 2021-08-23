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

func (*oktaToken) Verify(secret string) (*bool, error) {
	return nil, nil
}

// select * from code_secret where src = '02d0YZgNSJwlNew6lZG-6qGThisisatest-token'
