package secrets

import "regexp"

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/turbot/go-kit/types"
// )

func init() {
	RegisterMatcher(&twitterSecretKey{})
}

type twitterSecretKey struct{}

func (*twitterSecretKey) Type() string {
	return "twitter_secret_key"
}

func (*twitterSecretKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)twitter(.{0,20})?['\"][0-9a-z]{35,44}`),
	}
}

func (*twitterSecretKey) Verify(secret string) (*bool, error) {
	return nil, nil
}
