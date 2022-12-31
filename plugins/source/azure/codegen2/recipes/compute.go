// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"

func init() {
	tables := []Table{
		{
			Service:        "armcompute",
			Name:           "cloud_services",
			Struct:         &armcompute.CloudService{},
			ResponseStruct: &armcompute.CloudServicesClientListAllResponse{},
			Client:         &armcompute.CloudServicesClient{},
			ListFunc:       (&armcompute.CloudServicesClient{}).NewListAllPager,
			NewFunc:        armcompute.NewCloudServicesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/cloudServices",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "disk_accesses",
			Struct:         &armcompute.DiskAccess{},
			ResponseStruct: &armcompute.DiskAccessesClientListResponse{},
			Client:         &armcompute.DiskAccessesClient{},
			ListFunc:       (&armcompute.DiskAccessesClient{}).NewListPager,
			NewFunc:        armcompute.NewDiskAccessesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskAccesses",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "disk_encryption_sets",
			Struct:         &armcompute.DiskEncryptionSet{},
			ResponseStruct: &armcompute.DiskEncryptionSetsClientListResponse{},
			Client:         &armcompute.DiskEncryptionSetsClient{},
			ListFunc:       (&armcompute.DiskEncryptionSetsClient{}).NewListPager,
			NewFunc:        armcompute.NewDiskEncryptionSetsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "disks",
			Struct:         &armcompute.Disk{},
			ResponseStruct: &armcompute.DisksClientListResponse{},
			Client:         &armcompute.DisksClient{},
			ListFunc:       (&armcompute.DisksClient{}).NewListPager,
			NewFunc:        armcompute.NewDisksClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "galleries",
			Struct:         &armcompute.Gallery{},
			ResponseStruct: &armcompute.GalleriesClientListResponse{},
			Client:         &armcompute.GalleriesClient{},
			ListFunc:       (&armcompute.GalleriesClient{}).NewListPager,
			NewFunc:        armcompute.NewGalleriesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/galleries",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "images",
			Struct:         &armcompute.Image{},
			ResponseStruct: &armcompute.ImagesClientListResponse{},
			Client:         &armcompute.ImagesClient{},
			ListFunc:       (&armcompute.ImagesClient{}).NewListPager,
			NewFunc:        armcompute.NewImagesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "restore_point_collections",
			Struct:         &armcompute.RestorePointCollection{},
			ResponseStruct: &armcompute.RestorePointCollectionsClientListAllResponse{},
			Client:         &armcompute.RestorePointCollectionsClient{},
			ListFunc:       (&armcompute.RestorePointCollectionsClient{}).NewListAllPager,
			NewFunc:        armcompute.NewRestorePointCollectionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/restorePointCollections",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "snapshots",
			Struct:         &armcompute.Snapshot{},
			ResponseStruct: &armcompute.SnapshotsClientListResponse{},
			Client:         &armcompute.SnapshotsClient{},
			ListFunc:       (&armcompute.SnapshotsClient{}).NewListPager,
			NewFunc:        armcompute.NewSnapshotsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcompute",
			Name:           "virtual_machine_scale_sets",
			Struct:         &armcompute.VirtualMachineScaleSet{},
			ResponseStruct: &armcompute.VirtualMachineScaleSetsClientListAllResponse{},
			Client:         &armcompute.VirtualMachineScaleSetsClient{},
			ListFunc:       (&armcompute.VirtualMachineScaleSetsClient{}).NewListAllPager,
			NewFunc:        armcompute.NewVirtualMachineScaleSetsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachineScaleSets",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_compute)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
