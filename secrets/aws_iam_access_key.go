package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&awsAccessKeyID{})
}

type awsAccessKeyID struct{}

func (*awsAccessKeyID) Type() string {
	return "aws_access_key_id"
}

// https://aws.amazon.com/blogs/security/a-safer-way-to-distribute-aws-credentials-to-ec2/''

// grep -RP '(?<![A-Z0-9])[A-Z0-9]{20}(?![A-Z0-9])' *
// grep -RP '(?<![A-Za-z0-9/+=])[A-Za-z0-9/+=]{40}(?![A-Za-z0-9/+=])' *

func (*awsAccessKeyID) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("(?m)AKIA[0-9A-Z]{16}"),
	}
}

func (*awsAccessKeyID) Verify(secret string) (VerifiedValue, error) {
	return UNVERIFIED, nil
}
