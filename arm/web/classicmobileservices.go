package web

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ClassicMobileServicesClient is the use these APIs to manage Azure Websites
// resources through the Azure Resource Manager. All task operations conform
// to the HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type ClassicMobileServicesClient struct {
	ManagementClient
}

// NewClassicMobileServicesClient creates an instance of the
// ClassicMobileServicesClient client.
func NewClassicMobileServicesClient(subscriptionID string) ClassicMobileServicesClient {
	return NewClassicMobileServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClassicMobileServicesClientWithBaseURI creates an instance of the
// ClassicMobileServicesClient client.
func NewClassicMobileServicesClientWithBaseURI(baseURI string, subscriptionID string) ClassicMobileServicesClient {
	return ClassicMobileServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DeleteClassicMobileService sends the delete classic mobile service request.
//
// resourceGroupName is name of resource group name is name of mobile service
func (client ClassicMobileServicesClient) DeleteClassicMobileService(resourceGroupName string, name string) (result SetObject, err error) {
	req, err := client.DeleteClassicMobileServicePreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "DeleteClassicMobileService", nil, "Failure preparing request")
	}

	resp, err := client.DeleteClassicMobileServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "DeleteClassicMobileService", resp, "Failure sending request")
	}

	result, err = client.DeleteClassicMobileServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "DeleteClassicMobileService", resp, "Failure responding to request")
	}

	return
}

// DeleteClassicMobileServicePreparer prepares the DeleteClassicMobileService request.
func (client ClassicMobileServicesClient) DeleteClassicMobileServicePreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/classicMobileServices/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteClassicMobileServiceSender sends the DeleteClassicMobileService request. The method will close the
// http.Response Body if it receives an error.
func (client ClassicMobileServicesClient) DeleteClassicMobileServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteClassicMobileServiceResponder handles the response to the DeleteClassicMobileService request. The method always
// closes the http.Response Body.
func (client ClassicMobileServicesClient) DeleteClassicMobileServiceResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetClassicMobileService sends the get classic mobile service request.
//
// resourceGroupName is name of resource group name is name of mobile service
func (client ClassicMobileServicesClient) GetClassicMobileService(resourceGroupName string, name string) (result ClassicMobileService, err error) {
	req, err := client.GetClassicMobileServicePreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileService", nil, "Failure preparing request")
	}

	resp, err := client.GetClassicMobileServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileService", resp, "Failure sending request")
	}

	result, err = client.GetClassicMobileServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileService", resp, "Failure responding to request")
	}

	return
}

// GetClassicMobileServicePreparer prepares the GetClassicMobileService request.
func (client ClassicMobileServicesClient) GetClassicMobileServicePreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/classicMobileServices/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetClassicMobileServiceSender sends the GetClassicMobileService request. The method will close the
// http.Response Body if it receives an error.
func (client ClassicMobileServicesClient) GetClassicMobileServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetClassicMobileServiceResponder handles the response to the GetClassicMobileService request. The method always
// closes the http.Response Body.
func (client ClassicMobileServicesClient) GetClassicMobileServiceResponder(resp *http.Response) (result ClassicMobileService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetClassicMobileServices sends the get classic mobile services request.
//
// resourceGroupName is name of resource group
func (client ClassicMobileServicesClient) GetClassicMobileServices(resourceGroupName string) (result ClassicMobileServiceCollection, err error) {
	req, err := client.GetClassicMobileServicesPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileServices", nil, "Failure preparing request")
	}

	resp, err := client.GetClassicMobileServicesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileServices", resp, "Failure sending request")
	}

	result, err = client.GetClassicMobileServicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileServices", resp, "Failure responding to request")
	}

	return
}

// GetClassicMobileServicesPreparer prepares the GetClassicMobileServices request.
func (client ClassicMobileServicesClient) GetClassicMobileServicesPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/classicMobileServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetClassicMobileServicesSender sends the GetClassicMobileServices request. The method will close the
// http.Response Body if it receives an error.
func (client ClassicMobileServicesClient) GetClassicMobileServicesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetClassicMobileServicesResponder handles the response to the GetClassicMobileServices request. The method always
// closes the http.Response Body.
func (client ClassicMobileServicesClient) GetClassicMobileServicesResponder(resp *http.Response) (result ClassicMobileServiceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
