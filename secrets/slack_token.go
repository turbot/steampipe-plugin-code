package secrets

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func init() {
	RegisterMatcher(&slackToken{})
}

type slackToken struct{}

func (*slackToken) Type() string {
	return "slack_token"
}

func (*slackToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)xox(?:a|b|p|o|s|r)-(?:\d+-)+[a-z0-9]+`),
	}
}

func (*slackToken) Verify(secret string) (*bool, error) {
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
	result := res["ok"].(bool)
	return &result, nil
}
