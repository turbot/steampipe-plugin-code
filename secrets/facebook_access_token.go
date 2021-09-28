package secrets

import "regexp"

func init() {
	RegisterMatcher(&facebookAccessToken{})
}

// https://github.com/zricethezav/gitleaks/blob/6f5ad9dc0b385c872f652324188ce91da7157c7c/config/default.go
// https://github.com/l4yton/RegHex#facebook-access-token
type facebookAccessToken struct{}

func (*facebookAccessToken) Type() string {
	return "facebook_access_token"
}

func (*facebookAccessToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)EAACEdEose0cBA[0-9A-Za-z]+`),
	}
}

func (*facebookAccessToken) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
