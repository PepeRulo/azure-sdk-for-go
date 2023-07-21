//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// RestorePointCollectionsClient contains the methods for the RestorePointCollections group.
// Don't use this type directly, use NewRestorePointCollectionsClient() instead.
type RestorePointCollectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorePointCollectionsClient creates a new instance of RestorePointCollectionsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorePointCollectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorePointCollectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".RestorePointCollectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorePointCollectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - The operation to create or update the restore point collection. Please refer to https://aka.ms/RestorePoints
// for more details. When updating a restore point collection, only tags may be modified.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - restorePointCollectionName - The name of the restore point collection.
//   - parameters - Parameters supplied to the Create or Update restore point collection operation.
//   - options - RestorePointCollectionsClientCreateOrUpdateOptions contains the optional parameters for the RestorePointCollectionsClient.CreateOrUpdate
//     method.
func (client *RestorePointCollectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollection, options *RestorePointCollectionsClientCreateOrUpdateOptions) (RestorePointCollectionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "RestorePointCollectionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, restorePointCollectionName, parameters, options)
	if err != nil {
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RestorePointCollectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollection, options *RestorePointCollectionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RestorePointCollectionsClient) createOrUpdateHandleResponse(resp *http.Response) (RestorePointCollectionsClientCreateOrUpdateResponse, error) {
	result := RestorePointCollectionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollection); err != nil {
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - The operation to delete the restore point collection. This operation will also delete all the contained restore
// points.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - restorePointCollectionName - The name of the Restore Point Collection.
//   - options - RestorePointCollectionsClientBeginDeleteOptions contains the optional parameters for the RestorePointCollectionsClient.BeginDelete
//     method.
func (client *RestorePointCollectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientBeginDeleteOptions) (*runtime.Poller[RestorePointCollectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, restorePointCollectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[RestorePointCollectionsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RestorePointCollectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete the restore point collection. This operation will also delete all the contained restore
// points.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *RestorePointCollectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RestorePointCollectionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, restorePointCollectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RestorePointCollectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation to get the restore point collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - restorePointCollectionName - The name of the restore point collection.
//   - options - RestorePointCollectionsClientGetOptions contains the optional parameters for the RestorePointCollectionsClient.Get
//     method.
func (client *RestorePointCollectionsClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientGetOptions) (RestorePointCollectionsClientGetResponse, error) {
	var err error
	const operationName = "RestorePointCollectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, restorePointCollectionName, options)
	if err != nil {
		return RestorePointCollectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RestorePointCollectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RestorePointCollectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RestorePointCollectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorePointCollectionsClient) getHandleResponse(resp *http.Response) (RestorePointCollectionsClientGetResponse, error) {
	result := RestorePointCollectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollection); err != nil {
		return RestorePointCollectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of restore point collections in a resource group.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - options - RestorePointCollectionsClientListOptions contains the optional parameters for the RestorePointCollectionsClient.NewListPager
//     method.
func (client *RestorePointCollectionsClient) NewListPager(resourceGroupName string, options *RestorePointCollectionsClientListOptions) *runtime.Pager[RestorePointCollectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorePointCollectionsClientListResponse]{
		More: func(page RestorePointCollectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RestorePointCollectionsClientListResponse) (RestorePointCollectionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RestorePointCollectionsClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RestorePointCollectionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorePointCollectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorePointCollectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RestorePointCollectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *RestorePointCollectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorePointCollectionsClient) listHandleResponse(resp *http.Response) (RestorePointCollectionsClientListResponse, error) {
	result := RestorePointCollectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollectionListResult); err != nil {
		return RestorePointCollectionsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets the list of restore point collections in the subscription. Use nextLink property in the response
// to get the next page of restore point collections. Do this till nextLink is not null to fetch all
// the restore point collections.
//
// Generated from API version 2023-03-01
//   - options - RestorePointCollectionsClientListAllOptions contains the optional parameters for the RestorePointCollectionsClient.NewListAllPager
//     method.
func (client *RestorePointCollectionsClient) NewListAllPager(options *RestorePointCollectionsClientListAllOptions) *runtime.Pager[RestorePointCollectionsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorePointCollectionsClientListAllResponse]{
		More: func(page RestorePointCollectionsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RestorePointCollectionsClientListAllResponse) (RestorePointCollectionsClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RestorePointCollectionsClient.NewListAllPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RestorePointCollectionsClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorePointCollectionsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorePointCollectionsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *RestorePointCollectionsClient) listAllCreateRequest(ctx context.Context, options *RestorePointCollectionsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/restorePointCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *RestorePointCollectionsClient) listAllHandleResponse(resp *http.Response) (RestorePointCollectionsClientListAllResponse, error) {
	result := RestorePointCollectionsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollectionListResult); err != nil {
		return RestorePointCollectionsClientListAllResponse{}, err
	}
	return result, nil
}

// Update - The operation to update the restore point collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - restorePointCollectionName - The name of the restore point collection.
//   - parameters - Parameters supplied to the Update restore point collection operation.
//   - options - RestorePointCollectionsClientUpdateOptions contains the optional parameters for the RestorePointCollectionsClient.Update
//     method.
func (client *RestorePointCollectionsClient) Update(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollectionUpdate, options *RestorePointCollectionsClientUpdateOptions) (RestorePointCollectionsClientUpdateResponse, error) {
	var err error
	const operationName = "RestorePointCollectionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, restorePointCollectionName, parameters, options)
	if err != nil {
		return RestorePointCollectionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RestorePointCollectionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RestorePointCollectionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *RestorePointCollectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollectionUpdate, options *RestorePointCollectionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *RestorePointCollectionsClient) updateHandleResponse(resp *http.Response) (RestorePointCollectionsClientUpdateResponse, error) {
	result := RestorePointCollectionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollection); err != nil {
		return RestorePointCollectionsClientUpdateResponse{}, err
	}
	return result, nil
}
