// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"

func Armsupport() []*Table {
	tables := []*Table{
		{
			NewFunc:        armsupport.NewServicesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
			URL:            "/providers/Microsoft.Support/services",
			Namespace:      "microsoft.support",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_support)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ServicesClientListResponse",
		},
		{
			NewFunc:        armsupport.NewTicketsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets",
			Namespace:      "microsoft.support",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_support)`,
			Pager:          `NewListPager`,
			ResponseStruct: "TicketsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsupport())
}
