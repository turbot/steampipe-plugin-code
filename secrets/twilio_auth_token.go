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
		regexp.MustCompile(`(?m)AC[a-z0-9]{32}`),
		regexp.MustCompile(`(?m)SK[a-z0-9]{32}`),
	}
}

func (*twilioAuthToken) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}
