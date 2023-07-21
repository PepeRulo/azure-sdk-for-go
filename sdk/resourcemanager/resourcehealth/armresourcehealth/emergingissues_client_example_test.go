//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b74978708bb95475562412d4654c00fbcedd9f89/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/EmergingIssues_List.json
func ExampleEmergingIssuesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmergingIssuesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.EmergingIssueListResult = armresourcehealth.EmergingIssueListResult{
		// 	Value: []*armresourcehealth.EmergingIssuesGetResult{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("/providers/Microsoft.ResourceHealth/emergingissues"),
		// 			ID: to.Ptr("/providers/Microsoft.ResourceHealth/emergingissues/default"),
		// 			Properties: &armresourcehealth.EmergingIssue{
		// 				RefreshTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-17T09:12:00Z"); return t}()),
		// 				StatusActiveEvents: []*armresourcehealth.StatusActiveEvent{
		// 					{
		// 						Description: to.Ptr("Virtual Machines case"),
		// 						Cloud: to.Ptr("Public"),
		// 						Impacts: []*armresourcehealth.EmergingIssueImpact{
		// 							{
		// 								Name: to.Ptr("Virtual Machines"),
		// 								ID: to.Ptr("virtual-machines"),
		// 								Regions: []*armresourcehealth.ImpactedRegion{
		// 									{
		// 										Name: to.Ptr("Central US"),
		// 										ID: to.Ptr("us-central"),
		// 									},
		// 									{
		// 										Name: to.Ptr("East US"),
		// 										ID: to.Ptr("us-east"),
		// 								}},
		// 						}},
		// 						LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T08:10:00Z"); return t}()),
		// 						Published: to.Ptr(true),
		// 						Severity: to.Ptr(armresourcehealth.SeverityValuesInformation),
		// 						Stage: to.Ptr(armresourcehealth.StageValuesActive),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T08:06:00Z"); return t}()),
		// 						Title: to.Ptr("Automatic updates to the dial tone page from ACM - SHD event"),
		// 						TrackingID: to.Ptr("KTTK-TZ8"),
		// 					},
		// 					{
		// 						Description: to.Ptr("All Azure SQL service management requests (create, update, delete, etc.) are serviced through the SQL Control Plane infrastructure. Engineers determined that during this issue, the control plane service became unhealthy after receiving a significant increase in internally generated operations and reaching an operational threshold. This led to service management requests becoming unable to complete, which in-turn resulted in timeouts and throttling. Subsequent investigation by engineers determined this issue was due service requests from an internal Azure group that triggered a bug which caused an excessive number of internally generated operations."),
		// 						Cloud: to.Ptr("Public"),
		// 						Impacts: []*armresourcehealth.EmergingIssueImpact{
		// 							{
		// 								Name: to.Ptr("SQL Database"),
		// 								ID: to.Ptr("sql-database"),
		// 								Regions: []*armresourcehealth.ImpactedRegion{
		// 									{
		// 										Name: to.Ptr("West Europe"),
		// 										ID: to.Ptr("europe-west"),
		// 								}},
		// 						}},
		// 						LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-16T05:11:00Z"); return t}()),
		// 						Published: to.Ptr(true),
		// 						Severity: to.Ptr(armresourcehealth.SeverityValuesError),
		// 						Stage: to.Ptr(armresourcehealth.StageValuesActive),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-16T05:11:00Z"); return t}()),
		// 						Title: to.Ptr("Azure SQL Database - West Europe"),
		// 						TrackingID: to.Ptr("4KHK-LS8"),
		// 				}},
		// 				StatusBanners: []*armresourcehealth.StatusBanner{
		// 					{
		// 						Cloud: to.Ptr("Public"),
		// 						LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T08:04:00Z"); return t}()),
		// 						Message: to.Ptr("Testing backup site"),
		// 						Title: to.Ptr("Automatic updates to the dial tone page from ACM - banner"),
		// 					},
		// 					{
		// 						Cloud: to.Ptr("Public"),
		// 						LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T10:15:00Z"); return t}()),
		// 						Message: to.Ptr("<span style=\"color: #323237; font-family: &quot;Segoe UI&quot;, SegoeUI, &quot;Segoe WP&quot;, Tahoma, Arial, sans-serif; font-size: 16px; background-color: #ffffff\">A subset of customers using Storage in West Europe experienced service availability issues. In addition, resources with dependencies on the impacted storage scale units may have experienced downstream impact in the form of availability issues, connection failures, or high latency. Engineers are investigating the root cause.</span>"),
		// 						Title: to.Ptr("Storage - West Europe"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b74978708bb95475562412d4654c00fbcedd9f89/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/EmergingIssues_Get.json
func ExampleEmergingIssuesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmergingIssuesClient().Get(ctx, armresourcehealth.IssueNameParameterDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EmergingIssuesGetResult = armresourcehealth.EmergingIssuesGetResult{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("/providers/Microsoft.ResourceHealth/emergingissues"),
	// 	ID: to.Ptr("/providers/Microsoft.ResourceHealth/emergingissues/default"),
	// 	Properties: &armresourcehealth.EmergingIssue{
	// 		RefreshTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-17T09:12:00Z"); return t}()),
	// 		StatusActiveEvents: []*armresourcehealth.StatusActiveEvent{
	// 			{
	// 				Description: to.Ptr("Virtual Machines case"),
	// 				Cloud: to.Ptr("Public"),
	// 				Impacts: []*armresourcehealth.EmergingIssueImpact{
	// 					{
	// 						Name: to.Ptr("Virtual Machines"),
	// 						ID: to.Ptr("virtual-machines"),
	// 						Regions: []*armresourcehealth.ImpactedRegion{
	// 							{
	// 								Name: to.Ptr("Central US"),
	// 								ID: to.Ptr("us-central"),
	// 							},
	// 							{
	// 								Name: to.Ptr("East US"),
	// 								ID: to.Ptr("us-east"),
	// 						}},
	// 				}},
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T08:10:00Z"); return t}()),
	// 				Published: to.Ptr(true),
	// 				Severity: to.Ptr(armresourcehealth.SeverityValuesInformation),
	// 				Stage: to.Ptr(armresourcehealth.StageValuesActive),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T08:06:00Z"); return t}()),
	// 				Title: to.Ptr("Automatic updates to the dial tone page from ACM - SHD event"),
	// 				TrackingID: to.Ptr("KTTK-TZ8"),
	// 			},
	// 			{
	// 				Description: to.Ptr("All Azure SQL service management requests (create, update, delete, etc.) are serviced through the SQL Control Plane infrastructure. Engineers determined that during this issue, the control plane service became unhealthy after receiving a significant increase in internally generated operations and reaching an operational threshold. This led to service management requests becoming unable to complete, which in-turn resulted in timeouts and throttling. Subsequent investigation by engineers determined this issue was due service requests from an internal Azure group that triggered a bug which caused an excessive number of internally generated operations."),
	// 				Cloud: to.Ptr("Public"),
	// 				Impacts: []*armresourcehealth.EmergingIssueImpact{
	// 					{
	// 						Name: to.Ptr("SQL Database"),
	// 						ID: to.Ptr("sql-database"),
	// 						Regions: []*armresourcehealth.ImpactedRegion{
	// 							{
	// 								Name: to.Ptr("West Europe"),
	// 								ID: to.Ptr("europe-west"),
	// 						}},
	// 				}},
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-16T05:11:00Z"); return t}()),
	// 				Published: to.Ptr(true),
	// 				Severity: to.Ptr(armresourcehealth.SeverityValuesError),
	// 				Stage: to.Ptr(armresourcehealth.StageValuesActive),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-16T05:11:00Z"); return t}()),
	// 				Title: to.Ptr("Azure SQL Database - West Europe"),
	// 				TrackingID: to.Ptr("4KHK-LS8"),
	// 		}},
	// 		StatusBanners: []*armresourcehealth.StatusBanner{
	// 			{
	// 				Cloud: to.Ptr("Public"),
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T08:04:00Z"); return t}()),
	// 				Message: to.Ptr("Testing backup site"),
	// 				Title: to.Ptr("Automatic updates to the dial tone page from ACM - banner"),
	// 			},
	// 			{
	// 				Cloud: to.Ptr("Public"),
	// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-15T10:15:00Z"); return t}()),
	// 				Message: to.Ptr("<span style=\"color: #323237; font-family: &quot;Segoe UI&quot;, SegoeUI, &quot;Segoe WP&quot;, Tahoma, Arial, sans-serif; font-size: 16px; background-color: #ffffff\">A subset of customers using Storage in West Europe experienced service availability issues. In addition, resources with dependencies on the impacted storage scale units may have experienced downstream impact in the form of availability issues, connection failures, or high latency. Engineers are investigating the root cause.</span>"),
	// 				Title: to.Ptr("Storage - West Europe"),
	// 		}},
	// 	},
	// }
}
