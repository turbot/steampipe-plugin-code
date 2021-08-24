package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&azureStorageAccountKey{})
}

type azureStorageAccountKey struct{}

func (*azureStorageAccountKey) Type() string {
	return "azure_storage_account_key"
}

func (*azureStorageAccountKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`AccountKey=[a-zA-Z0-9+/=]{88}`),
		regexp.MustCompile(`[a-zA-Z0-9+/=]{88}`),
	}
}

func (*azureStorageAccountKey) Verify(secret string) (VerifiedResult, error) {
	return UNVERIFIED, nil
}
