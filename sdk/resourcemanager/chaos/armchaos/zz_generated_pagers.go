//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// CapabilitiesClientListPager provides operations for iterating over paged responses.
type CapabilitiesClientListPager struct {
	client    *CapabilitiesClient
	current   CapabilitiesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CapabilitiesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CapabilitiesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CapabilitiesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CapabilityListResult.NextLink == nil || len(*p.current.CapabilityListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current CapabilitiesClientListResponse page.
func (p *CapabilitiesClientListPager) PageResponse() CapabilitiesClientListResponse {
	return p.current
}

// CapabilityTypesClientListPager provides operations for iterating over paged responses.
type CapabilityTypesClientListPager struct {
	client    *CapabilityTypesClient
	current   CapabilityTypesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CapabilityTypesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CapabilityTypesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CapabilityTypesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CapabilityTypeListResult.NextLink == nil || len(*p.current.CapabilityTypeListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current CapabilityTypesClientListResponse page.
func (p *CapabilityTypesClientListPager) PageResponse() CapabilityTypesClientListResponse {
	return p.current
}

// ExperimentsClientListAllPager provides operations for iterating over paged responses.
type ExperimentsClientListAllPager struct {
	client    *ExperimentsClient
	current   ExperimentsClientListAllResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ExperimentsClientListAllResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ExperimentsClientListAllPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ExperimentsClientListAllPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExperimentListResult.NextLink == nil || len(*p.current.ExperimentListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAllHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ExperimentsClientListAllResponse page.
func (p *ExperimentsClientListAllPager) PageResponse() ExperimentsClientListAllResponse {
	return p.current
}

// ExperimentsClientListAllStatusesPager provides operations for iterating over paged responses.
type ExperimentsClientListAllStatusesPager struct {
	client    *ExperimentsClient
	current   ExperimentsClientListAllStatusesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ExperimentsClientListAllStatusesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ExperimentsClientListAllStatusesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ExperimentsClientListAllStatusesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExperimentStatusListResult.NextLink == nil || len(*p.current.ExperimentStatusListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAllStatusesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ExperimentsClientListAllStatusesResponse page.
func (p *ExperimentsClientListAllStatusesPager) PageResponse() ExperimentsClientListAllStatusesResponse {
	return p.current
}

// ExperimentsClientListExecutionDetailsPager provides operations for iterating over paged responses.
type ExperimentsClientListExecutionDetailsPager struct {
	client    *ExperimentsClient
	current   ExperimentsClientListExecutionDetailsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ExperimentsClientListExecutionDetailsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ExperimentsClientListExecutionDetailsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ExperimentsClientListExecutionDetailsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExperimentExecutionDetailsListResult.NextLink == nil || len(*p.current.ExperimentExecutionDetailsListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listExecutionDetailsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ExperimentsClientListExecutionDetailsResponse page.
func (p *ExperimentsClientListExecutionDetailsPager) PageResponse() ExperimentsClientListExecutionDetailsResponse {
	return p.current
}

// ExperimentsClientListPager provides operations for iterating over paged responses.
type ExperimentsClientListPager struct {
	client    *ExperimentsClient
	current   ExperimentsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ExperimentsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ExperimentsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ExperimentsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExperimentListResult.NextLink == nil || len(*p.current.ExperimentListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ExperimentsClientListResponse page.
func (p *ExperimentsClientListPager) PageResponse() ExperimentsClientListResponse {
	return p.current
}

// OperationsClientListAllPager provides operations for iterating over paged responses.
type OperationsClientListAllPager struct {
	client    *OperationsClient
	current   OperationsClientListAllResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListAllResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListAllPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListAllPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAllHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListAllResponse page.
func (p *OperationsClientListAllPager) PageResponse() OperationsClientListAllResponse {
	return p.current
}

// TargetTypesClientListPager provides operations for iterating over paged responses.
type TargetTypesClientListPager struct {
	client    *TargetTypesClient
	current   TargetTypesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TargetTypesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TargetTypesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TargetTypesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TargetTypeListResult.NextLink == nil || len(*p.current.TargetTypeListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TargetTypesClientListResponse page.
func (p *TargetTypesClientListPager) PageResponse() TargetTypesClientListResponse {
	return p.current
}

// TargetsClientListPager provides operations for iterating over paged responses.
type TargetsClientListPager struct {
	client    *TargetsClient
	current   TargetsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TargetsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TargetsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TargetsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TargetListResult.NextLink == nil || len(*p.current.TargetListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TargetsClientListResponse page.
func (p *TargetsClientListPager) PageResponse() TargetsClientListResponse {
	return p.current
}
