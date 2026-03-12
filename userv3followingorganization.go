// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserV3FollowingOrganizationService contains methods and other services that help
// with interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3FollowingOrganizationService] method instead.
type UserV3FollowingOrganizationService struct {
	options []option.RequestOption
}

// NewUserV3FollowingOrganizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserV3FollowingOrganizationService(opts ...option.RequestOption) (r UserV3FollowingOrganizationService) {
	r = UserV3FollowingOrganizationService{}
	r.options = opts
	return
}

// List the organizations followed by a user
//
// Source file:
// `api-server/src/controllers/users/v3/get-following-organizations.controller.ts`
func (r *UserV3FollowingOrganizationService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3FollowingOrganizationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/following/organizations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Toggle following an organization for the current user
//
// Source file:
// `api-server/src/controllers/users/v3/toggle-following-organization.controller.ts`
func (r *UserV3FollowingOrganizationService) Toggle(ctx context.Context, id string, body UserV3FollowingOrganizationToggleParams, opts ...option.RequestOption) (res *UserV3FollowingOrganizationToggleResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/following/organizations/toggle", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UserV3FollowingOrganizationListResponse struct {
	Organizations []UserV3FollowingOrganizationListResponseOrganization `json:"organizations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organizations respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3FollowingOrganizationListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingOrganizationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingOrganizationListResponseOrganization struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3FollowingOrganizationListResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingOrganizationListResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingOrganizationToggleResponse struct {
	Following    bool   `json:"following" api:"required"`
	Organization string `json:"organization" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Following    respjson.Field
		Organization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3FollowingOrganizationToggleResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3FollowingOrganizationToggleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3FollowingOrganizationToggleParams struct {
	Organization string `json:"organization" api:"required"`
	paramObj
}

func (r UserV3FollowingOrganizationToggleParams) MarshalJSON() (data []byte, err error) {
	type shadow UserV3FollowingOrganizationToggleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserV3FollowingOrganizationToggleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
