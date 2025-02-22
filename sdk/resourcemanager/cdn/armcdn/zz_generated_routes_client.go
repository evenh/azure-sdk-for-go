//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// RoutesClient contains the methods for the Routes group.
// Don't use this type directly, use NewRoutesClient() instead.
type RoutesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRoutesClient creates a new instance of RoutesClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoutesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RoutesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RoutesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Creates a new route with the specified route name under the specified subscription, resource group, profile,
// and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// routeName - Name of the routing rule.
// route - Route properties
// options - RoutesClientBeginCreateOptions contains the optional parameters for the RoutesClient.BeginCreate method.
func (client *RoutesClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, route Route, options *RoutesClientBeginCreateOptions) (RoutesClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, routeName, route, options)
	if err != nil {
		return RoutesClientCreatePollerResponse{}, err
	}
	result := RoutesClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoutesClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return RoutesClientCreatePollerResponse{}, err
	}
	result.Poller = &RoutesClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new route with the specified route name under the specified subscription, resource group, profile, and
// AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RoutesClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, route Route, options *RoutesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, route, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *RoutesClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, route Route, options *RoutesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, route)
}

// BeginDelete - Deletes an existing route with the specified route name under the specified subscription, resource group,
// profile, and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// routeName - Name of the routing rule.
// options - RoutesClientBeginDeleteOptions contains the optional parameters for the RoutesClient.BeginDelete method.
func (client *RoutesClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientBeginDeleteOptions) (RoutesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, routeName, options)
	if err != nil {
		return RoutesClientDeletePollerResponse{}, err
	}
	result := RoutesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoutesClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return RoutesClientDeletePollerResponse{}, err
	}
	result.Poller = &RoutesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing route with the specified route name under the specified subscription, resource group, profile,
// and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RoutesClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, options)
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
func (client *RoutesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an existing route with the specified route name under the specified subscription, resource group, profile, and
// AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// routeName - Name of the routing rule.
// options - RoutesClientGetOptions contains the optional parameters for the RoutesClient.Get method.
func (client *RoutesClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientGetOptions) (RoutesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, options)
	if err != nil {
		return RoutesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoutesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoutesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoutesClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *RoutesClient) getHandleResponse(resp *http.Response) (RoutesClientGetResponse, error) {
	result := RoutesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Route); err != nil {
		return RoutesClientGetResponse{}, err
	}
	return result, nil
}

// ListByEndpoint - Lists all of the existing origins within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// options - RoutesClientListByEndpointOptions contains the optional parameters for the RoutesClient.ListByEndpoint method.
func (client *RoutesClient) ListByEndpoint(resourceGroupName string, profileName string, endpointName string, options *RoutesClientListByEndpointOptions) *RoutesClientListByEndpointPager {
	return &RoutesClientListByEndpointPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
		},
		advancer: func(ctx context.Context, resp RoutesClientListByEndpointResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RouteListResult.NextLink)
		},
	}
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *RoutesClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *RoutesClientListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *RoutesClient) listByEndpointHandleResponse(resp *http.Response) (RoutesClientListByEndpointResponse, error) {
	result := RoutesClientListByEndpointResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteListResult); err != nil {
		return RoutesClientListByEndpointResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing route with the specified route name under the specified subscription, resource group,
// profile, and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// routeName - Name of the routing rule.
// routeUpdateProperties - Route update properties
// options - RoutesClientBeginUpdateOptions contains the optional parameters for the RoutesClient.BeginUpdate method.
func (client *RoutesClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, routeUpdateProperties RouteUpdateParameters, options *RoutesClientBeginUpdateOptions) (RoutesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, routeName, routeUpdateProperties, options)
	if err != nil {
		return RoutesClientUpdatePollerResponse{}, err
	}
	result := RoutesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoutesClient.Update", "azure-async-operation", resp, client.pl)
	if err != nil {
		return RoutesClientUpdatePollerResponse{}, err
	}
	result.Poller = &RoutesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an existing route with the specified route name under the specified subscription, resource group, profile,
// and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RoutesClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, routeUpdateProperties RouteUpdateParameters, options *RoutesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, routeUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *RoutesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, routeUpdateProperties RouteUpdateParameters, options *RoutesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, routeUpdateProperties)
}
