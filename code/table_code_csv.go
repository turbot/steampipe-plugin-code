package code

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableCodeCSV(ctx context.Context) *plugin.Table {

	path := ctx.Value("path").(string)
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	plugin.Logger(ctx).Warn("tableCodeCSV", "path", path)

	r := csv.NewReader(csvFile)

	r.Comma = ','
	r.Comment = '#'
	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	cols := []*plugin.Column{}
	for _, i := range header {
		cols = append(cols, &plugin.Column{Name: i, Type: proto.ColumnType_STRING, Transform: transform.FromField(i)})
	}

	return &plugin.Table{
		Name:        path,
		Description: "TODO",
		List: &plugin.ListConfig{
			Hydrate: listCSVWithPath(path),
		},
		Columns: cols,
	}
}

func listCSVWithPath(path string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

		csvFile, err := os.Open(path)
		//csvFile, err := os.Open("/Users/nathan/src/steampipe-plugin-code/test.csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}

		r := csv.NewReader(csvFile)

		r.Comma = ','
		r.Comment = '#'

		records, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range records[1:] {
			row := map[string]string{}
			for idx, j := range i {
				plugin.Logger(ctx).Warn("listCSV", "idx", idx)
				plugin.Logger(ctx).Warn("listCSV", "records[0][idx]", records[0][idx])
				plugin.Logger(ctx).Warn("listCSV", "j", j)
				row[records[0][idx]] = j
			}
			d.StreamListItem(ctx, row)
		}

		return nil, nil
	}
}
