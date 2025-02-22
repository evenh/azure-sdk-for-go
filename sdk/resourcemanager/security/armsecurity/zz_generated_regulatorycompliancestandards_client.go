//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// RegulatoryComplianceStandardsClient contains the methods for the RegulatoryComplianceStandards group.
// Don't use this type directly, use NewRegulatoryComplianceStandardsClient() instead.
type RegulatoryComplianceStandardsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRegulatoryComplianceStandardsClient creates a new instance of RegulatoryComplianceStandardsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegulatoryComplianceStandardsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RegulatoryComplianceStandardsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RegulatoryComplianceStandardsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Supported regulatory compliance details state for selected standard
// If the operation fails it returns an *azcore.ResponseError type.
// regulatoryComplianceStandardName - Name of the regulatory compliance standard object
// options - RegulatoryComplianceStandardsClientGetOptions contains the optional parameters for the RegulatoryComplianceStandardsClient.Get
// method.
func (client *RegulatoryComplianceStandardsClient) Get(ctx context.Context, regulatoryComplianceStandardName string, options *RegulatoryComplianceStandardsClientGetOptions) (RegulatoryComplianceStandardsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regulatoryComplianceStandardName, options)
	if err != nil {
		return RegulatoryComplianceStandardsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegulatoryComplianceStandardsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegulatoryComplianceStandardsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegulatoryComplianceStandardsClient) getCreateRequest(ctx context.Context, regulatoryComplianceStandardName string, options *RegulatoryComplianceStandardsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regulatoryComplianceStandardName == "" {
		return nil, errors.New("parameter regulatoryComplianceStandardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceStandardName}", url.PathEscape(regulatoryComplianceStandardName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegulatoryComplianceStandardsClient) getHandleResponse(resp *http.Response) (RegulatoryComplianceStandardsClientGetResponse, error) {
	result := RegulatoryComplianceStandardsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceStandard); err != nil {
		return RegulatoryComplianceStandardsClientGetResponse{}, err
	}
	return result, nil
}

// List - Supported regulatory compliance standards details and state
// If the operation fails it returns an *azcore.ResponseError type.
// options - RegulatoryComplianceStandardsClientListOptions contains the optional parameters for the RegulatoryComplianceStandardsClient.List
// method.
func (client *RegulatoryComplianceStandardsClient) List(options *RegulatoryComplianceStandardsClientListOptions) *RegulatoryComplianceStandardsClientListPager {
	return &RegulatoryComplianceStandardsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp RegulatoryComplianceStandardsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RegulatoryComplianceStandardList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RegulatoryComplianceStandardsClient) listCreateRequest(ctx context.Context, options *RegulatoryComplianceStandardsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegulatoryComplianceStandardsClient) listHandleResponse(resp *http.Response) (RegulatoryComplianceStandardsClientListResponse, error) {
	result := RegulatoryComplianceStandardsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceStandardList); err != nil {
		return RegulatoryComplianceStandardsClientListResponse{}, err
	}
	return result, nil
}
