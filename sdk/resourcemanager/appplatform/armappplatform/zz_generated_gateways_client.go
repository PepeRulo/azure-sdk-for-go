//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// GatewaysClient contains the methods for the Gateways group.
// Don't use this type directly, use NewGatewaysClient() instead.
type GatewaysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGatewaysClient creates a new instance of GatewaysClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GatewaysClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &GatewaysClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create the default Spring Cloud Gateway or update the existing Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// gatewayResource - The gateway for the create or update operation
// options - GatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the GatewaysClient.BeginCreateOrUpdate
// method.
func (client *GatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayResource GatewayResource, options *GatewaysClientBeginCreateOrUpdateOptions) (GatewaysClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, gatewayName, gatewayResource, options)
	if err != nil {
		return GatewaysClientCreateOrUpdatePollerResponse{}, err
	}
	result := GatewaysClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return GatewaysClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &GatewaysClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create the default Spring Cloud Gateway or update the existing Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayResource GatewayResource, options *GatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, gatewayResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayResource GatewayResource, options *GatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, gatewayResource)
}

// BeginDelete - Disable the default Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// options - GatewaysClientBeginDeleteOptions contains the optional parameters for the GatewaysClient.BeginDelete method.
func (client *GatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (GatewaysClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, gatewayName, options)
	if err != nil {
		return GatewaysClientDeletePollerResponse{}, err
	}
	result := GatewaysClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GatewaysClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return GatewaysClientDeletePollerResponse{}, err
	}
	result.Poller = &GatewaysClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Disable the default Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the Spring Cloud Gateway and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// options - GatewaysClientGetOptions contains the optional parameters for the GatewaysClient.Get method.
func (client *GatewaysClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientGetOptions) (GatewaysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
	if err != nil {
		return GatewaysClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewaysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewaysClient) getHandleResponse(resp *http.Response) (GatewaysClientGetResponse, error) {
	result := GatewaysClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResource); err != nil {
		return GatewaysClientGetResponse{}, err
	}
	return result, nil
}

// List - Handles requests to list all resources in a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - GatewaysClientListOptions contains the optional parameters for the GatewaysClient.List method.
func (client *GatewaysClient) List(resourceGroupName string, serviceName string, options *GatewaysClientListOptions) *GatewaysClientListPager {
	return &GatewaysClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp GatewaysClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GatewayResourceCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *GatewaysClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *GatewaysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GatewaysClient) listHandleResponse(resp *http.Response) (GatewaysClientListResponse, error) {
	result := GatewaysClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResourceCollection); err != nil {
		return GatewaysClientListResponse{}, err
	}
	return result, nil
}

// ValidateDomain - Check the domains are valid as well as not in use.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// validatePayload - Custom domain payload to be validated
// options - GatewaysClientValidateDomainOptions contains the optional parameters for the GatewaysClient.ValidateDomain method.
func (client *GatewaysClient) ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, validatePayload CustomDomainValidatePayload, options *GatewaysClientValidateDomainOptions) (GatewaysClientValidateDomainResponse, error) {
	req, err := client.validateDomainCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, validatePayload, options)
	if err != nil {
		return GatewaysClientValidateDomainResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewaysClientValidateDomainResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewaysClientValidateDomainResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateDomainHandleResponse(resp)
}

// validateDomainCreateRequest creates the ValidateDomain request.
func (client *GatewaysClient) validateDomainCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, validatePayload CustomDomainValidatePayload, options *GatewaysClientValidateDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/validateDomain"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, validatePayload)
}

// validateDomainHandleResponse handles the ValidateDomain response.
func (client *GatewaysClient) validateDomainHandleResponse(resp *http.Response) (GatewaysClientValidateDomainResponse, error) {
	result := GatewaysClientValidateDomainResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomainValidateResult); err != nil {
		return GatewaysClientValidateDomainResponse{}, err
	}
	return result, nil
}