package secrets

import "regexp"

func init() {
	RegisterMatcher(&githubRefreshToken{})
}

//// Github Refresh Token

type githubRefreshToken struct{}

func (*githubRefreshToken) Type() string {
	return "github_refresh_token"
}

func (*githubRefreshToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)ghr_[0-9a-zA-Z]{76}`),
	}
}

func (*githubRefreshToken) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
