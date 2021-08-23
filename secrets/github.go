package secrets

import "regexp"

func init() {
	RegisterMatcher(&githubToken{})
}

type githubToken struct{}

func (*githubToken) Type() string {
	return "github_token"
}

func (*githubToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		// https://github.blog/2021-04-05-behind-githubs-new-authentication-token-formats/
		regexp.MustCompile(`(ghp|gho|ghu|ghs|ghr)_[A-Za-z0-9_]{36}`),
		regexp.MustCompile(`(?i)github(.{0,20})?(?-i)['\"][0-9a-zA-Z]{35,40}`),
		regexp.MustCompile(`[0-9a-f]{40}`), // https://bl.ocks.org/magnetikonline/073afe7909ffdd6f10ef06a00bc3bc88
	}
}

func (*githubToken) Verify(secret string) (*bool, error) {
	return nil, nil
}

// select * from code_secret where src = '45ab6f911111f9f376a5b52c25d22113f2b45fa1'
