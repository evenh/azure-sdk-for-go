package labservices

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

// EnvironmentSettingsClient is the the Managed Labs Client.
type EnvironmentSettingsClient struct {
	BaseClient
}

// NewEnvironmentSettingsClient creates an instance of the EnvironmentSettingsClient client.
func NewEnvironmentSettingsClient(subscriptionID string) EnvironmentSettingsClient {
	return NewEnvironmentSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEnvironmentSettingsClientWithBaseURI creates an instance of the EnvironmentSettingsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewEnvironmentSettingsClientWithBaseURI(baseURI string, subscriptionID string) EnvironmentSettingsClient {
	return EnvironmentSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ClaimAny claims a random environment for a user in an environment settings
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
func (client EnvironmentSettingsClient) ClaimAny(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.ClaimAny")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ClaimAnyPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "ClaimAny", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClaimAnySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "ClaimAny", resp, "Failure sending request")
		return
	}

	result, err = client.ClaimAnyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "ClaimAny", resp, "Failure responding to request")
		return
	}

	return
}

// ClaimAnyPreparer prepares the ClaimAny request.
func (client EnvironmentSettingsClient) ClaimAnyPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/claimAny", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ClaimAnySender sends the ClaimAny request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) ClaimAnySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ClaimAnyResponder handles the response to the ClaimAny request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) ClaimAnyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create or replace an existing Environment Setting. This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentSetting - represents settings of an environment, from which environment instances would be
// created
func (client EnvironmentSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentSetting EnvironmentSetting) (result EnvironmentSettingsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: environmentSetting,
			Constraints: []validation.Constraint{{Target: "environmentSetting.EnvironmentSettingProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "environmentSetting.EnvironmentSettingProperties.ResourceSettings", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "environmentSetting.EnvironmentSettingProperties.ResourceSettings.ReferenceVM", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "environmentSetting.EnvironmentSettingProperties.ResourceSettings.ReferenceVM.UserName", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("labservices.EnvironmentSettingsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentSetting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EnvironmentSettingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentSetting EnvironmentSetting) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", pathParameters),
		autorest.WithJSON(environmentSetting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) CreateOrUpdateSender(req *http.Request) (future EnvironmentSettingsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result EnvironmentSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete environment setting. This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
func (client EnvironmentSettingsClient) Delete(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result EnvironmentSettingsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EnvironmentSettingsClient) DeletePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) DeleteSender(req *http.Request) (future EnvironmentSettingsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get environment setting
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// expand - specify the $expand query. Example: 'properties($select=publishingState)'
func (client EnvironmentSettingsClient) Get(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, expand string) (result EnvironmentSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client EnvironmentSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) GetResponder(resp *http.Response) (result EnvironmentSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list environment setting in a given lab.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// expand - specify the $expand query. Example: 'properties($select=publishingState)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client EnvironmentSettingsClient) List(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationEnvironmentSettingPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.List")
		defer func() {
			sc := -1
			if result.rwces.Response.Response != nil {
				sc = result.rwces.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labAccountName, labName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwces.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "List", resp, "Failure sending request")
		return
	}

	result.rwces, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rwces.hasNextLink() && result.rwces.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EnvironmentSettingsClient) ListPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationEnvironmentSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EnvironmentSettingsClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationEnvironmentSetting) (result ResponseWithContinuationEnvironmentSetting, err error) {
	req, err := lastResults.responseWithContinuationEnvironmentSettingPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EnvironmentSettingsClient) ListComplete(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationEnvironmentSettingIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labAccountName, labName, expand, filter, top, orderby)
	return
}

// Publish provisions/deprovisions required resources for an environment setting based on current state of the
// lab/environment setting.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// publishPayload - payload for Publish operation on EnvironmentSetting.
func (client EnvironmentSettingsClient) Publish(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, publishPayload PublishPayload) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.Publish")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PublishPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, publishPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Publish", nil, "Failure preparing request")
		return
	}

	resp, err := client.PublishSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Publish", resp, "Failure sending request")
		return
	}

	result, err = client.PublishResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Publish", resp, "Failure responding to request")
		return
	}

	return
}

// PublishPreparer prepares the Publish request.
func (client EnvironmentSettingsClient) PublishPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, publishPayload PublishPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/publish", pathParameters),
		autorest.WithJSON(publishPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishSender sends the Publish request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) PublishSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PublishResponder handles the response to the Publish request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) PublishResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start starts a template by starting all resources inside the template. This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
func (client EnvironmentSettingsClient) Start(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result EnvironmentSettingsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.Start")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Start", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client EnvironmentSettingsClient) StartPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) StartSender(req *http.Request) (future EnvironmentSettingsStartFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop starts a template by starting all resources inside the template. This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
func (client EnvironmentSettingsClient) Stop(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result EnvironmentSettingsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.Stop")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client EnvironmentSettingsClient) StopPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) StopSender(req *http.Request) (future EnvironmentSettingsStopFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update modify properties of environment setting.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentSetting - represents settings of an environment, from which environment instances would be
// created
func (client EnvironmentSettingsClient) Update(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentSetting EnvironmentSettingFragment) (result EnvironmentSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentSettingsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentSetting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentSettingsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EnvironmentSettingsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentSetting EnvironmentSettingFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", pathParameters),
		autorest.WithJSON(environmentSetting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentSettingsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EnvironmentSettingsClient) UpdateResponder(resp *http.Response) (result EnvironmentSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
