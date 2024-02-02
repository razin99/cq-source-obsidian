package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-obsidian/client"
	"github.com/razin99/cq-source-obsidian/obsidian"
)

func AllEventsTable() *schema.Table {
	return &schema.Table{
		Name:     "obsidian_all_events",
		Resolver: fetchAllEvents,
		Transform: transformers.TransformWithStruct(
			&obsidian.Event{},
			transformers.WithPrimaryKeys("Id"),
		),
	}
}

func fetchAllEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := ""

	for {
		response, err := cl.Obsidian.GetAllEvents(ctx, cursor)
		if err != nil {
			return err
		}

		for _, event := range response.Data.GetEvents.Results {
			res <- event
		}

		cursor = response.Data.GetEvents.Cursor

		if !response.Data.GetEvents.HasMoreResults || response.Data.GetEvents.Cursor == "" {
			break
		}
	}

	return nil
}
