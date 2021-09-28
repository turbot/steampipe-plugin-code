package secrets

import "regexp"

func init() {
	RegisterMatcher(&githubAppToken{})
}

type githubAppToken struct{}

func (*githubAppToken) Type() string {
	return "github_app_token"
}

func (*githubAppToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)(ghu|ghs)_[0-9a-zA-Z]{36}`),
	}
}

func (*githubAppToken) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
