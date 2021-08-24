package secrets

import "regexp"

func init() {
	RegisterMatcher(&twilioAuthToken{})
}

type twilioAuthToken struct{}

func (*twilioAuthToken) Type() string {
	return "twilio_auth_token"
}

func (*twilioAuthToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)SK[a-z0-9]{32}`),
	}
}

func (*twilioAuthToken) Verify(secret string) (VerifiedValue, error) {
	return UNVERIFIED, nil
}
