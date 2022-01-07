//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomproviders

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

// AssociationsClient contains the methods for the Associations group.
// Don't use this type directly, use NewAssociationsClient() instead.
type AssociationsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewAssociationsClient creates a new instance of AssociationsClient with the specified values.
func NewAssociationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AssociationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AssociationsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Create or update an association.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssociationsClient) BeginCreateOrUpdate(ctx context.Context, scope string, associationName string, association Association, options *AssociationsBeginCreateOrUpdateOptions) (AssociationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, scope, associationName, association, options)
	if err != nil {
		return AssociationsCreateOrUpdatePollerResponse{}, err
	}
	result := AssociationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AssociationsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return AssociationsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &AssociationsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update an association.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssociationsClient) createOrUpdate(ctx context.Context, scope string, associationName string, association Association, options *AssociationsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, associationName, association, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssociationsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, associationName string, association Association, options *AssociationsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, association)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AssociationsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Delete an association.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssociationsClient) BeginDelete(ctx context.Context, scope string, associationName string, options *AssociationsBeginDeleteOptions) (AssociationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, scope, associationName, options)
	if err != nil {
		return AssociationsDeletePollerResponse{}, err
	}
	result := AssociationsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AssociationsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return AssociationsDeletePollerResponse{}, err
	}
	result.Poller = &AssociationsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete an association.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssociationsClient) deleteOperation(ctx context.Context, scope string, associationName string, options *AssociationsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, scope, associationName, options)
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
func (client *AssociationsClient) deleteCreateRequest(ctx context.Context, scope string, associationName string, options *AssociationsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AssociationsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get an association.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssociationsClient) Get(ctx context.Context, scope string, associationName string, options *AssociationsGetOptions) (AssociationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, associationName, options)
	if err != nil {
		return AssociationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssociationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssociationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssociationsClient) getCreateRequest(ctx context.Context, scope string, associationName string, options *AssociationsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssociationsClient) getHandleResponse(resp *http.Response) (AssociationsGetResponse, error) {
	result := AssociationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Association); err != nil {
		return AssociationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AssociationsClient) getHandleError(resp *http.Response) error {
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

// ListAll - Gets all association for the given scope.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssociationsClient) ListAll(scope string, options *AssociationsListAllOptions) *AssociationsListAllPager {
	return &AssociationsListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp AssociationsListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AssociationsList.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *AssociationsClient) listAllCreateRequest(ctx context.Context, scope string, options *AssociationsListAllOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *AssociationsClient) listAllHandleResponse(resp *http.Response) (AssociationsListAllResponse, error) {
	result := AssociationsListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssociationsList); err != nil {
		return AssociationsListAllResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *AssociationsClient) listAllHandleError(resp *http.Response) error {
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
