package secrets

import "regexp"

func init() {
	RegisterMatcher(&twilioAccessKey{})
}

type twilioAccessKey struct{}

func (*twilioAccessKey) Type() string {
	return "twilio_access_key"
}

func (*twilioAccessKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`AC[a-z0-9]{32}`), // Account SID (ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx)
		regexp.MustCompile(`SK[a-z0-9]{32}`), // Auth token (SKxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx)
	}
}

func (*twilioAccessKey) Verify(secret string) (*bool, error) {
	return nil, nil
}
