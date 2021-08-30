package code

import (
	"context"
	"path/filepath"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {

	tables := map[string]*plugin.Table{
		"code_secret": tableCodeSecret(ctx),
	}
	paths := []string{"/Users/nathan/src/steampipe-plugin-code/test.csv", "/Users/nathan/src/steampipe-plugin-code/test2.csv", "/Users/nathan/Downloads/downloaded-logs.csv"}
	for _, p := range paths {
		tableCtx := context.WithValue(ctx, "path", p)
		base := filepath.Base(p)
		tables[base[0:len(base)-len(filepath.Ext(base))]] = tableCodeCSV(tableCtx)
	}

	p := &plugin.Plugin{
		Name:             "steampipe-plugin-code",
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap:         tables,
	}
	return p
}
