// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"

func init() {
	tables := []Table{
		{
			Service:        "armdashboard",
			Name:           "grafana",
			Struct:         &armdashboard.ManagedGrafana{},
			ResponseStruct: &armdashboard.GrafanaClientListResponse{},
			Client:         &armdashboard.GrafanaClient{},
			ListFunc:       (&armdashboard.GrafanaClient{}).NewListPager,
			NewFunc:        armdashboard.NewGrafanaClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Dashboard/grafana",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_dashboard)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
