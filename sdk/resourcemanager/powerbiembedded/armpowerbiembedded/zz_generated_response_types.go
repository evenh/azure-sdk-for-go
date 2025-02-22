//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiembedded

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ManagementClientGetAvailableOperationsResponse contains the response from method ManagementClient.GetAvailableOperations.
type ManagementClientGetAvailableOperationsResponse struct {
	ManagementClientGetAvailableOperationsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementClientGetAvailableOperationsResult contains the result from method ManagementClient.GetAvailableOperations.
type ManagementClientGetAvailableOperationsResult struct {
	OperationList
}

// WorkspaceCollectionsClientCheckNameAvailabilityResponse contains the response from method WorkspaceCollectionsClient.CheckNameAvailability.
type WorkspaceCollectionsClientCheckNameAvailabilityResponse struct {
	WorkspaceCollectionsClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientCheckNameAvailabilityResult contains the result from method WorkspaceCollectionsClient.CheckNameAvailability.
type WorkspaceCollectionsClientCheckNameAvailabilityResult struct {
	CheckNameResponse
}

// WorkspaceCollectionsClientCreateResponse contains the response from method WorkspaceCollectionsClient.Create.
type WorkspaceCollectionsClientCreateResponse struct {
	WorkspaceCollectionsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientCreateResult contains the result from method WorkspaceCollectionsClient.Create.
type WorkspaceCollectionsClientCreateResult struct {
	WorkspaceCollection
}

// WorkspaceCollectionsClientDeletePollerResponse contains the response from method WorkspaceCollectionsClient.Delete.
type WorkspaceCollectionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspaceCollectionsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l WorkspaceCollectionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspaceCollectionsClientDeleteResponse, error) {
	respType := WorkspaceCollectionsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspaceCollectionsClientDeletePollerResponse from the provided client and resume token.
func (l *WorkspaceCollectionsClientDeletePollerResponse) Resume(ctx context.Context, client *WorkspaceCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspaceCollectionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &WorkspaceCollectionsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspaceCollectionsClientDeleteResponse contains the response from method WorkspaceCollectionsClient.Delete.
type WorkspaceCollectionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientGetAccessKeysResponse contains the response from method WorkspaceCollectionsClient.GetAccessKeys.
type WorkspaceCollectionsClientGetAccessKeysResponse struct {
	WorkspaceCollectionsClientGetAccessKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientGetAccessKeysResult contains the result from method WorkspaceCollectionsClient.GetAccessKeys.
type WorkspaceCollectionsClientGetAccessKeysResult struct {
	WorkspaceCollectionAccessKeys
}

// WorkspaceCollectionsClientGetByNameResponse contains the response from method WorkspaceCollectionsClient.GetByName.
type WorkspaceCollectionsClientGetByNameResponse struct {
	WorkspaceCollectionsClientGetByNameResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientGetByNameResult contains the result from method WorkspaceCollectionsClient.GetByName.
type WorkspaceCollectionsClientGetByNameResult struct {
	WorkspaceCollection
}

// WorkspaceCollectionsClientListByResourceGroupResponse contains the response from method WorkspaceCollectionsClient.ListByResourceGroup.
type WorkspaceCollectionsClientListByResourceGroupResponse struct {
	WorkspaceCollectionsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientListByResourceGroupResult contains the result from method WorkspaceCollectionsClient.ListByResourceGroup.
type WorkspaceCollectionsClientListByResourceGroupResult struct {
	WorkspaceCollectionList
}

// WorkspaceCollectionsClientListBySubscriptionResponse contains the response from method WorkspaceCollectionsClient.ListBySubscription.
type WorkspaceCollectionsClientListBySubscriptionResponse struct {
	WorkspaceCollectionsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientListBySubscriptionResult contains the result from method WorkspaceCollectionsClient.ListBySubscription.
type WorkspaceCollectionsClientListBySubscriptionResult struct {
	WorkspaceCollectionList
}

// WorkspaceCollectionsClientMigrateResponse contains the response from method WorkspaceCollectionsClient.Migrate.
type WorkspaceCollectionsClientMigrateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientRegenerateKeyResponse contains the response from method WorkspaceCollectionsClient.RegenerateKey.
type WorkspaceCollectionsClientRegenerateKeyResponse struct {
	WorkspaceCollectionsClientRegenerateKeyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientRegenerateKeyResult contains the result from method WorkspaceCollectionsClient.RegenerateKey.
type WorkspaceCollectionsClientRegenerateKeyResult struct {
	WorkspaceCollectionAccessKeys
}

// WorkspaceCollectionsClientUpdateResponse contains the response from method WorkspaceCollectionsClient.Update.
type WorkspaceCollectionsClientUpdateResponse struct {
	WorkspaceCollectionsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceCollectionsClientUpdateResult contains the result from method WorkspaceCollectionsClient.Update.
type WorkspaceCollectionsClientUpdateResult struct {
	WorkspaceCollection
}

// WorkspacesClientListResponse contains the response from method WorkspacesClient.List.
type WorkspacesClientListResponse struct {
	WorkspacesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesClientListResult contains the result from method WorkspacesClient.List.
type WorkspacesClientListResult struct {
	WorkspaceList
}
