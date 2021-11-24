//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

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

// PolicyExemptionsClient contains the methods for the PolicyExemptions group.
// Don't use this type directly, use NewPolicyExemptionsClient() instead.
type PolicyExemptionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPolicyExemptionsClient creates a new instance of PolicyExemptionsClient with the specified values.
func NewPolicyExemptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PolicyExemptionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PolicyExemptionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - This operation creates or updates a policy exemption with the given scope and name. Policy exemptions apply to all resources contained
// within their scope. For example, when you create a policy
// exemption at resource group scope for a policy assignment at the same or above level, the exemption exempts to all applicable resources in the resource
// group.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyExemptionsClient) CreateOrUpdate(ctx context.Context, scope string, policyExemptionName string, parameters PolicyExemption, options *PolicyExemptionsCreateOrUpdateOptions) (PolicyExemptionsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, policyExemptionName, parameters, options)
	if err != nil {
		return PolicyExemptionsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyExemptionsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PolicyExemptionsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PolicyExemptionsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, policyExemptionName string, parameters PolicyExemption, options *PolicyExemptionsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if policyExemptionName == "" {
		return nil, errors.New("parameter policyExemptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyExemptionName}", url.PathEscape(policyExemptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PolicyExemptionsClient) createOrUpdateHandleResponse(resp *http.Response) (PolicyExemptionsCreateOrUpdateResponse, error) {
	result := PolicyExemptionsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyExemption); err != nil {
		return PolicyExemptionsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PolicyExemptionsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - This operation deletes a policy exemption, given its name and the scope it was created in. The scope of a policy exemption is the part of its
// ID preceding
// '/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}'.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyExemptionsClient) Delete(ctx context.Context, scope string, policyExemptionName string, options *PolicyExemptionsDeleteOptions) (PolicyExemptionsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, policyExemptionName, options)
	if err != nil {
		return PolicyExemptionsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyExemptionsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PolicyExemptionsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return PolicyExemptionsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PolicyExemptionsClient) deleteCreateRequest(ctx context.Context, scope string, policyExemptionName string, options *PolicyExemptionsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if policyExemptionName == "" {
		return nil, errors.New("parameter policyExemptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyExemptionName}", url.PathEscape(policyExemptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PolicyExemptionsClient) deleteHandleError(resp *http.Response) error {
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

// Get - This operation retrieves a single policy exemption, given its name and the scope it was created at.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyExemptionsClient) Get(ctx context.Context, scope string, policyExemptionName string, options *PolicyExemptionsGetOptions) (PolicyExemptionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, policyExemptionName, options)
	if err != nil {
		return PolicyExemptionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyExemptionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyExemptionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PolicyExemptionsClient) getCreateRequest(ctx context.Context, scope string, policyExemptionName string, options *PolicyExemptionsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if policyExemptionName == "" {
		return nil, errors.New("parameter policyExemptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyExemptionName}", url.PathEscape(policyExemptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PolicyExemptionsClient) getHandleResponse(resp *http.Response) (PolicyExemptionsGetResponse, error) {
	result := PolicyExemptionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyExemption); err != nil {
		return PolicyExemptionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PolicyExemptionsClient) getHandleError(resp *http.Response) error {
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

// List - This operation retrieves the list of all policy exemptions associated with the given subscription that match the optional given $filter. Valid
// values for $filter are: 'atScope()', 'atExactScope()',
// 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter is not provided, the unfiltered list includes all policy exemptions associated with
// the subscription, including those that apply
// directly or from management groups that contain the given subscription, as well as any applied to objects contained within the subscription.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyExemptionsClient) List(options *PolicyExemptionsListOptions) *PolicyExemptionsListPager {
	return &PolicyExemptionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PolicyExemptionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyExemptionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PolicyExemptionsClient) listCreateRequest(ctx context.Context, options *PolicyExemptionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyExemptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PolicyExemptionsClient) listHandleResponse(resp *http.Response) (PolicyExemptionsListResponse, error) {
	result := PolicyExemptionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyExemptionListResult); err != nil {
		return PolicyExemptionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PolicyExemptionsClient) listHandleError(resp *http.Response) error {
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

// ListForManagementGroup - This operation retrieves the list of all policy exemptions applicable to the management group that match the given $filter.
// Valid values for $filter are: 'atScope()', 'atExactScope()',
// 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter=atScope() is provided, the returned list includes all policy exemptions that are
// assigned to the management group or the management
// group's ancestors.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyExemptionsClient) ListForManagementGroup(managementGroupID string, options *PolicyExemptionsListForManagementGroupOptions) *PolicyExemptionsListForManagementGroupPager {
	return &PolicyExemptionsListForManagementGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForManagementGroupCreateRequest(ctx, managementGroupID, options)
		},
		advancer: func(ctx context.Context, resp PolicyExemptionsListForManagementGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyExemptionListResult.NextLink)
		},
	}
}

// listForManagementGroupCreateRequest creates the ListForManagementGroup request.
func (client *PolicyExemptionsClient) listForManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *PolicyExemptionsListForManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyExemptions"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForManagementGroupHandleResponse handles the ListForManagementGroup response.
func (client *PolicyExemptionsClient) listForManagementGroupHandleResponse(resp *http.Response) (PolicyExemptionsListForManagementGroupResponse, error) {
	result := PolicyExemptionsListForManagementGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyExemptionListResult); err != nil {
		return PolicyExemptionsListForManagementGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listForManagementGroupHandleError handles the ListForManagementGroup error response.
func (client *PolicyExemptionsClient) listForManagementGroupHandleError(resp *http.Response) error {
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

// ListForResource - This operation retrieves the list of all policy exemptions associated with the specified resource in the given resource group and subscription
// that match the optional given $filter. Valid values for
// $filter are: 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter is not provided, the unfiltered list
// includes all policy exemptions associated with the
// resource, including those that apply directly or from all containing scopes, as well as any applied to resources contained within the resource. Three
// parameters plus the resource name are used to
// identify a specific resource. If the resource is not part of a parent resource (the more common case), the parent resource path should not be provided
// (or provided as ''). For example a web app could
// be specified as ({resourceProviderNamespace} == 'Microsoft.Web', {parentResourcePath} == '', {resourceType} == 'sites', {resourceName} == 'MyWebApp').
// If the resource is part of a parent resource,
// then all parameters should be provided. For example a virtual machine DNS name could be specified as ({resourceProviderNamespace} == 'Microsoft.Compute',
// {parentResourcePath} ==
// 'virtualMachines/MyVirtualMachine', {resourceType} == 'domainNames', {resourceName} == 'MyComputerName'). A convenient alternative to providing the namespace
// and type name separately is to provide
// both in the {resourceType} parameter, format: ({resourceProviderNamespace} == '', {parentResourcePath} == '', {resourceType} == 'Microsoft.Web/sites',
// {resourceName} == 'MyWebApp').
// If the operation fails it returns the *CloudError error type.
func (client *PolicyExemptionsClient) ListForResource(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *PolicyExemptionsListForResourceOptions) *PolicyExemptionsListForResourcePager {
	return &PolicyExemptionsListForResourcePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, options)
		},
		advancer: func(ctx context.Context, resp PolicyExemptionsListForResourceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyExemptionListResult.NextLink)
		},
	}
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *PolicyExemptionsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *PolicyExemptionsListForResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/policyExemptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourcePath}", parentResourcePath)
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *PolicyExemptionsClient) listForResourceHandleResponse(resp *http.Response) (PolicyExemptionsListForResourceResponse, error) {
	result := PolicyExemptionsListForResourceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyExemptionListResult); err != nil {
		return PolicyExemptionsListForResourceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listForResourceHandleError handles the ListForResource error response.
func (client *PolicyExemptionsClient) listForResourceHandleError(resp *http.Response) error {
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

// ListForResourceGroup - This operation retrieves the list of all policy exemptions associated with the given resource group in the given subscription
// that match the optional given $filter. Valid values for $filter are:
// 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter is not provided, the unfiltered list includes all
// policy exemptions associated with the resource
// group, including those that apply directly or apply from containing scopes, as well as any applied to resources contained within the resource group.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyExemptionsClient) ListForResourceGroup(resourceGroupName string, options *PolicyExemptionsListForResourceGroupOptions) *PolicyExemptionsListForResourceGroupPager {
	return &PolicyExemptionsListForResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PolicyExemptionsListForResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyExemptionListResult.NextLink)
		},
	}
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *PolicyExemptionsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PolicyExemptionsListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyExemptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *PolicyExemptionsClient) listForResourceGroupHandleResponse(resp *http.Response) (PolicyExemptionsListForResourceGroupResponse, error) {
	result := PolicyExemptionsListForResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyExemptionListResult); err != nil {
		return PolicyExemptionsListForResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listForResourceGroupHandleError handles the ListForResourceGroup error response.
func (client *PolicyExemptionsClient) listForResourceGroupHandleError(resp *http.Response) error {
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