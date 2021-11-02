//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CommunityGalleryImageVersionsClient contains the methods for the CommunityGalleryImageVersions group.
// Don't use this type directly, use NewCommunityGalleryImageVersionsClient() instead.
type CommunityGalleryImageVersionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCommunityGalleryImageVersionsClient creates a new instance of CommunityGalleryImageVersionsClient with the specified values.
func NewCommunityGalleryImageVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CommunityGalleryImageVersionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CommunityGalleryImageVersionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get a community gallery image version.
// If the operation fails it returns the *CloudError error type.
func (client *CommunityGalleryImageVersionsClient) Get(ctx context.Context, location string, publicGalleryName string, galleryImageName string, galleryImageVersionName string, options *CommunityGalleryImageVersionsGetOptions) (CommunityGalleryImageVersionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publicGalleryName, galleryImageName, galleryImageVersionName, options)
	if err != nil {
		return CommunityGalleryImageVersionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommunityGalleryImageVersionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunityGalleryImageVersionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CommunityGalleryImageVersionsClient) getCreateRequest(ctx context.Context, location string, publicGalleryName string, galleryImageName string, galleryImageVersionName string, options *CommunityGalleryImageVersionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publicGalleryName == "" {
		return nil, errors.New("parameter publicGalleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicGalleryName}", url.PathEscape(publicGalleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	if galleryImageVersionName == "" {
		return nil, errors.New("parameter galleryImageVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommunityGalleryImageVersionsClient) getHandleResponse(resp *http.Response) (CommunityGalleryImageVersionsGetResponse, error) {
	result := CommunityGalleryImageVersionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunityGalleryImageVersion); err != nil {
		return CommunityGalleryImageVersionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CommunityGalleryImageVersionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}