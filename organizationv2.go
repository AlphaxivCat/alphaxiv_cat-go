// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// OrganizationV2Service contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationV2Service] method instead.
type OrganizationV2Service struct {
	options []option.RequestOption
}

// NewOrganizationV2Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationV2Service(opts ...option.RequestOption) (r OrganizationV2Service) {
	r = OrganizationV2Service{}
	r.options = opts
	return
}

// Retrieve the top 20 organizations with images by paper count
//
// Source file:
// `api-server/src/controllers/organizations/v2/get-top-orgs.controller.ts`
func (r *OrganizationV2Service) ListTop(ctx context.Context, opts ...option.RequestOption) (res *[]OrganizationV2ListTopResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "organizations/v2/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve an organization's basic information given its ID.
//
// Source file:
// `api-server/src/controllers/organizations/v2/get-organization-by-id.controller.ts`
func (r *OrganizationV2Service) GetByID(ctx context.Context, id string, opts ...option.RequestOption) (res *OrganizationV2GetByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/v2/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve an organization's basic information given it's name – casing does not
// matter (i.e. google and GOOGLE return the same results).
//
// Source file:
// `api-server/src/controllers/organizations/v2/get-organization-by-name.controller.ts`
func (r *OrganizationV2Service) GetByName(ctx context.Context, name string, opts ...option.RequestOption) (res *OrganizationV2GetByNameResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/v2/by-name/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search organizations by name
//
// Source file: `api-server/src/controllers/organizations/v2/search.controller.ts`
func (r *OrganizationV2Service) Search(ctx context.Context, query OrganizationV2SearchParams, opts ...option.RequestOption) (res *[]OrganizationV2SearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "organizations/v2/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Toggle following an organization, affects the current user's profile
//
// Source file:
// `api-server/src/controllers/organizations/v2/toggle-follow-org.controller.ts`
func (r *OrganizationV2Service) ToggleFollow(ctx context.Context, id string, opts ...option.RequestOption) (res *OrganizationV2ToggleFollowResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/v2/%s/toggle", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type OrganizationV2ListTopResponse struct {
	ID    string `json:"id" api:"required" format:"uuid"`
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	Slug  string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Image       respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationV2ListTopResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationV2ListTopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationV2GetByIDResponse struct {
	ID    string `json:"id" api:"required" format:"uuid"`
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	Slug  string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Image       respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationV2GetByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationV2GetByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationV2GetByNameResponse struct {
	ID    string `json:"id" api:"required" format:"uuid"`
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	Slug  string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Image       respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationV2GetByNameResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationV2GetByNameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationV2SearchResponse struct {
	ID    string `json:"id" api:"required" format:"uuid"`
	Image string `json:"image" api:"required"`
	Name  string `json:"name" api:"required"`
	Slug  string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Image       respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationV2SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationV2SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationV2ToggleFollowResponse struct {
	Following bool `json:"following" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Following   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationV2ToggleFollowResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationV2ToggleFollowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationV2SearchParams struct {
	Q string `query:"q" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationV2SearchParams]'s query parameters as
// `url.Values`.
func (r OrganizationV2SearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
