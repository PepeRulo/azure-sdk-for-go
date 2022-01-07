//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsignalr_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalRPrivateEndpointConnections_List.json
func ExampleSignalRPrivateEndpointConnectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<resource-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("PrivateEndpointConnection.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalRPrivateEndpointConnections_Get.json
func ExampleSignalRPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<private-endpoint-connection-name>",
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateEndpointConnection.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalRPrivateEndpointConnections_Update.json
func ExampleSignalRPrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<private-endpoint-connection-name>",
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.PrivateEndpointConnection{
			Properties: &armsignalr.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armsignalr.PrivateEndpoint{
					ID: to.StringPtr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armsignalr.PrivateLinkServiceConnectionState{
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          armsignalr.PrivateLinkServiceConnectionStatusApproved.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateEndpointConnection.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalRPrivateEndpointConnections_Delete.json
func ExampleSignalRPrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<private-endpoint-connection-name>",
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
