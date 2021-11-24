//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

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

// ScriptActionsClient contains the methods for the ScriptActions group.
// Don't use this type directly, use NewScriptActionsClient() instead.
type ScriptActionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewScriptActionsClient creates a new instance of ScriptActionsClient with the specified values.
func NewScriptActionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ScriptActionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ScriptActionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Delete - Deletes a specified persisted script action of the cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ScriptActionsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, scriptName string, options *ScriptActionsDeleteOptions) (ScriptActionsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, scriptName, options)
	if err != nil {
		return ScriptActionsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptActionsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ScriptActionsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ScriptActionsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ScriptActionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, scriptName string, options *ScriptActionsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptActions/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ScriptActionsClient) deleteHandleError(resp *http.Response) error {
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

// GetExecutionAsyncOperationStatus - Gets the async operation status of execution operation.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ScriptActionsClient) GetExecutionAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, operationID string, options *ScriptActionsGetExecutionAsyncOperationStatusOptions) (ScriptActionsGetExecutionAsyncOperationStatusResponse, error) {
	req, err := client.getExecutionAsyncOperationStatusCreateRequest(ctx, resourceGroupName, clusterName, operationID, options)
	if err != nil {
		return ScriptActionsGetExecutionAsyncOperationStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptActionsGetExecutionAsyncOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScriptActionsGetExecutionAsyncOperationStatusResponse{}, client.getExecutionAsyncOperationStatusHandleError(resp)
	}
	return client.getExecutionAsyncOperationStatusHandleResponse(resp)
}

// getExecutionAsyncOperationStatusCreateRequest creates the GetExecutionAsyncOperationStatus request.
func (client *ScriptActionsClient) getExecutionAsyncOperationStatusCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, operationID string, options *ScriptActionsGetExecutionAsyncOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/executeScriptActions/azureasyncoperations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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

// getExecutionAsyncOperationStatusHandleResponse handles the GetExecutionAsyncOperationStatus response.
func (client *ScriptActionsClient) getExecutionAsyncOperationStatusHandleResponse(resp *http.Response) (ScriptActionsGetExecutionAsyncOperationStatusResponse, error) {
	result := ScriptActionsGetExecutionAsyncOperationStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AsyncOperationResult); err != nil {
		return ScriptActionsGetExecutionAsyncOperationStatusResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getExecutionAsyncOperationStatusHandleError handles the GetExecutionAsyncOperationStatus error response.
func (client *ScriptActionsClient) getExecutionAsyncOperationStatusHandleError(resp *http.Response) error {
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

// GetExecutionDetail - Gets the script execution detail for the given script execution ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ScriptActionsClient) GetExecutionDetail(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string, options *ScriptActionsGetExecutionDetailOptions) (ScriptActionsGetExecutionDetailResponse, error) {
	req, err := client.getExecutionDetailCreateRequest(ctx, resourceGroupName, clusterName, scriptExecutionID, options)
	if err != nil {
		return ScriptActionsGetExecutionDetailResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptActionsGetExecutionDetailResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScriptActionsGetExecutionDetailResponse{}, client.getExecutionDetailHandleError(resp)
	}
	return client.getExecutionDetailHandleResponse(resp)
}

// getExecutionDetailCreateRequest creates the GetExecutionDetail request.
func (client *ScriptActionsClient) getExecutionDetailCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string, options *ScriptActionsGetExecutionDetailOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory/{scriptExecutionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if scriptExecutionID == "" {
		return nil, errors.New("parameter scriptExecutionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionId}", url.PathEscape(scriptExecutionID))
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

// getExecutionDetailHandleResponse handles the GetExecutionDetail response.
func (client *ScriptActionsClient) getExecutionDetailHandleResponse(resp *http.Response) (ScriptActionsGetExecutionDetailResponse, error) {
	result := ScriptActionsGetExecutionDetailResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuntimeScriptActionDetail); err != nil {
		return ScriptActionsGetExecutionDetailResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getExecutionDetailHandleError handles the GetExecutionDetail error response.
func (client *ScriptActionsClient) getExecutionDetailHandleError(resp *http.Response) error {
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

// ListByCluster - Lists all the persisted script actions for the specified cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ScriptActionsClient) ListByCluster(resourceGroupName string, clusterName string, options *ScriptActionsListByClusterOptions) *ScriptActionsListByClusterPager {
	return &ScriptActionsListByClusterPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
		},
		advancer: func(ctx context.Context, resp ScriptActionsListByClusterResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ScriptActionsList.NextLink)
		},
	}
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *ScriptActionsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ScriptActionsListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptActions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
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

// listByClusterHandleResponse handles the ListByCluster response.
func (client *ScriptActionsClient) listByClusterHandleResponse(resp *http.Response) (ScriptActionsListByClusterResponse, error) {
	result := ScriptActionsListByClusterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptActionsList); err != nil {
		return ScriptActionsListByClusterResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByClusterHandleError handles the ListByCluster error response.
func (client *ScriptActionsClient) listByClusterHandleError(resp *http.Response) error {
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