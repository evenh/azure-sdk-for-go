//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvisualstudio

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

// ExtensionsClient contains the methods for the Extensions group.
// Don't use this type directly, use NewExtensionsClient() instead.
type ExtensionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExtensionsClient creates a new instance of ExtensionsClient with the specified values.
// subscriptionID - The Azure subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExtensionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ExtensionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Create - Registers the extension with a Visual Studio Team Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the Azure subscription.
// accountResourceName - The name of the Visual Studio Team Services account resource.
// extensionResourceName - The name of the extension.
// body - An object containing additional information related to the extension request.
// options - ExtensionsClientCreateOptions contains the optional parameters for the ExtensionsClient.Create method.
func (client *ExtensionsClient) Create(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, body ExtensionResourceRequest, options *ExtensionsClientCreateOptions) (ExtensionsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountResourceName, extensionResourceName, body, options)
	if err != nil {
		return ExtensionsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ExtensionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, body ExtensionResourceRequest, options *ExtensionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountResourceName == "" {
		return nil, errors.New("parameter accountResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountResourceName}", url.PathEscape(accountResourceName))
	if extensionResourceName == "" {
		return nil, errors.New("parameter extensionResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionResourceName}", url.PathEscape(extensionResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleResponse handles the Create response.
func (client *ExtensionsClient) createHandleResponse(resp *http.Response) (ExtensionsClientCreateResponse, error) {
	result := ExtensionsClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionResource); err != nil {
		return ExtensionsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Removes an extension resource registration for a Visual Studio Team Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the Azure subscription.
// accountResourceName - The name of the Visual Studio Team Services account resource.
// extensionResourceName - The name of the extension.
// options - ExtensionsClientDeleteOptions contains the optional parameters for the ExtensionsClient.Delete method.
func (client *ExtensionsClient) Delete(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, options *ExtensionsClientDeleteOptions) (ExtensionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountResourceName, extensionResourceName, options)
	if err != nil {
		return ExtensionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ExtensionsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, options *ExtensionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountResourceName == "" {
		return nil, errors.New("parameter accountResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountResourceName}", url.PathEscape(accountResourceName))
	if extensionResourceName == "" {
		return nil, errors.New("parameter extensionResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionResourceName}", url.PathEscape(extensionResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the details of an extension associated with a Visual Studio Team Services account resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the Azure subscription.
// accountResourceName - The name of the Visual Studio Team Services account resource.
// extensionResourceName - The name of the extension.
// options - ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
func (client *ExtensionsClient) Get(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, options *ExtensionsClientGetOptions) (ExtensionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountResourceName, extensionResourceName, options)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return ExtensionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, options *ExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountResourceName == "" {
		return nil, errors.New("parameter accountResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountResourceName}", url.PathEscape(accountResourceName))
	if extensionResourceName == "" {
		return nil, errors.New("parameter extensionResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionResourceName}", url.PathEscape(extensionResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionsClient) getHandleResponse(resp *http.Response) (ExtensionsClientGetResponse, error) {
	result := ExtensionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionResource); err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByAccount - Gets the details of the extension resources created within the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the Azure subscription.
// accountResourceName - The name of the Visual Studio Team Services account resource.
// options - ExtensionsClientListByAccountOptions contains the optional parameters for the ExtensionsClient.ListByAccount
// method.
func (client *ExtensionsClient) ListByAccount(ctx context.Context, resourceGroupName string, accountResourceName string, options *ExtensionsClientListByAccountOptions) (ExtensionsClientListByAccountResponse, error) {
	req, err := client.listByAccountCreateRequest(ctx, resourceGroupName, accountResourceName, options)
	if err != nil {
		return ExtensionsClientListByAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientListByAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientListByAccountResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByAccountHandleResponse(resp)
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *ExtensionsClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountResourceName string, options *ExtensionsClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountResourceName == "" {
		return nil, errors.New("parameter accountResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountResourceName}", url.PathEscape(accountResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *ExtensionsClient) listByAccountHandleResponse(resp *http.Response) (ExtensionsClientListByAccountResponse, error) {
	result := ExtensionsClientListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionResourceListResult); err != nil {
		return ExtensionsClientListByAccountResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing extension registration for the Visual Studio Team Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the Azure subscription.
// accountResourceName - The name of the Visual Studio Team Services account resource.
// extensionResourceName - The name of the extension.
// body - An object containing additional information related to the extension request.
// options - ExtensionsClientUpdateOptions contains the optional parameters for the ExtensionsClient.Update method.
func (client *ExtensionsClient) Update(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, body ExtensionResourceRequest, options *ExtensionsClientUpdateOptions) (ExtensionsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountResourceName, extensionResourceName, body, options)
	if err != nil {
		return ExtensionsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string, body ExtensionResourceRequest, options *ExtensionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountResourceName == "" {
		return nil, errors.New("parameter accountResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountResourceName}", url.PathEscape(accountResourceName))
	if extensionResourceName == "" {
		return nil, errors.New("parameter extensionResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionResourceName}", url.PathEscape(extensionResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *ExtensionsClient) updateHandleResponse(resp *http.Response) (ExtensionsClientUpdateResponse, error) {
	result := ExtensionsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionResource); err != nil {
		return ExtensionsClientUpdateResponse{}, err
	}
	return result, nil
}
