package code

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"

	"gopkg.in/ini.v1"
)

func tableParseIni(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "parse_ini",
		Description: "Table representation of an INI file.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("path"),
			Hydrate:    parseIniFIle,
		},
		Columns: []*plugin.Column{
			{Name: "path", Type: proto.ColumnType_STRING, Description: "Specifies the path of the ini file."},
			{Name: "section", Type: proto.ColumnType_STRING, Description: "Specifies the name of the section."},
			{Name: "key", Type: proto.ColumnType_STRING, Description: "The name of the key."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "The value of corresponding key."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "nested_values", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

type parseFormat struct {
	Path         string
	Section      string
	Key          string
	Value        string
	Comment      string
	NestedValues []string
}

func parseIniFIle(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	givenPath := d.KeyColumnQuals["path"].GetStringValue()

	var opts ini.LoadOptions
	opts.AllowNestedValues = true
	opts.AllowPythonMultilineValues = true

	cfg, err := ini.LoadSources(opts, givenPath)
	if err != nil {
		panic(fmt.Errorf("fail to read file: %v", err))
	}

	for _, i := range cfg.Sections() {
		// Extract keys of a section
		for _, key := range cfg.Section(i.Name()).Keys() {
			d.StreamListItem(ctx, parseFormat{
				Path:         givenPath,
				Section:      i.Name(),
				Key:          key.Name(),
				Value:        parseValue(cfg, key.String()),
				Comment:      key.Comment,
				NestedValues: key.NestedValues(),
			})
		}
	}
	return nil, nil
}

// parseValue will parse env variable and other variable references with its actual value
func parseValue(cfg *ini.File, str string) string {
	// Check for value of the environment variable references
	isEnvVar, _ := regexp.MatchString(".*\\${.*}.*", str)
	if isEnvVar {
		regexExp := regexp.MustCompile(`\$\{(.*?)\}`)
		matchedStr := regexExp.FindStringSubmatch(str)
		if len(matchedStr) > 1 {
			// Check for reference from other section, i.e. path = ${Common.system_dir}/Library/Frameworks/
			if strings.Contains(matchedStr[1], ".") {
				splitStr := strings.Split(matchedStr[1], ".")
				sec := strings.Join(splitStr[:len(splitStr)-1], ".")
				key := splitStr[len(splitStr)-1]
				value := cfg.Section(sec).Key(key).String()
				str = strings.Replace(str, matchedStr[0], value, -1)
			} else {
				// Replace the matched string with env variable value
				str = strings.Replace(str, matchedStr[0], os.Getenv(matchedStr[1]), -1)
			}
		}
	}
	return str
}
