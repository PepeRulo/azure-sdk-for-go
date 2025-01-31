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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/Operations_List.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListResult = armresourcehealth.OperationListResult{
	// 	Value: []*armresourcehealth.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/Operations/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Get the operations available for the Microsoft ResourceHealth"),
	// 				Operation: to.Ptr("Get Operations"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Availability Status"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/Notifications/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Receives ARM notifications"),
	// 				Operation: to.Ptr("Receive notification"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Notification"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Gets the availability statuses for all resources in the specified scope"),
	// 				Operation: to.Ptr("Get Availability Statuses"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Availability Status"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses/current/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Gets the availability status for the specified resource"),
	// 				Operation: to.Ptr("Get Availability Status"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Availability Status"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/events/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Get Service Health Events for given subscription"),
	// 				Operation: to.Ptr("Get Service Health Events"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Events"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/emergingissues/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Get Azure services' emerging issues"),
	// 				Operation: to.Ptr("Get Azure Emerging Issues"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Emerging Issues"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/potentialoutages/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/impactedResources/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Get Impacted Resources for given subscription"),
	// 				Operation: to.Ptr("Get Impacted Resources"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Impacted Resources"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/register/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Registers the subscription for the Microsoft ResourceHealth"),
	// 				Operation: to.Ptr("Register with the Provider"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Registration"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/unregister/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Unregisters the subscription for the Microsoft ResourceHealth"),
	// 				Operation: to.Ptr("Unregister with the Provider"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Unregistration"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Resourcehealth/healthevent/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Denotes the change in health state for the specified resource"),
	// 				Operation: to.Ptr("HealthEvent Change"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Availability Status"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Resourcehealth/healthevent/Activated/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Denotes the change in health state for the specified resource"),
	// 				Operation: to.Ptr("Health Event Activated"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Health Event"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Resourcehealth/healthevent/Updated/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Denotes the change in health state for the specified resource"),
	// 				Operation: to.Ptr("Health Event Updated"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Health Event"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Resourcehealth/healthevent/Resolved/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Denotes the change in health state for the specified resource"),
	// 				Operation: to.Ptr("Health Event Resolved"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Health Event"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Resourcehealth/healthevent/InProgress/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Denotes the change in health state for the specified resource"),
	// 				Operation: to.Ptr("Health Event InProgress"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Health Event"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Resourcehealth/healthevent/Pending/action"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Denotes the change in health state for the specified resource"),
	// 				Operation: to.Ptr("Health Event Pending"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Health Event"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.ResourceHealth/metadata/read"),
	// 			Display: &armresourcehealth.OperationDisplay{
	// 				Description: to.Ptr("Gets Metadata"),
	// 				Operation: to.Ptr("Read Metadata"),
	// 				Provider: to.Ptr("Microsoft ResourceHealth"),
	// 				Resource: to.Ptr("Metadata"),
	// 			},
	// 	}},
	// }
}
