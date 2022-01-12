//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// MachineExtensionsClient contains the methods for the MachineExtensions group.
// Don't use this type directly, use NewMachineExtensionsClient() instead.
type MachineExtensionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMachineExtensionsClient creates a new instance of MachineExtensionsClient with the specified values.
func NewMachineExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MachineExtensionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &MachineExtensionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - The operation to create or update the extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtension, options *MachineExtensionsBeginCreateOrUpdateOptions) (MachineExtensionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
	if err != nil {
		return MachineExtensionsCreateOrUpdatePollerResponse{}, err
	}
	result := MachineExtensionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MachineExtensionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return MachineExtensionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &MachineExtensionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update the extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtension, options *MachineExtensionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MachineExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtension, options *MachineExtensionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *MachineExtensionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - The operation to delete the extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsBeginDeleteOptions) (MachineExtensionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, machineName, extensionName, options)
	if err != nil {
		return MachineExtensionsDeletePollerResponse{}, err
	}
	result := MachineExtensionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MachineExtensionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return MachineExtensionsDeletePollerResponse{}, err
	}
	result.Poller = &MachineExtensionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The operation to delete the extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, machineName, extensionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MachineExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *MachineExtensionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - The operation to get the extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) Get(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsGetOptions) (MachineExtensionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, extensionName, options)
	if err != nil {
		return MachineExtensionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MachineExtensionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MachineExtensionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MachineExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MachineExtensionsClient) getHandleResponse(resp *http.Response) (MachineExtensionsGetResponse, error) {
	result := MachineExtensionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineExtension); err != nil {
		return MachineExtensionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MachineExtensionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - The operation to get all extensions of a non-Azure machine
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) List(resourceGroupName string, machineName string, options *MachineExtensionsListOptions) *MachineExtensionsListPager {
	return &MachineExtensionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, machineName, options)
		},
		advancer: func(ctx context.Context, resp MachineExtensionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MachineExtensionsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MachineExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *MachineExtensionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MachineExtensionsClient) listHandleResponse(resp *http.Response) (MachineExtensionsListResponse, error) {
	result := MachineExtensionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineExtensionsListResult); err != nil {
		return MachineExtensionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MachineExtensionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - The operation to create or update the extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtensionUpdate, options *MachineExtensionsBeginUpdateOptions) (MachineExtensionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
	if err != nil {
		return MachineExtensionsUpdatePollerResponse{}, err
	}
	result := MachineExtensionsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MachineExtensionsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return MachineExtensionsUpdatePollerResponse{}, err
	}
	result.Poller = &MachineExtensionsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The operation to create or update the extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MachineExtensionsClient) update(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtensionUpdate, options *MachineExtensionsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *MachineExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtensionUpdate, options *MachineExtensionsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}

// updateHandleError handles the Update error response.
func (client *MachineExtensionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}