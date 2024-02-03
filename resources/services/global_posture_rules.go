package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-obsidian/client"
	"github.com/razin99/cq-source-obsidian/obsidian"
)

func GlobalPostureRulesTable() *schema.Table {
	return &schema.Table{
		Name:     "obsidian_global_posture_rules",
		Resolver: fetchGlobalPostureRules,
		Transform: transformers.TransformWithStruct(
			&obsidian.PostureRule{},
			transformers.WithPrimaryKeys("RuleId"),
		),
	}
}

func fetchGlobalPostureRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := ""

	for {
		response, err := cl.Obsidian.ListGlobalPostureRules(ctx, cursor)
		if err != nil {
			return err
		}

		for _, rule := range response.Data.ListGlobalPostureRules.Rules {
			res <- rule
		}

		cursor = response.Data.ListGlobalPostureRules.Cursor

		if !response.Data.ListGlobalPostureRules.HasMoreResults || response.Data.ListGlobalPostureRules.Cursor == "" {
			break
		}
	}

	return nil
}
