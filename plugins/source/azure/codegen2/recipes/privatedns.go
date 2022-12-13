// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"

func init() {
	tables := []Table{
		{
			Service:        "armprivatedns",
			Name:           "private_zones",
			Struct:         &armprivatedns.PrivateZone{},
			ResponseStruct: &armprivatedns.PrivateZonesClientListResponse{},
			Client:         &armprivatedns.PrivateZonesClient{},
			ListFunc:       (&armprivatedns.PrivateZonesClient{}).NewListPager,
			NewFunc:        armprivatedns.NewPrivateZonesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateDnsZones",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
