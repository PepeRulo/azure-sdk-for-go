//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// MetricAlertsClient contains the methods for the MetricAlerts group.
// Don't use this type directly, use NewMetricAlertsClient() instead.
type MetricAlertsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMetricAlertsClient creates a new instance of MetricAlertsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMetricAlertsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MetricAlertsClient, error) {
	cl, err := arm.NewClient(moduleName+".MetricAlertsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MetricAlertsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an metric alert definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - parameters - The parameters of the rule to create or update.
//   - options - MetricAlertsClientCreateOrUpdateOptions contains the optional parameters for the MetricAlertsClient.CreateOrUpdate
//     method.
func (client *MetricAlertsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResource, options *MetricAlertsClientCreateOrUpdateOptions) (MetricAlertsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MetricAlertsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResource, options *MetricAlertsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MetricAlertsClient) createOrUpdateHandleResponse(resp *http.Response) (MetricAlertsClientCreateOrUpdateResponse, error) {
	result := MetricAlertsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResource); err != nil {
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an alert rule definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - options - MetricAlertsClientDeleteOptions contains the optional parameters for the MetricAlertsClient.Delete method.
func (client *MetricAlertsClient) Delete(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsClientDeleteOptions) (MetricAlertsClientDeleteResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return MetricAlertsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientDeleteResponse{}, err
	}
	return MetricAlertsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MetricAlertsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve an alert rule definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - options - MetricAlertsClientGetOptions contains the optional parameters for the MetricAlertsClient.Get method.
func (client *MetricAlertsClient) Get(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsClientGetOptions) (MetricAlertsClientGetResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return MetricAlertsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MetricAlertsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MetricAlertsClient) getHandleResponse(resp *http.Response) (MetricAlertsClientGetResponse, error) {
	result := MetricAlertsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResource); err != nil {
		return MetricAlertsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve alert rule definitions in a resource group.
//
// Generated from API version 2018-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MetricAlertsClientListByResourceGroupOptions contains the optional parameters for the MetricAlertsClient.NewListByResourceGroupPager
//     method.
func (client *MetricAlertsClient) NewListByResourceGroupPager(resourceGroupName string, options *MetricAlertsClientListByResourceGroupOptions) *runtime.Pager[MetricAlertsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetricAlertsClientListByResourceGroupResponse]{
		More: func(page MetricAlertsClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *MetricAlertsClientListByResourceGroupResponse) (MetricAlertsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MetricAlertsClient.NewListByResourceGroupPager")
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return MetricAlertsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return MetricAlertsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MetricAlertsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MetricAlertsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *MetricAlertsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts"
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
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MetricAlertsClient) listByResourceGroupHandleResponse(resp *http.Response) (MetricAlertsClientListByResourceGroupResponse, error) {
	result := MetricAlertsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResourceCollection); err != nil {
		return MetricAlertsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieve alert rule definitions in a subscription.
//
// Generated from API version 2018-03-01
//   - options - MetricAlertsClientListBySubscriptionOptions contains the optional parameters for the MetricAlertsClient.NewListBySubscriptionPager
//     method.
func (client *MetricAlertsClient) NewListBySubscriptionPager(options *MetricAlertsClientListBySubscriptionOptions) *runtime.Pager[MetricAlertsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetricAlertsClientListBySubscriptionResponse]{
		More: func(page MetricAlertsClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *MetricAlertsClientListBySubscriptionResponse) (MetricAlertsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MetricAlertsClient.NewListBySubscriptionPager")
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return MetricAlertsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return MetricAlertsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MetricAlertsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MetricAlertsClient) listBySubscriptionCreateRequest(ctx context.Context, options *MetricAlertsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/metricAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MetricAlertsClient) listBySubscriptionHandleResponse(resp *http.Response) (MetricAlertsClientListBySubscriptionResponse, error) {
	result := MetricAlertsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResourceCollection); err != nil {
		return MetricAlertsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an metric alert definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - parameters - The parameters of the rule to update.
//   - options - MetricAlertsClientUpdateOptions contains the optional parameters for the MetricAlertsClient.Update method.
func (client *MetricAlertsClient) Update(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResourcePatch, options *MetricAlertsClientUpdateOptions) (MetricAlertsClientUpdateResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return MetricAlertsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *MetricAlertsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResourcePatch, options *MetricAlertsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MetricAlertsClient) updateHandleResponse(resp *http.Response) (MetricAlertsClientUpdateResponse, error) {
	result := MetricAlertsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResource); err != nil {
		return MetricAlertsClientUpdateResponse{}, err
	}
	return result, nil
}
