// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// APIKeyV1Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyV1Service] method instead.
type APIKeyV1Service struct {
	options []option.RequestOption
}

// NewAPIKeyV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIKeyV1Service(opts ...option.RequestOption) (r APIKeyV1Service) {
	r = APIKeyV1Service{}
	r.options = opts
	return
}

// Create a new API key for the current user.
//
// Source file:
// `api-server/src/controllers/api-keys/v1/create-api-key.controller.ts`
func (r *APIKeyV1Service) New(ctx context.Context, body APIKeyV1NewParams, opts ...option.RequestOption) (res *APIKeyV1NewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api-keys/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API keys for the current user.
//
// Source file: `api-server/src/controllers/api-keys/v1/get-api-keys.controller.ts`
func (r *APIKeyV1Service) List(ctx context.Context, opts ...option.RequestOption) (res *[]APIKeyV1ListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api-keys/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new API key for the current user.
//
// Source file:
// `api-server/src/controllers/api-keys/v1/create-impersonation-api-key.controller.ts`
func (r *APIKeyV1Service) NewImpersonation(ctx context.Context, body APIKeyV1NewImpersonationParams, opts ...option.RequestOption) (res *APIKeyV1NewImpersonationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api-keys/v1/impersonate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Revoke an API key for the authenticated user. No-op if already revoked.
//
// Source file:
// `api-server/src/controllers/api-keys/v1/revoke-api-key.controller.ts`
func (r *APIKeyV1Service) Revoke(ctx context.Context, apiKeyID APIKeyV1RevokeParamsAPIKeyID, opts ...option.RequestOption) (res *APIKeyV1RevokeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("api-keys/v1/%v/revoke", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type APIKeyV1NewResponse struct {
	APIKey APIKeyV1NewResponseAPIKey `json:"apiKey" api:"required"`
	Secret string                    `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyV1NewResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyV1NewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1NewResponseAPIKey struct {
	ID         string `json:"id" api:"required" format:"uuid"`
	Label      string `json:"label" api:"required"`
	LastUsedAt string `json:"lastUsedAt" api:"required"`
	Prefix     string `json:"prefix" api:"required"`
	RevokedAt  string `json:"revokedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		LastUsedAt  respjson.Field
		Prefix      respjson.Field
		RevokedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyV1NewResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKeyV1NewResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1ListResponse struct {
	ID         string `json:"id" api:"required" format:"uuid"`
	Label      string `json:"label" api:"required"`
	LastUsedAt string `json:"lastUsedAt" api:"required"`
	Prefix     string `json:"prefix" api:"required"`
	RevokedAt  string `json:"revokedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		LastUsedAt  respjson.Field
		Prefix      respjson.Field
		RevokedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyV1ListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyV1ListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1NewImpersonationResponse struct {
	APIKey APIKeyV1NewImpersonationResponseAPIKey `json:"apiKey" api:"required"`
	Secret string                                 `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyV1NewImpersonationResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyV1NewImpersonationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1NewImpersonationResponseAPIKey struct {
	ID         string `json:"id" api:"required" format:"uuid"`
	Label      string `json:"label" api:"required"`
	LastUsedAt string `json:"lastUsedAt" api:"required"`
	Prefix     string `json:"prefix" api:"required"`
	RevokedAt  string `json:"revokedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		LastUsedAt  respjson.Field
		Prefix      respjson.Field
		RevokedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyV1NewImpersonationResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKeyV1NewImpersonationResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1RevokeResponse struct {
	ID         string `json:"id" api:"required" format:"uuid"`
	Label      string `json:"label" api:"required"`
	LastUsedAt string `json:"lastUsedAt" api:"required"`
	Prefix     string `json:"prefix" api:"required"`
	RevokedAt  string `json:"revokedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		LastUsedAt  respjson.Field
		Prefix      respjson.Field
		RevokedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyV1RevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyV1RevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1NewParams struct {
	Label string `json:"label" api:"required"`
	paramObj
}

func (r APIKeyV1NewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyV1NewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyV1NewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1NewImpersonationParams struct {
	User string `json:"user" api:"required" format:"uuid"`
	paramObj
}

func (r APIKeyV1NewImpersonationParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyV1NewImpersonationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyV1NewImpersonationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyV1RevokeParamsAPIKeyID string

const (
	APIKeyV1RevokeParamsAPIKeyIDOwn APIKeyV1RevokeParamsAPIKeyID = "own"
)
