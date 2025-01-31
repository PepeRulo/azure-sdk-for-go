//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetPreconfiguredEndpoints.json
func ExamplePreconfiguredEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPreconfiguredEndpointsClient().NewListPager("MyResourceGroup", "MyProfile", nil)
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
		// page.PreconfiguredEndpointList = armfrontdoor.PreconfiguredEndpointList{
		// 	Value: []*armfrontdoor.PreconfiguredEndpoint{
		// 		{
		// 			Name: to.Ptr("Endpoint 1"),
		// 			Properties: &armfrontdoor.PreconfiguredEndpointProperties{
		// 				Description: to.Ptr("this is the endpoint 1 preconfigured endpoint."),
		// 				Backend: to.Ptr("WESTUS"),
		// 				Endpoint: to.Ptr("endpoint1.net"),
		// 				EndpointType: to.Ptr(armfrontdoor.EndpointTypeAFD),
		// 			},
		// 	}},
		// }
	}
}
