// Code generated by codegen; DO NOT EDIT.

package rulesets

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Rulesets() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_rulesets",
		Description: `https://developer.pagerduty.com/api-reference/633f1ecb6c03b-list-rulesets`,
		Resolver:    fetchRulesets,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "self",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Self"),
			},
			{
				Name:     "routing_keys",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RoutingKeys"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "creator",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Creator"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "updater",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Updater"),
			},
			{
				Name:     "team",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Team"),
			},
		},

		Relations: []*schema.Table{
			RulesetRules(),
		},
	}
}
