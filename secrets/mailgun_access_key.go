package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&mailgunAccessKey{})
}

type mailgunAccessKey struct{}

func (*mailgunAccessKey) Type() string {
	return "mailgun_access_key"
}

func (*mailgunAccessKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("(?m)key-[0-9a-zA-Z]{32}"),
	}
}

func (*mailgunAccessKey) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	return NOT_IMPLEMENTED, nil
}

// select * from code_secret where src = 'key-3ax6xnjp29jd6fds4gc373sgvjxteol0'