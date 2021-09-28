package secrets

import (
	"fmt"
	"regexp"
)

func init() {
	RegisterMatcher(&basicAuth{})
}

type basicAuth struct{}

func (*basicAuth) Type() string {
	return "basic_auth"
}

// These characters are derived from RFC 3986 Section 2.2.
//
// We don't expect any of these delimiter characters to appear in
// the username/password component of the URL, seeing that this would probably
// result in an unexpected URL parsing (and probably won't even work).

func basicAuthReservedCharacters() string {
	return `:/?#[]@`
}

func basicAuthSubDelimiterCharacters() string {
	return `!$&\'()*+,;=`
}

func basicAuthRegexString() string {
	allChars := basicAuthReservedCharacters() + basicAuthSubDelimiterCharacters()
	return fmt.Sprintf(`(?m)([a-zA-Z0-9+-\.]+://[^%s\s]+:[^%s\s]+)@`, regexp.QuoteMeta(allChars), regexp.QuoteMeta(allChars))
}

func (*basicAuth) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(basicAuthRegexString()),
	}
}

func (*basicAuth) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	// Not supported
	return NOT_IMPLEMENTED, nil
}
