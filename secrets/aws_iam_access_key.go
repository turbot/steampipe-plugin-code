package secrets

import (
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/turbot/go-kit/helpers"
)

func init() {
	RegisterMatcher(&awsAccessKeyID{})
}

type awsAccessKeyID struct{}

func (*awsAccessKeyID) Type() string {
	return "aws_access_key_id"
}

// https://aws.amazon.com/blogs/security/a-safer-way-to-distribute-aws-credentials-to-ec2/

// grep -RP '(?<![A-Z0-9])[A-Z0-9]{20}(?![A-Z0-9])' *
// grep -RP '(?<![A-Za-z0-9/+=])[A-Za-z0-9/+=]{40}(?![A-Za-z0-9/+=])' *

func (*awsAccessKeyID) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("(?m)AKIA[0-9A-Z]{16}"),
	}
}

func (*awsAccessKeyID) Verify(secret string, src string) (VerifiedResult, error) {
	//  This examines the variable name to identify AWS secret tokens.
	//  The order is important since we want to prefer finding `AKIA`-based
	//  keys (since they can be verified), rather than the secret tokens.
	// re := regexp.MustCompile("(?m)aws.{0,20}?['\"]([0-9a-zA-Z/+]{40})['\"]")
	re := regexp.MustCompile(`(?m)([0-9a-zA-Z\/+]{40})`)

	matchGroups := re.FindAllStringSubmatchIndex(src, -1)
	var secrets = make([]string, 0)
	for _, m := range matchGroups {
		var startOffset, endOffset int
		if len(m) > 2 {
			// If the regexp in the secret matcher has a match group, then use it
			// as the "secret" from the string. For example "user:(secret)".
			startOffset = m[2]
			endOffset = m[3]
		} else {
			// If the regexp has no match group, then use the full match as the secret.
			// e.g. "tok-[a-z]+"
			startOffset = m[0]
			endOffset = m[1]
		}

		secret_key := src[startOffset:endOffset]
		secrets = append(secrets, secret_key)
	}

	if len(secrets) == 0 {
		return UNVERIFIED, nil
	}

	for _, secret_key := range secrets {
		sessionOptions := session.Options{
			Config: aws.Config{
				Region:      aws.String("us-east-1"),
				MaxRetries:  aws.Int(0),
				Credentials: credentials.NewStaticCredentials(secret, secret_key, ""),
			},
		}

		sess, err := session.NewSessionWithOptions(sessionOptions)
		if err != nil {
			return nil, err
		}
		svc := sts.New(sess)

		callerIdentity, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				// Valid means keys match the regex
				// SignatureDoesNotMatch - In case Access key is invalid but secret key is not is correct
				// IncompleteSignature - When the access key is invalid
				// InvalidClientTokenId - When access key and secret key are valid but expired
				if helpers.StringSliceContains([]string{"SignatureDoesNotMatch", "IncompleteSignature", "InvalidClientTokenId"}, awsErr.Code()) {
					return VERIFIED_FALSE, nil
				}
			}
		}

		if callerIdentity != nil {
			if callerIdentity.Account != nil {
				return VERIFIED_TRUE, nil
			}
		}
	}

	return UNVERIFIED, nil
}
