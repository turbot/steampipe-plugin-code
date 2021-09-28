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

func (*githubRefreshToken) Verify(secret string, src string) (VerifiedResult, error) {
	return UNVERIFIED, nil
}
