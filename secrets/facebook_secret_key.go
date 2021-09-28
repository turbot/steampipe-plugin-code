package secrets

import "regexp"

func init() {
	RegisterMatcher(&facebookSecretKey{})
}

// https://github.com/l4yton/RegHex#facebook-secret-key
type facebookSecretKey struct{}

func (*facebookSecretKey) Type() string {
	return "facebook_secret_key"
}

func (*facebookSecretKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?im)(facebook|fb)(.{0,20})?(?-i)['\"][0-9a-f]{32}`),
	}
}

func (*facebookSecretKey) Verify(secret string) (VerifiedResult, error) {
	return UNVERIFIED, nil
}