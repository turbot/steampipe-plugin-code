package code

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-code/secrets"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCodeSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "code_secret",
		Description: "TODO",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("src"),
			Hydrate:    listSecret,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "secret_type", Type: proto.ColumnType_STRING, Description: "Secret type."},
			{Name: "secret", Type: proto.ColumnType_STRING, Description: "Secret string."},
			{Name: "start_offset", Type: proto.ColumnType_INT, Description: "Offset of the first character of the secret string."},
			{Name: "end_offset", Type: proto.ColumnType_INT, Description: "Offset of the last character of the secret string."},
			{Name: "line", Type: proto.ColumnType_INT, Description: "Line number of the first character of the secret string."},
			{Name: "col", Type: proto.ColumnType_INT, Description: "Column on the line of the first character of the secret string."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Hydrate: getVerified, Transform: transform.FromValue(), Description: "True if the secret has been verified as active."},
			// Other columns
			{Name: "src", Type: proto.ColumnType_STRING, Transform: transform.FromQual("src"), Description: "The source code to scan."},
		},
	}
}

type secretMatch struct {
	Matcher     secrets.SecretMatcher `json:"matcher"`
	SecretType  string                `json:"secret_type"`
	Secret      string                `json:"text"`
	StartOffset int                   `json:"start_offset"`
	EndOffset   int                   `json:"end_offset"`
	Line        int                   `json:"line"`
	Col         int                   `json:"col"`
	Verified    *bool                 `json:"verified"`
}

func listSecret(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	src := quals["src"].GetStringValue()
	for _, sm := range secrets.Matchers() {
		for _, re := range sm.DenyList() {
			matchGroups := re.FindAllStringSubmatchIndex(src, -1)
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
				secret := src[startOffset:endOffset]
				prefix := src[0:startOffset]
				line := strings.Count(prefix, "\n") + 1
				col := startOffset - strings.LastIndex(prefix, "\n")
				d.StreamListItem(ctx, secretMatch{
					Matcher:     sm,
					SecretType:  sm.Type(),
					Secret:      secret,
					StartOffset: startOffset,
					EndOffset:   endOffset,
					Line:        line,
					Col:         col,
				})
			}
		}
	}
	return nil, nil
}

func getVerified(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	sm := h.Item.(secretMatch)
	verified, err := sm.Matcher.Verify(sm.Secret)
	if err != nil {
		return nil, err
	}
	return verified, nil
}
