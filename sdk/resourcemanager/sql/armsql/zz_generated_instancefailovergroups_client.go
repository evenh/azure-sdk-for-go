//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// InstanceFailoverGroupsClient contains the methods for the InstanceFailoverGroups group.
// Don't use this type directly, use NewInstanceFailoverGroupsClient() instead.
type InstanceFailoverGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInstanceFailoverGroupsClient creates a new instance of InstanceFailoverGroupsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInstanceFailoverGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *InstanceFailoverGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &InstanceFailoverGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// locationName - The name of the region where the resource is located.
// failoverGroupName - The name of the failover group.
// parameters - The failover group parameters.
// options - InstanceFailoverGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the InstanceFailoverGroupsClient.BeginCreateOrUpdate
// method.
func (client *InstanceFailoverGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup, options *InstanceFailoverGroupsClientBeginCreateOrUpdateOptions) (InstanceFailoverGroupsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, locationName, failoverGroupName, parameters, options)
	if err != nil {
		return InstanceFailoverGroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result := InstanceFailoverGroupsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return InstanceFailoverGroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstanceFailoverGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup, options *InstanceFailoverGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InstanceFailoverGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup, options *InstanceFailoverGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// locationName - The name of the region where the resource is located.
// failoverGroupName - The name of the failover group.
// options - InstanceFailoverGroupsClientBeginDeleteOptions contains the optional parameters for the InstanceFailoverGroupsClient.BeginDelete
// method.
func (client *InstanceFailoverGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginDeleteOptions) (InstanceFailoverGroupsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsClientDeletePollerResponse{}, err
	}
	result := InstanceFailoverGroupsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.Delete", "", resp, client.pl)
	if err != nil {
		return InstanceFailoverGroupsClientDeletePollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstanceFailoverGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
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
func (client *InstanceFailoverGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginFailover - Fails over from the current primary managed instance to this managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// locationName - The name of the region where the resource is located.
// failoverGroupName - The name of the failover group.
// options - InstanceFailoverGroupsClientBeginFailoverOptions contains the optional parameters for the InstanceFailoverGroupsClient.BeginFailover
// method.
func (client *InstanceFailoverGroupsClient) BeginFailover(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginFailoverOptions) (InstanceFailoverGroupsClientFailoverPollerResponse, error) {
	resp, err := client.failover(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsClientFailoverPollerResponse{}, err
	}
	result := InstanceFailoverGroupsClientFailoverPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.Failover", "", resp, client.pl)
	if err != nil {
		return InstanceFailoverGroupsClientFailoverPollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsClientFailoverPoller{
		pt: pt,
	}
	return result, nil
}

// Failover - Fails over from the current primary managed instance to this managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstanceFailoverGroupsClient) failover(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginFailoverOptions) (*http.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
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

// failoverCreateRequest creates the Failover request.
func (client *InstanceFailoverGroupsClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginForceFailoverAllowDataLoss - Fails over from the current primary managed instance to this managed instance. This operation
// might result in data loss.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// locationName - The name of the region where the resource is located.
// failoverGroupName - The name of the failover group.
// options - InstanceFailoverGroupsClientBeginForceFailoverAllowDataLossOptions contains the optional parameters for the InstanceFailoverGroupsClient.BeginForceFailoverAllowDataLoss
// method.
func (client *InstanceFailoverGroupsClient) BeginForceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginForceFailoverAllowDataLossOptions) (InstanceFailoverGroupsClientForceFailoverAllowDataLossPollerResponse, error) {
	resp, err := client.forceFailoverAllowDataLoss(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsClientForceFailoverAllowDataLossPollerResponse{}, err
	}
	result := InstanceFailoverGroupsClientForceFailoverAllowDataLossPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.ForceFailoverAllowDataLoss", "", resp, client.pl)
	if err != nil {
		return InstanceFailoverGroupsClientForceFailoverAllowDataLossPollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsClientForceFailoverAllowDataLossPoller{
		pt: pt,
	}
	return result, nil
}

// ForceFailoverAllowDataLoss - Fails over from the current primary managed instance to this managed instance. This operation
// might result in data loss.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstanceFailoverGroupsClient) forceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginForceFailoverAllowDataLossOptions) (*http.Response, error) {
	req, err := client.forceFailoverAllowDataLossCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
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

// forceFailoverAllowDataLossCreateRequest creates the ForceFailoverAllowDataLoss request.
func (client *InstanceFailoverGroupsClient) forceFailoverAllowDataLossCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientBeginForceFailoverAllowDataLossOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}/forceFailoverAllowDataLoss"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// locationName - The name of the region where the resource is located.
// failoverGroupName - The name of the failover group.
// options - InstanceFailoverGroupsClientGetOptions contains the optional parameters for the InstanceFailoverGroupsClient.Get
// method.
func (client *InstanceFailoverGroupsClient) Get(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientGetOptions) (InstanceFailoverGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstanceFailoverGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstanceFailoverGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InstanceFailoverGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InstanceFailoverGroupsClient) getHandleResponse(resp *http.Response) (InstanceFailoverGroupsClientGetResponse, error) {
	result := InstanceFailoverGroupsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InstanceFailoverGroup); err != nil {
		return InstanceFailoverGroupsClientGetResponse{}, err
	}
	return result, nil
}

// ListByLocation - Lists the failover groups in a location.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// locationName - The name of the region where the resource is located.
// options - InstanceFailoverGroupsClientListByLocationOptions contains the optional parameters for the InstanceFailoverGroupsClient.ListByLocation
// method.
func (client *InstanceFailoverGroupsClient) ListByLocation(resourceGroupName string, locationName string, options *InstanceFailoverGroupsClientListByLocationOptions) *InstanceFailoverGroupsClientListByLocationPager {
	return &InstanceFailoverGroupsClientListByLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLocationCreateRequest(ctx, resourceGroupName, locationName, options)
		},
		advancer: func(ctx context.Context, resp InstanceFailoverGroupsClientListByLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.InstanceFailoverGroupListResult.NextLink)
		},
	}
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *InstanceFailoverGroupsClient) listByLocationCreateRequest(ctx context.Context, resourceGroupName string, locationName string, options *InstanceFailoverGroupsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *InstanceFailoverGroupsClient) listByLocationHandleResponse(resp *http.Response) (InstanceFailoverGroupsClientListByLocationResponse, error) {
	result := InstanceFailoverGroupsClientListByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InstanceFailoverGroupListResult); err != nil {
		return InstanceFailoverGroupsClientListByLocationResponse{}, err
	}
	return result, nil
}
