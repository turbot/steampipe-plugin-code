package secrets

import "regexp"

func init() {
	RegisterMatcher(&googleApiKey{})
}

type googleApiKey struct{}

func (*googleApiKey) Type() string {
	return "google_api_key"
}

func (*googleApiKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)AIza[0-9A-Za-z\\-_]{35}`),
	}
}

func (*googleApiKey) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
