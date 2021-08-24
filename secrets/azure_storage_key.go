package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&azureStorageKey{})
}

type azureStorageKey struct{}

func (*azureStorageKey) Type() string {
	return "azure_storage_account_access_key"
}

func (*azureStorageKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`AccountKey=[a-zA-Z0-9+/=]{88}`),
		regexp.MustCompile(`[a-zA-Z0-9+/=]{88}`),
	}
}

func (*azureStorageKey) Verify(secret string) (VerifiedValue, error) {
	return UNVERIFIED, nil
}
