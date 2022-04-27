package code

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-code/secrets"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableCodeSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "code_secret",
		Description: "Detect, and verify if possible, secrets in a given source string.",
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
			{Name: "authenticated", Type: proto.ColumnType_STRING, Hydrate: getAuthenticated, Transform: transform.FromValue(), Description: "Authentication status of the secret. Valid values are \"authenticated\", \"unauthenticated\", \"not_implemented\", and \"unknown\"."},
			// Other columns
			{Name: "src", Type: proto.ColumnType_STRING, Transform: transform.FromQual("src"), Description: "The source code to scan."},
		},
	}
}

type secretMatch struct {
	Matcher       secrets.SecretMatcher `json:"matcher"`
	SecretType    string                `json:"secret_type"`
	Secret        string                `json:"text"`
	StartOffset   int                   `json:"start_offset"`
	EndOffset     int                   `json:"end_offset"`
	Line          int                   `json:"line"`
	Col           int                   `json:"col"`
	Authenticated string                `json:"authenticated"`
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

func getAuthenticated(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	sm := h.Item.(secretMatch)
	quals := d.KeyColumnQuals
	src := quals["src"].GetStringValue()
	authenticated, err := sm.Matcher.Authenticate(sm.Secret, src)
	if err != nil {
		return nil, err
	}
	return authenticated, nil
}
