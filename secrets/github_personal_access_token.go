package secrets

import "regexp"

func init() {
	RegisterMatcher(&githubPersonalAccessToken{})
}

type githubPersonalAccessToken struct{}

func (*githubPersonalAccessToken) Type() string {
	return "github_personal_access_token"
}

func (*githubPersonalAccessToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		// https://github.blog/2021-04-05-behind-githubs-new-authentication-token-formats/
		regexp.MustCompile(`(?m)(ghp|gho|ghu|ghs|ghr)_[A-Za-z0-9_]{36}`),
		regexp.MustCompile(`(?m)[0-9a-f]{40}`), // https://bl.ocks.org/magnetikonline/073afe7909ffdd6f10ef06a00bc3bc88
	}
}

func (*githubPersonalAccessToken) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
