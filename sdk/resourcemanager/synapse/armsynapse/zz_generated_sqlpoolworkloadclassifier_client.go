//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SQLPoolWorkloadClassifierClient contains the methods for the SQLPoolWorkloadClassifier group.
// Don't use this type directly, use NewSQLPoolWorkloadClassifierClient() instead.
type SQLPoolWorkloadClassifierClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSQLPoolWorkloadClassifierClient creates a new instance of SQLPoolWorkloadClassifierClient with the specified values.
func NewSQLPoolWorkloadClassifierClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SQLPoolWorkloadClassifierClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SQLPoolWorkloadClassifierClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Create Or Update workload classifier for a Sql pool's workload group.
// If the operation fails it returns a generic error.
func (client *SQLPoolWorkloadClassifierClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *SQLPoolWorkloadClassifierBeginCreateOrUpdateOptions) (SQLPoolWorkloadClassifierCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, workloadClassifierName, parameters, options)
	if err != nil {
		return SQLPoolWorkloadClassifierCreateOrUpdatePollerResponse{}, err
	}
	result := SQLPoolWorkloadClassifierCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLPoolWorkloadClassifierClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SQLPoolWorkloadClassifierCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SQLPoolWorkloadClassifierCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create Or Update workload classifier for a Sql pool's workload group.
// If the operation fails it returns a generic error.
func (client *SQLPoolWorkloadClassifierClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *SQLPoolWorkloadClassifierBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, workloadClassifierName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLPoolWorkloadClassifierClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *SQLPoolWorkloadClassifierBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SQLPoolWorkloadClassifierClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Remove workload classifier of a Sql pool's workload group.
// If the operation fails it returns a generic error.
func (client *SQLPoolWorkloadClassifierClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, options *SQLPoolWorkloadClassifierBeginDeleteOptions) (SQLPoolWorkloadClassifierDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, workloadClassifierName, options)
	if err != nil {
		return SQLPoolWorkloadClassifierDeletePollerResponse{}, err
	}
	result := SQLPoolWorkloadClassifierDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLPoolWorkloadClassifierClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SQLPoolWorkloadClassifierDeletePollerResponse{}, err
	}
	result.Poller = &SQLPoolWorkloadClassifierDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Remove workload classifier of a Sql pool's workload group.
// If the operation fails it returns a generic error.
func (client *SQLPoolWorkloadClassifierClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, options *SQLPoolWorkloadClassifierBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, workloadClassifierName, options)
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
func (client *SQLPoolWorkloadClassifierClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, options *SQLPoolWorkloadClassifierBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SQLPoolWorkloadClassifierClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get a workload classifier of Sql pool's workload group.
// If the operation fails it returns a generic error.
func (client *SQLPoolWorkloadClassifierClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, options *SQLPoolWorkloadClassifierGetOptions) (SQLPoolWorkloadClassifierGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, workloadClassifierName, options)
	if err != nil {
		return SQLPoolWorkloadClassifierGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolWorkloadClassifierGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolWorkloadClassifierGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolWorkloadClassifierClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, workloadClassifierName string, options *SQLPoolWorkloadClassifierGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLPoolWorkloadClassifierClient) getHandleResponse(resp *http.Response) (SQLPoolWorkloadClassifierGetResponse, error) {
	result := SQLPoolWorkloadClassifierGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadClassifier); err != nil {
		return SQLPoolWorkloadClassifierGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SQLPoolWorkloadClassifierClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Get list of Sql pool's workload classifier for workload groups.
// If the operation fails it returns a generic error.
func (client *SQLPoolWorkloadClassifierClient) List(resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, options *SQLPoolWorkloadClassifierListOptions) *SQLPoolWorkloadClassifierListPager {
	return &SQLPoolWorkloadClassifierListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, options)
		},
		advancer: func(ctx context.Context, resp SQLPoolWorkloadClassifierListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkloadClassifierListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SQLPoolWorkloadClassifierClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, options *SQLPoolWorkloadClassifierListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLPoolWorkloadClassifierClient) listHandleResponse(resp *http.Response) (SQLPoolWorkloadClassifierListResponse, error) {
	result := SQLPoolWorkloadClassifierListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadClassifierListResult); err != nil {
		return SQLPoolWorkloadClassifierListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SQLPoolWorkloadClassifierClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}