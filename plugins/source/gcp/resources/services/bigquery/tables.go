// Code generated by codegen; DO NOT EDIT.

package bigquery

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigquery_tables",
		Resolver:            fetchTables,
		PreResourceResolver: tableGet,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "clone_definition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloneDefinition"),
			},
			{
				Name:     "clustering",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Clustering"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "default_collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCollation"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "encryption_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionConfiguration"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "expiration_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ExpirationTime"),
			},
			{
				Name:     "external_data_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExternalDataConfiguration"),
			},
			{
				Name:     "friendly_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FriendlyName"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "materialized_view",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaterializedView"),
			},
			{
				Name:     "max_staleness",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MaxStaleness"),
			},
			{
				Name:     "model",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Model"),
			},
			{
				Name:     "num_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumBytes"),
			},
			{
				Name:     "num_long_term_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumLongTermBytes"),
			},
			{
				Name:     "num_physical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumPhysicalBytes"),
			},
			{
				Name:     "num_rows",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumRows"),
			},
			{
				Name:     "num_active_logical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumActiveLogicalBytes"),
			},
			{
				Name:     "num_active_physical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumActivePhysicalBytes"),
			},
			{
				Name:     "num_long_term_logical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumLongTermLogicalBytes"),
			},
			{
				Name:     "num_long_term_physical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumLongTermPhysicalBytes"),
			},
			{
				Name:     "num_partitions",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumPartitions"),
			},
			{
				Name:     "num_time_travel_physical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumTimeTravelPhysicalBytes"),
			},
			{
				Name:     "num_total_logical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumTotalLogicalBytes"),
			},
			{
				Name:     "num_total_physical_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumTotalPhysicalBytes"),
			},
			{
				Name:     "range_partitioning",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RangePartitioning"),
			},
			{
				Name:     "require_partition_filter",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RequirePartitionFilter"),
			},
			{
				Name:     "schema",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Schema"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
			},
			{
				Name:     "snapshot_definition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SnapshotDefinition"),
			},
			{
				Name:     "streaming_buffer",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StreamingBuffer"),
			},
			{
				Name:     "table_reference",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TableReference"),
			},
			{
				Name:     "time_partitioning",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TimePartitioning"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "view",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("View"),
			},
		},
	}
}
