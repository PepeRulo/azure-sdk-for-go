//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ListCertificates.json
func ExampleCertificatesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificatesClient("<subscription-id>", cred, nil)
	pager := client.List(&armappservice.CertificatesListOptions{Filter: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Certificate.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ListCertificatesByResourceGroup.json
func ExampleCertificatesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificatesClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Certificate.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetCertificate.json
func ExampleCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Certificate.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/CreateOrUpdateCertificate.json
func ExampleCertificatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.Certificate{
			Resource: armappservice.Resource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armappservice.CertificateProperties{
				HostNames: []*string{
					to.StringPtr("ServerCert")},
				Password: to.StringPtr("<password>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Certificate.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/DeleteCertificate.json
func ExampleCertificatesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificatesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/PatchCertificate.json
func ExampleCertificatesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.CertificatePatchResource{
			Properties: &armappservice.CertificatePatchResourceProperties{
				Password: to.StringPtr("<password>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Certificate.ID: %s\n", *res.ID)
}
