package secrets

import "regexp"

func init() {
	RegisterMatcher(&herokuApiKey{})
}

type herokuApiKey struct{}

func (*herokuApiKey) Type() string {
	return "heroku_api_key"
}

func (*herokuApiKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`),
	}
}

func (*herokuApiKey) Verify(secret string, src string) (VerifiedResult, error) {
	return UNVERIFIED, nil
}

// select * from code_secret where src = 'def66a66-3411-44a3-ad5f-a6af6f316f92';
