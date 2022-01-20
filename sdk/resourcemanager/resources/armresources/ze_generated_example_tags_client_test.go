//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutTagsResource.json
func ExampleTagsClient_CreateOrUpdateAtScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresources.NewTagsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateAtScope(ctx,
		"<scope>",
		armresources.TagsResource{
			Properties: &armresources.Tags{
				Tags: map[string]*string{
					"tagKey1": to.StringPtr("tag-value-1"),
					"tagKey2": to.StringPtr("tag-value-2"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagsClientCreateOrUpdateAtScopeResult)
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetTagsResource.json
func ExampleTagsClient_GetAtScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresources.NewTagsClient("<subscription-id>", cred, nil)
	res, err := client.GetAtScope(ctx,
		"<scope>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagsClientGetAtScopeResult)
}