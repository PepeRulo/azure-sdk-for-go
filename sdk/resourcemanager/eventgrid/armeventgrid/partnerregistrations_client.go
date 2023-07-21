//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PartnerRegistrationsClient contains the methods for the PartnerRegistrations group.
// Don't use this type directly, use NewPartnerRegistrationsClient() instead.
type PartnerRegistrationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPartnerRegistrationsClient creates a new instance of PartnerRegistrationsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPartnerRegistrationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PartnerRegistrationsClient, error) {
	cl, err := arm.NewClient(moduleName+".PartnerRegistrationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PartnerRegistrationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new partner registration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerRegistrationName - Name of the partner registration.
//   - partnerRegistrationInfo - PartnerRegistration information.
//   - options - PartnerRegistrationsClientBeginCreateOrUpdateOptions contains the optional parameters for the PartnerRegistrationsClient.BeginCreateOrUpdate
//     method.
func (client *PartnerRegistrationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationInfo PartnerRegistration, options *PartnerRegistrationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PartnerRegistrationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, partnerRegistrationName, partnerRegistrationInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PartnerRegistrationsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PartnerRegistrationsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new partner registration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *PartnerRegistrationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationInfo PartnerRegistration, options *PartnerRegistrationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerRegistrationName, partnerRegistrationInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PartnerRegistrationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationInfo PartnerRegistration, options *PartnerRegistrationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerRegistrationName == "" {
		return nil, errors.New("parameter partnerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerRegistrationName}", url.PathEscape(partnerRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, partnerRegistrationInfo)
}

// BeginDelete - Deletes a partner registration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerRegistrationName - Name of the partner registration.
//   - options - PartnerRegistrationsClientBeginDeleteOptions contains the optional parameters for the PartnerRegistrationsClient.BeginDelete
//     method.
func (client *PartnerRegistrationsClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerRegistrationName string, options *PartnerRegistrationsClientBeginDeleteOptions) (*runtime.Poller[PartnerRegistrationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, partnerRegistrationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PartnerRegistrationsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PartnerRegistrationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a partner registration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *PartnerRegistrationsClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerRegistrationName string, options *PartnerRegistrationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerRegistrationName, options)
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
func (client *PartnerRegistrationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerRegistrationName string, options *PartnerRegistrationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerRegistrationName == "" {
		return nil, errors.New("parameter partnerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerRegistrationName}", url.PathEscape(partnerRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a partner registration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerRegistrationName - Name of the partner registration.
//   - options - PartnerRegistrationsClientGetOptions contains the optional parameters for the PartnerRegistrationsClient.Get
//     method.
func (client *PartnerRegistrationsClient) Get(ctx context.Context, resourceGroupName string, partnerRegistrationName string, options *PartnerRegistrationsClientGetOptions) (PartnerRegistrationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerRegistrationName, options)
	if err != nil {
		return PartnerRegistrationsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerRegistrationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerRegistrationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PartnerRegistrationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerRegistrationName string, options *PartnerRegistrationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerRegistrationName == "" {
		return nil, errors.New("parameter partnerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerRegistrationName}", url.PathEscape(partnerRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerRegistrationsClient) getHandleResponse(resp *http.Response) (PartnerRegistrationsClientGetResponse, error) {
	result := PartnerRegistrationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerRegistration); err != nil {
		return PartnerRegistrationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the partner registrations under a resource group.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - options - PartnerRegistrationsClientListByResourceGroupOptions contains the optional parameters for the PartnerRegistrationsClient.NewListByResourceGroupPager
//     method.
func (client *PartnerRegistrationsClient) NewListByResourceGroupPager(resourceGroupName string, options *PartnerRegistrationsClientListByResourceGroupOptions) *runtime.Pager[PartnerRegistrationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerRegistrationsClientListByResourceGroupResponse]{
		More: func(page PartnerRegistrationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerRegistrationsClientListByResourceGroupResponse) (PartnerRegistrationsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PartnerRegistrationsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PartnerRegistrationsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PartnerRegistrationsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PartnerRegistrationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PartnerRegistrationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations"
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
	reqQP.Set("api-version", "2023-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PartnerRegistrationsClient) listByResourceGroupHandleResponse(resp *http.Response) (PartnerRegistrationsClientListByResourceGroupResponse, error) {
	result := PartnerRegistrationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerRegistrationsListResult); err != nil {
		return PartnerRegistrationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the partner registrations under an Azure subscription.
//
// Generated from API version 2023-06-01-preview
//   - options - PartnerRegistrationsClientListBySubscriptionOptions contains the optional parameters for the PartnerRegistrationsClient.NewListBySubscriptionPager
//     method.
func (client *PartnerRegistrationsClient) NewListBySubscriptionPager(options *PartnerRegistrationsClientListBySubscriptionOptions) *runtime.Pager[PartnerRegistrationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerRegistrationsClientListBySubscriptionResponse]{
		More: func(page PartnerRegistrationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerRegistrationsClientListBySubscriptionResponse) (PartnerRegistrationsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PartnerRegistrationsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PartnerRegistrationsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PartnerRegistrationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PartnerRegistrationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *PartnerRegistrationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/partnerRegistrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PartnerRegistrationsClient) listBySubscriptionHandleResponse(resp *http.Response) (PartnerRegistrationsClientListBySubscriptionResponse, error) {
	result := PartnerRegistrationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerRegistrationsListResult); err != nil {
		return PartnerRegistrationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a partner registration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerRegistrationName - Name of the partner registration.
//   - partnerRegistrationUpdateParameters - Partner registration update information.
//   - options - PartnerRegistrationsClientBeginUpdateOptions contains the optional parameters for the PartnerRegistrationsClient.BeginUpdate
//     method.
func (client *PartnerRegistrationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationUpdateParameters PartnerRegistrationUpdateParameters, options *PartnerRegistrationsClientBeginUpdateOptions) (*runtime.Poller[PartnerRegistrationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, partnerRegistrationName, partnerRegistrationUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PartnerRegistrationsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PartnerRegistrationsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a partner registration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *PartnerRegistrationsClient) update(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationUpdateParameters PartnerRegistrationUpdateParameters, options *PartnerRegistrationsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, partnerRegistrationName, partnerRegistrationUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *PartnerRegistrationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationUpdateParameters PartnerRegistrationUpdateParameters, options *PartnerRegistrationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerRegistrationName == "" {
		return nil, errors.New("parameter partnerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerRegistrationName}", url.PathEscape(partnerRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, partnerRegistrationUpdateParameters)
}
