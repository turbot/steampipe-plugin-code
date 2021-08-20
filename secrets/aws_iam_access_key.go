package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&awsIAMAccessKey{})
}

type awsIAMAccessKey struct{}

func (*awsIAMAccessKey) Type() string {
	return "aws_iam_access_key"
}

func (*awsIAMAccessKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("(?m)AKIA[0-9A-Z]{16}"),
	}
}

func (*awsIAMAccessKey) Verify(secret string) (*bool, error) {
	return nil, nil
}
