// Code generated by codegen; DO NOT EDIT.

package {{.SubService}}

import (
  "github.com/cloudquery/plugin-sdk/schema"
)

func {{.SubService | ToCamel}}() *schema.Table {
  return &schema.Table{{template "table.go.tpl" .Table}}
}