package secrets

import "regexp"

func init() {
	RegisterMatcher(&facebookOauth{})
}

// https://github.com/l4yton/RegHex#facebook-oauth
type facebookOauth struct{}

func (*facebookOauth) Type() string {
	return "facebook_oauth"
}

func (*facebookOauth) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K].*['|\"][0-9a-f]{32}['|\"]`),
	}
}

func (*facebookOauth) Verify(secret string) (VerifiedResult, error) {
	return UNVERIFIED, nil
}
