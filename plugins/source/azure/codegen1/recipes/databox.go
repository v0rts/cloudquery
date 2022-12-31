// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"

func Armdatabox() []*Table {
	tables := []*Table{
		{
			NewFunc:        armdatabox.NewJobsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs",
			Namespace:      "microsoft.databox",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_databox)`,
			Pager:          `NewListPager`,
			ResponseStruct: "JobsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdatabox())
}
