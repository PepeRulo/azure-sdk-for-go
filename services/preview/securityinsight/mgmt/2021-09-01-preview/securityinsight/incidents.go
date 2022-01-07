package securityinsight

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IncidentsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type IncidentsClient struct {
	BaseClient
}

// NewIncidentsClient creates an instance of the IncidentsClient client.
func NewIncidentsClient(subscriptionID string) IncidentsClient {
	return NewIncidentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIncidentsClientWithBaseURI creates an instance of the IncidentsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIncidentsClientWithBaseURI(baseURI string, subscriptionID string) IncidentsClient {
	return IncidentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the incident.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
// incident - the incident
func (client IncidentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident) (result Incident, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: incident,
			Constraints: []validation.Constraint{{Target: "incident.IncidentProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "incident.IncidentProperties.Title", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, incidentID, incident)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IncidentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}", pathParameters),
		autorest.WithJSON(incident),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IncidentsClient) CreateOrUpdateResponder(resp *http.Response) (result Incident, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateTeam creates a Microsoft team to investigate the incident by sharing information and insights between
// participants.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
// teamProperties - team properties
func (client IncidentsClient) CreateTeam(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, teamProperties TeamProperties) (result TeamInformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.CreateTeam")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: teamProperties,
			Constraints: []validation.Constraint{{Target: "teamProperties.TeamName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "CreateTeam", err.Error())
	}

	req, err := client.CreateTeamPreparer(ctx, resourceGroupName, workspaceName, incidentID, teamProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateTeam", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateTeamSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateTeam", resp, "Failure sending request")
		return
	}

	result, err = client.CreateTeamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateTeam", resp, "Failure responding to request")
		return
	}

	return
}

// CreateTeamPreparer prepares the CreateTeam request.
func (client IncidentsClient) CreateTeamPreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, teamProperties TeamProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/createTeam", pathParameters),
		autorest.WithJSON(teamProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateTeamSender sends the CreateTeam request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) CreateTeamSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateTeamResponder handles the response to the CreateTeam request. The method always
// closes the http.Response Body.
func (client IncidentsClient) CreateTeamResponder(resp *http.Response) (result TeamInformation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the incident.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
func (client IncidentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, incidentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IncidentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IncidentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an incident.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
func (client IncidentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result Incident, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, incidentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IncidentsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IncidentsClient) GetResponder(resp *http.Response) (result Incident, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all incidents.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// filter - filters the results, based on a Boolean condition. Optional.
// orderby - sorts the results. Optional.
// top - returns only the first n results. Optional.
// skipToken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls. Optional.
func (client IncidentsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result IncidentListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.List")
		defer func() {
			sc := -1
			if result.il.Response.Response != nil {
				sc = result.il.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, filter, orderby, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.il.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "List", resp, "Failure sending request")
		return
	}

	result.il, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.il.hasNextLink() && result.il.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IncidentsClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IncidentsClient) ListResponder(resp *http.Response) (result IncidentList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IncidentsClient) listNextResults(ctx context.Context, lastResults IncidentList) (result IncidentList, err error) {
	req, err := lastResults.incidentListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IncidentsClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result IncidentListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName, filter, orderby, top, skipToken)
	return
}

// ListAlerts gets all incident alerts.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
func (client IncidentsClient) ListAlerts(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result IncidentAlertList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.ListAlerts")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "ListAlerts", err.Error())
	}

	req, err := client.ListAlertsPreparer(ctx, resourceGroupName, workspaceName, incidentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListAlerts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAlertsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListAlerts", resp, "Failure sending request")
		return
	}

	result, err = client.ListAlertsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListAlerts", resp, "Failure responding to request")
		return
	}

	return
}

// ListAlertsPreparer prepares the ListAlerts request.
func (client IncidentsClient) ListAlertsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAlertsSender sends the ListAlerts request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) ListAlertsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAlertsResponder handles the response to the ListAlerts request. The method always
// closes the http.Response Body.
func (client IncidentsClient) ListAlertsResponder(resp *http.Response) (result IncidentAlertList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBookmarks gets all incident bookmarks.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
func (client IncidentsClient) ListBookmarks(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result IncidentBookmarkList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.ListBookmarks")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "ListBookmarks", err.Error())
	}

	req, err := client.ListBookmarksPreparer(ctx, resourceGroupName, workspaceName, incidentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListBookmarks", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBookmarksSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListBookmarks", resp, "Failure sending request")
		return
	}

	result, err = client.ListBookmarksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListBookmarks", resp, "Failure responding to request")
		return
	}

	return
}

// ListBookmarksPreparer prepares the ListBookmarks request.
func (client IncidentsClient) ListBookmarksPreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/bookmarks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBookmarksSender sends the ListBookmarks request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) ListBookmarksSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBookmarksResponder handles the response to the ListBookmarks request. The method always
// closes the http.Response Body.
func (client IncidentsClient) ListBookmarksResponder(resp *http.Response) (result IncidentBookmarkList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListEntities gets all incident related entities.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
func (client IncidentsClient) ListEntities(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result IncidentEntitiesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.ListEntities")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "ListEntities", err.Error())
	}

	req, err := client.ListEntitiesPreparer(ctx, resourceGroupName, workspaceName, incidentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListEntities", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListEntitiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListEntities", resp, "Failure sending request")
		return
	}

	result, err = client.ListEntitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "ListEntities", resp, "Failure responding to request")
		return
	}

	return
}

// ListEntitiesPreparer prepares the ListEntities request.
func (client IncidentsClient) ListEntitiesPreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/entities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListEntitiesSender sends the ListEntities request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) ListEntitiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListEntitiesResponder handles the response to the ListEntities request. The method always
// closes the http.Response Body.
func (client IncidentsClient) ListEntitiesResponder(resp *http.Response) (result IncidentEntitiesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
