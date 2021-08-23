package secrets

import "regexp"

func init() {
	RegisterMatcher(&githubPersonalAccessToken{})
	RegisterMatcher(&githubOAuthAccessToken{})
	RegisterMatcher(&githubAppToken{})
	RegisterMatcher(&githubRefreshToken{})
}

//// Github Personal Access Token

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

func (*githubPersonalAccessToken) Verify(secret string) (*bool, error) {
	return nil, nil
}

//// Github OAuth Access Token

type githubOAuthAccessToken struct{}

func (*githubOAuthAccessToken) Type() string {
	return "github_oauth_access_token"
}

func (*githubOAuthAccessToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)gho_[0-9a-zA-Z]{36}`),
	}
}

func (*githubOAuthAccessToken) Verify(secret string) (*bool, error) {
	return nil, nil
}

//// Github App Token

type githubAppToken struct{}

func (*githubAppToken) Type() string {
	return "github_app_token"
}

func (*githubAppToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)(ghu|ghs)_[0-9a-zA-Z]{36}`),
	}
}

func (*githubAppToken) Verify(secret string) (*bool, error) {
	return nil, nil
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

func (*githubRefreshToken) Verify(secret string) (*bool, error) {
	return nil, nil
}

// select * from code_secret where src = '45ab6f911111f9f376a5b52c25d22113f2b45fa1'
