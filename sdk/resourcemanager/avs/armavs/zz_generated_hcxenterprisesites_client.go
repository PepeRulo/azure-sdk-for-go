//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// HcxEnterpriseSitesClient contains the methods for the HcxEnterpriseSites group.
// Don't use this type directly, use NewHcxEnterpriseSitesClient() instead.
type HcxEnterpriseSitesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewHcxEnterpriseSitesClient creates a new instance of HcxEnterpriseSitesClient with the specified values.
func NewHcxEnterpriseSitesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *HcxEnterpriseSitesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &HcxEnterpriseSitesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update an HCX Enterprise Site in a private cloud
// If the operation fails it returns the *CloudError error type.
func (client *HcxEnterpriseSitesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, hcxEnterpriseSite HcxEnterpriseSite, options *HcxEnterpriseSitesCreateOrUpdateOptions) (HcxEnterpriseSitesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, hcxEnterpriseSiteName, hcxEnterpriseSite, options)
	if err != nil {
		return HcxEnterpriseSitesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HcxEnterpriseSitesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HcxEnterpriseSitesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HcxEnterpriseSitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, hcxEnterpriseSite HcxEnterpriseSite, options *HcxEnterpriseSitesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites/{hcxEnterpriseSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if hcxEnterpriseSiteName == "" {
		return nil, errors.New("parameter hcxEnterpriseSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcxEnterpriseSiteName}", url.PathEscape(hcxEnterpriseSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, hcxEnterpriseSite)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HcxEnterpriseSitesClient) createOrUpdateHandleResponse(resp *http.Response) (HcxEnterpriseSitesCreateOrUpdateResponse, error) {
	result := HcxEnterpriseSitesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcxEnterpriseSite); err != nil {
		return HcxEnterpriseSitesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *HcxEnterpriseSitesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete an HCX Enterprise Site in a private cloud
// If the operation fails it returns the *CloudError error type.
func (client *HcxEnterpriseSitesClient) Delete(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesDeleteOptions) (HcxEnterpriseSitesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, hcxEnterpriseSiteName, options)
	if err != nil {
		return HcxEnterpriseSitesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HcxEnterpriseSitesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HcxEnterpriseSitesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return HcxEnterpriseSitesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HcxEnterpriseSitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites/{hcxEnterpriseSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if hcxEnterpriseSiteName == "" {
		return nil, errors.New("parameter hcxEnterpriseSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcxEnterpriseSiteName}", url.PathEscape(hcxEnterpriseSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *HcxEnterpriseSitesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get an HCX Enterprise Site by name in a private cloud
// If the operation fails it returns the *CloudError error type.
func (client *HcxEnterpriseSitesClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesGetOptions) (HcxEnterpriseSitesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, hcxEnterpriseSiteName, options)
	if err != nil {
		return HcxEnterpriseSitesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HcxEnterpriseSitesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HcxEnterpriseSitesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HcxEnterpriseSitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites/{hcxEnterpriseSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if hcxEnterpriseSiteName == "" {
		return nil, errors.New("parameter hcxEnterpriseSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcxEnterpriseSiteName}", url.PathEscape(hcxEnterpriseSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HcxEnterpriseSitesClient) getHandleResponse(resp *http.Response) (HcxEnterpriseSitesGetResponse, error) {
	result := HcxEnterpriseSitesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcxEnterpriseSite); err != nil {
		return HcxEnterpriseSitesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *HcxEnterpriseSitesClient) getHandleError(resp *http.Response) error {
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

// List - List HCX Enterprise Sites in a private cloud
// If the operation fails it returns the *CloudError error type.
func (client *HcxEnterpriseSitesClient) List(resourceGroupName string, privateCloudName string, options *HcxEnterpriseSitesListOptions) *HcxEnterpriseSitesListPager {
	return &HcxEnterpriseSitesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
		},
		advancer: func(ctx context.Context, resp HcxEnterpriseSitesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.HcxEnterpriseSiteList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *HcxEnterpriseSitesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *HcxEnterpriseSitesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HcxEnterpriseSitesClient) listHandleResponse(resp *http.Response) (HcxEnterpriseSitesListResponse, error) {
	result := HcxEnterpriseSitesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcxEnterpriseSiteList); err != nil {
		return HcxEnterpriseSitesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *HcxEnterpriseSitesClient) listHandleError(resp *http.Response) error {
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
