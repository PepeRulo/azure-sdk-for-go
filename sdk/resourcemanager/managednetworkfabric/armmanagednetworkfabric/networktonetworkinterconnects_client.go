//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NetworkToNetworkInterconnectsClient contains the methods for the NetworkToNetworkInterconnects group.
// Don't use this type directly, use NewNetworkToNetworkInterconnectsClient() instead.
type NetworkToNetworkInterconnectsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkToNetworkInterconnectsClient creates a new instance of NetworkToNetworkInterconnectsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkToNetworkInterconnectsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkToNetworkInterconnectsClient, error) {
	cl, err := arm.NewClient(moduleName+".NetworkToNetworkInterconnectsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkToNetworkInterconnectsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Configuration used to setup CE-PE connectivity PUT Method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the NetworkFabric.
//   - networkToNetworkInterconnectName - Name of the NetworkToNetworkInterconnectName
//   - body - Request payload.
//   - options - NetworkToNetworkInterconnectsClientBeginCreateOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.BeginCreate
//     method.
func (client *NetworkToNetworkInterconnectsClient) BeginCreate(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnect, options *NetworkToNetworkInterconnectsClientBeginCreateOptions) (*runtime.Poller[NetworkToNetworkInterconnectsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkToNetworkInterconnectsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkToNetworkInterconnectsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Configuration used to setup CE-PE connectivity PUT Method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *NetworkToNetworkInterconnectsClient) create(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnect, options *NetworkToNetworkInterconnectsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *NetworkToNetworkInterconnectsClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnect, options *NetworkToNetworkInterconnectsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Implements NetworkToNetworkInterconnects DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the NetworkFabric.
//   - networkToNetworkInterconnectName - Name of the NetworkToNetworkInterconnectName
//   - options - NetworkToNetworkInterconnectsClientBeginDeleteOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.BeginDelete
//     method.
func (client *NetworkToNetworkInterconnectsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientBeginDeleteOptions) (*runtime.Poller[NetworkToNetworkInterconnectsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkToNetworkInterconnectsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkToNetworkInterconnectsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Implements NetworkToNetworkInterconnects DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *NetworkToNetworkInterconnectsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkToNetworkInterconnectsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements NetworkToNetworkInterconnects GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the NetworkFabric.
//   - networkToNetworkInterconnectName - Name of the NetworkToNetworkInterconnect
//   - options - NetworkToNetworkInterconnectsClientGetOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.Get
//     method.
func (client *NetworkToNetworkInterconnectsClient) Get(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientGetOptions) (NetworkToNetworkInterconnectsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, options)
	if err != nil {
		return NetworkToNetworkInterconnectsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkToNetworkInterconnectsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkToNetworkInterconnectsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkToNetworkInterconnectsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkToNetworkInterconnectsClient) getHandleResponse(resp *http.Response) (NetworkToNetworkInterconnectsClientGetResponse, error) {
	result := NetworkToNetworkInterconnectsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkToNetworkInterconnect); err != nil {
		return NetworkToNetworkInterconnectsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Implements Network To Network Interconnects list by Network Fabric GET method.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the NetworkFabric.
//   - options - NetworkToNetworkInterconnectsClientListOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.NewListPager
//     method.
func (client *NetworkToNetworkInterconnectsClient) NewListPager(resourceGroupName string, networkFabricName string, options *NetworkToNetworkInterconnectsClientListOptions) *runtime.Pager[NetworkToNetworkInterconnectsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkToNetworkInterconnectsClientListResponse]{
		More: func(page NetworkToNetworkInterconnectsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkToNetworkInterconnectsClientListResponse) (NetworkToNetworkInterconnectsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkFabricName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkToNetworkInterconnectsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkToNetworkInterconnectsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkToNetworkInterconnectsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *NetworkToNetworkInterconnectsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, options *NetworkToNetworkInterconnectsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkToNetworkInterconnectsClient) listHandleResponse(resp *http.Response) (NetworkToNetworkInterconnectsClientListResponse, error) {
	result := NetworkToNetworkInterconnectsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkToNetworkInterconnectsList); err != nil {
		return NetworkToNetworkInterconnectsClientListResponse{}, err
	}
	return result, nil
}
