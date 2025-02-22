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

// RecoverableDatabasesClient contains the methods for the RecoverableDatabases group.
// Don't use this type directly, use NewRecoverableDatabasesClient() instead.
type RecoverableDatabasesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRecoverableDatabasesClient creates a new instance of RecoverableDatabasesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRecoverableDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RecoverableDatabasesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &RecoverableDatabasesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets a recoverable database, which is a resource representing a database's geo backup
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database
// options - RecoverableDatabasesClientGetOptions contains the optional parameters for the RecoverableDatabasesClient.Get
// method.
func (client *RecoverableDatabasesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *RecoverableDatabasesClientGetOptions) (RecoverableDatabasesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
	if err != nil {
		return RecoverableDatabasesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecoverableDatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecoverableDatabasesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecoverableDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *RecoverableDatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recoverableDatabases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoverableDatabasesClient) getHandleResponse(resp *http.Response) (RecoverableDatabasesClientGetResponse, error) {
	result := RecoverableDatabasesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableDatabase); err != nil {
		return RecoverableDatabasesClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Gets a list of recoverable databases
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - RecoverableDatabasesClientListByServerOptions contains the optional parameters for the RecoverableDatabasesClient.ListByServer
// method.
func (client *RecoverableDatabasesClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, options *RecoverableDatabasesClientListByServerOptions) (RecoverableDatabasesClientListByServerResponse, error) {
	req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return RecoverableDatabasesClientListByServerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecoverableDatabasesClientListByServerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecoverableDatabasesClientListByServerResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByServerHandleResponse(resp)
}

// listByServerCreateRequest creates the ListByServer request.
func (client *RecoverableDatabasesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *RecoverableDatabasesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recoverableDatabases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *RecoverableDatabasesClient) listByServerHandleResponse(resp *http.Response) (RecoverableDatabasesClientListByServerResponse, error) {
	result := RecoverableDatabasesClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableDatabaseListResult); err != nil {
		return RecoverableDatabasesClientListByServerResponse{}, err
	}
	return result, nil
}
