//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/EmailRegistrations_ActivateEmail.json
func ExampleEmailRegistrationsClient_ActivateEmail() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmailRegistrationsClient().ActivateEmail(ctx, "East US 2", armdatashare.EmailRegistration{
		ActivationCode: to.Ptr("djsfhakj2lekowd3wepfklpwe9lpflcd"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EmailRegistration = armdatashare.EmailRegistration{
	// 	ActivationExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-12T16:44:53.688049Z"); return t}()),
	// 	Email: to.Ptr("receiver@microsoft.com"),
	// 	RegistrationStatus: to.Ptr(armdatashare.RegistrationStatusActivated),
	// 	TenantID: to.Ptr("9f532315-b048-4374-8de1-14734d9b7f77"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/EmailRegistrations_RegisterEmail.json
func ExampleEmailRegistrationsClient_RegisterEmail() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmailRegistrationsClient().RegisterEmail(ctx, "East US 2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EmailRegistration = armdatashare.EmailRegistration{
	// 	ActivationExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-12T16:44:53.688049Z"); return t}()),
	// 	Email: to.Ptr("receiver@microsoft.com"),
	// 	RegistrationStatus: to.Ptr(armdatashare.RegistrationStatusActivationPending),
	// 	TenantID: to.Ptr("9f532315-b048-4374-8de1-14734d9b7f77"),
	// }
}
