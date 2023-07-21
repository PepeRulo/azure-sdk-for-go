//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GitHubOAuth.json
func ExampleDeveloperHubServiceClient_GitHubOAuth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeveloperHubServiceClient().GitHubOAuth(ctx, "eastus2euap", &armdevhub.DeveloperHubServiceClientGitHubOAuthOptions{Parameters: &armdevhub.GitHubOAuthCallRequest{
		RedirectURL: to.Ptr("https://ms.portal.azure.com/aks"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubOAuthInfoResponse = armdevhub.GitHubOAuthInfoResponse{
	// 	AuthURL: to.Ptr("https://github.com/login/oauth/authorize?client_id=11111111&redirect_uri=https://management.azure.com/subscriptions/subscriptionId1/providers/Microsoft.DevHub/locations/eastus2euap/githuboauth&state=2345678-3456-7890-5678-012345678901&scope=repo,user:email,workflow"),
	// 	Token: to.Ptr(""),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GitHubOAuthCallback.json
func ExampleDeveloperHubServiceClient_GitHubOAuthCallback() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeveloperHubServiceClient().GitHubOAuthCallback(ctx, "eastus2euap", "3584d83530557fdd1f46af8289938c8ef79f9dc5", "12345678-3456-7890-5678-012345678901", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubOAuthResponse = armdevhub.GitHubOAuthResponse{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DevHub/locations/githuboauth"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId1/providers/Microsoft.DevHub/locations/eastus2euap/githuboauth/default"),
	// 	Properties: &armdevhub.GitHubOAuthProperties{
	// 		Username: to.Ptr("user"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GitHubOAuth_List.json
func ExampleDeveloperHubServiceClient_ListGitHubOAuth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeveloperHubServiceClient().ListGitHubOAuth(ctx, "eastus2euap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubOAuthListResponse = armdevhub.GitHubOAuthListResponse{
	// 	Value: []*armdevhub.GitHubOAuthResponse{
	// 		{
	// 			Name: to.Ptr("default"),
	// 			Type: to.Ptr("Microsoft.DevHub/locations/githuboauth"),
	// 			ID: to.Ptr("/subscriptions/subscriptionId1/providers/Microsoft.DevHub/locations/eastus2euap/githuboauth/default"),
	// 			Properties: &armdevhub.GitHubOAuthProperties{
	// 				Username: to.Ptr("user"),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GeneratePreviewArtifacts.json
func ExampleDeveloperHubServiceClient_GeneratePreviewArtifacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeveloperHubServiceClient().GeneratePreviewArtifacts(ctx, "location1", armdevhub.ArtifactGenerationProperties{
		AppName:                   to.Ptr("my-app"),
		DockerfileGenerationMode:  to.Ptr(armdevhub.DockerfileGenerationModeEnabled),
		DockerfileOutputDirectory: to.Ptr("./"),
		GenerationLanguage:        to.Ptr(armdevhub.GenerationLanguageJavascript),
		ImageName:                 to.Ptr("myimage"),
		ImageTag:                  to.Ptr("latest"),
		LanguageVersion:           to.Ptr("14"),
		ManifestGenerationMode:    to.Ptr(armdevhub.ManifestGenerationModeEnabled),
		ManifestOutputDirectory:   to.Ptr("./"),
		ManifestType:              to.Ptr(armdevhub.GenerationManifestTypeKube),
		Namespace:                 to.Ptr("my-namespace"),
		Port:                      to.Ptr("80"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Value = map[string]*string{
	// 	"dockerfiles/example-dockerfile": to.Ptr("dockerfile-content \n including newlines"),
	// 	"manifests/example-manifest-file-1": to.Ptr("manifest file 1 content \n including newlines"),
	// 	"manifests/example-manifest-file-2": to.Ptr("manifest file 2 content \n including newlines"),
	// }
}
