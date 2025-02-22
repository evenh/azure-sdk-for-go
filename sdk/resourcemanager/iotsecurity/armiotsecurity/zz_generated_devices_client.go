//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

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

// DevicesClient contains the methods for the Devices group.
// Don't use this type directly, use NewDevicesClient() instead.
type DevicesClient struct {
	host                string
	subscriptionID      string
	iotDefenderLocation string
	pl                  runtime.Pipeline
}

// NewDevicesClient creates a new instance of DevicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// iotDefenderLocation - Defender for IoT location
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDevicesClient(subscriptionID string, iotDefenderLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *DevicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DevicesClient{
		subscriptionID:      subscriptionID,
		iotDefenderLocation: iotDefenderLocation,
		host:                string(cp.Endpoint),
		pl:                  armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get device
// If the operation fails it returns an *azcore.ResponseError type.
// deviceGroupName - Device group name
// deviceID - Device Id
// options - DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
func (client *DevicesClient) Get(ctx context.Context, deviceGroupName string, deviceID string, options *DevicesClientGetOptions) (DevicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceGroupName, deviceID, options)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DevicesClient) getCreateRequest(ctx context.Context, deviceGroupName string, deviceID string, options *DevicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/devices/{deviceId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceID == "" {
		return nil, errors.New("parameter deviceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceId}", url.PathEscape(deviceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevicesClient) getHandleResponse(resp *http.Response) (DevicesClientGetResponse, error) {
	result := DevicesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceModel); err != nil {
		return DevicesClientGetResponse{}, err
	}
	return result, nil
}

// List - List devices
// If the operation fails it returns an *azcore.ResponseError type.
// deviceGroupName - Device group name
// options - DevicesClientListOptions contains the optional parameters for the DevicesClient.List method.
func (client *DevicesClient) List(deviceGroupName string, options *DevicesClientListOptions) *DevicesClientListPager {
	return &DevicesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, deviceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DevicesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeviceList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DevicesClient) listCreateRequest(ctx context.Context, deviceGroupName string, options *DevicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/devices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DevicesClient) listHandleResponse(resp *http.Response) (DevicesClientListResponse, error) {
	result := DevicesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceList); err != nil {
		return DevicesClientListResponse{}, err
	}
	return result, nil
}
