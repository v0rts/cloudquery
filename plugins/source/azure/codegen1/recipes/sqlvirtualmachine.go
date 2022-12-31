// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"

func Armsqlvirtualmachine() []*Table {
	tables := []*Table{
		{
			NewFunc:        armsqlvirtualmachine.NewGroupsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
			Namespace:      "microsoft.sqlvirtualmachine",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_sqlvirtualmachine)`,
			Pager:          `NewListPager`,
			ResponseStruct: "GroupsClientListResponse",
		},
		{
			NewFunc:        armsqlvirtualmachine.NewSQLVirtualMachinesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
			Namespace:      "microsoft.sqlvirtualmachine",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_sqlvirtualmachine)`,
			Pager:          `NewListPager`,
			ResponseStruct: "SQLVirtualMachinesClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsqlvirtualmachine())
}
