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

// RegulatoryComplianceAssessmentsClient contains the methods for the RegulatoryComplianceAssessments group.
// Don't use this type directly, use NewRegulatoryComplianceAssessmentsClient() instead.
type RegulatoryComplianceAssessmentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRegulatoryComplianceAssessmentsClient creates a new instance of RegulatoryComplianceAssessmentsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegulatoryComplianceAssessmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RegulatoryComplianceAssessmentsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RegulatoryComplianceAssessmentsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Supported regulatory compliance details and state for selected assessment
// If the operation fails it returns an *azcore.ResponseError type.
// regulatoryComplianceStandardName - Name of the regulatory compliance standard object
// regulatoryComplianceControlName - Name of the regulatory compliance control object
// regulatoryComplianceAssessmentName - Name of the regulatory compliance assessment object
// options - RegulatoryComplianceAssessmentsClientGetOptions contains the optional parameters for the RegulatoryComplianceAssessmentsClient.Get
// method.
func (client *RegulatoryComplianceAssessmentsClient) Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string, options *RegulatoryComplianceAssessmentsClientGetOptions) (RegulatoryComplianceAssessmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, regulatoryComplianceAssessmentName, options)
	if err != nil {
		return RegulatoryComplianceAssessmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegulatoryComplianceAssessmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegulatoryComplianceAssessmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegulatoryComplianceAssessmentsClient) getCreateRequest(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string, options *RegulatoryComplianceAssessmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments/{regulatoryComplianceAssessmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regulatoryComplianceStandardName == "" {
		return nil, errors.New("parameter regulatoryComplianceStandardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceStandardName}", url.PathEscape(regulatoryComplianceStandardName))
	if regulatoryComplianceControlName == "" {
		return nil, errors.New("parameter regulatoryComplianceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceControlName}", url.PathEscape(regulatoryComplianceControlName))
	if regulatoryComplianceAssessmentName == "" {
		return nil, errors.New("parameter regulatoryComplianceAssessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceAssessmentName}", url.PathEscape(regulatoryComplianceAssessmentName))
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
func (client *RegulatoryComplianceAssessmentsClient) getHandleResponse(resp *http.Response) (RegulatoryComplianceAssessmentsClientGetResponse, error) {
	result := RegulatoryComplianceAssessmentsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceAssessment); err != nil {
		return RegulatoryComplianceAssessmentsClientGetResponse{}, err
	}
	return result, nil
}

// List - Details and state of assessments mapped to selected regulatory compliance control
// If the operation fails it returns an *azcore.ResponseError type.
// regulatoryComplianceStandardName - Name of the regulatory compliance standard object
// regulatoryComplianceControlName - Name of the regulatory compliance control object
// options - RegulatoryComplianceAssessmentsClientListOptions contains the optional parameters for the RegulatoryComplianceAssessmentsClient.List
// method.
func (client *RegulatoryComplianceAssessmentsClient) List(regulatoryComplianceStandardName string, regulatoryComplianceControlName string, options *RegulatoryComplianceAssessmentsClientListOptions) *RegulatoryComplianceAssessmentsClientListPager {
	return &RegulatoryComplianceAssessmentsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, options)
		},
		advancer: func(ctx context.Context, resp RegulatoryComplianceAssessmentsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RegulatoryComplianceAssessmentList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RegulatoryComplianceAssessmentsClient) listCreateRequest(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, options *RegulatoryComplianceAssessmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regulatoryComplianceStandardName == "" {
		return nil, errors.New("parameter regulatoryComplianceStandardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceStandardName}", url.PathEscape(regulatoryComplianceStandardName))
	if regulatoryComplianceControlName == "" {
		return nil, errors.New("parameter regulatoryComplianceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceControlName}", url.PathEscape(regulatoryComplianceControlName))
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
func (client *RegulatoryComplianceAssessmentsClient) listHandleResponse(resp *http.Response) (RegulatoryComplianceAssessmentsClientListResponse, error) {
	result := RegulatoryComplianceAssessmentsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceAssessmentList); err != nil {
		return RegulatoryComplianceAssessmentsClientListResponse{}, err
	}
	return result, nil
}
