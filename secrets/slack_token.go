package secrets

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func init() {
	RegisterMatcher(&slackApiToken{})
}

type slackApiToken struct{}

func (*slackApiToken) Type() string {
	return "slack_api_token"
}

func (*slackApiToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)xox(?:a|b|p|o|s|r)-(?:\d+-)+[a-z0-9A-X]+`),
	}
}

func (*slackApiToken) Authenticate(secret string, src string) (AuthenticatedResult, error) {
	resp, err := http.PostForm("https://slack.com/api/auth.test", url.Values{"token": {secret}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res["ok"].(bool) {
		return AUTHENTICATED, nil
	}
	return UNAUTHENTICATED, nil
}
