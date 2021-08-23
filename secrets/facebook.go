package secrets

import "regexp"

func init() {
	RegisterMatcher(&facebookAccessToken{})
	RegisterMatcher(&facebookOauth{})
	RegisterMatcher(&facebookSecretKey{})
}

// https://github.com/l4yton/RegHex#facebook-access-token
type facebookAccessToken struct{}

func (*facebookAccessToken) Type() string {
	return "facebook_access_token"
}

func (*facebookAccessToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`EAACEdEose0cBA[0-9A-Za-z]+`),
	}
}

func (*facebookAccessToken) Verify(secret string) (*bool, error) {
	return nil, nil
}

// https://github.com/l4yton/RegHex#facebook-oauth
type facebookOauth struct{}

func (*facebookOauth) Type() string {
	return "facebook_oauth"
}

func (*facebookOauth) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K].*['|\"][0-9a-f]{32}['|\"]`),
	}
}

func (*facebookOauth) Verify(secret string) (*bool, error) {
	return nil, nil
}

// https://github.com/l4yton/RegHex#facebook-secret-key
type facebookSecretKey struct{}

func (*facebookSecretKey) Type() string {
	return "facebook_secret_key"
}

func (*facebookSecretKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)(facebook|fb)(.{0,20})?(?-i)['\"][0-9a-f]{32}`),
	}
}

func (*facebookSecretKey) Verify(secret string) (*bool, error) {
	return nil, nil
}
