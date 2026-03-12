// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserPreferenceEmailService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPreferenceEmailService] method instead.
type UserPreferenceEmailService struct {
	options []option.RequestOption
}

// NewUserPreferenceEmailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserPreferenceEmailService(opts ...option.RequestOption) (r UserPreferenceEmailService) {
	r = UserPreferenceEmailService{}
	r.options = opts
	return
}

// Update email preferences for the authenticated user
//
// Source file:
// `api-server/src/controllers/v2/users/update-email-preferences.controller.ts`
func (r *UserPreferenceEmailService) Update(ctx context.Context, body UserPreferenceEmailUpdateParams, opts ...option.RequestOption) (res *UserPreferenceEmailUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/users/preferences/email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Get email preferences for the authenticated user
//
// Source file:
// `api-server/src/controllers/v2/users/get-email-preferences.controller.ts`
func (r *UserPreferenceEmailService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserPreferenceEmailGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/users/preferences/email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserPreferenceEmailUpdateResponse struct {
	Data UserPreferenceEmailUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceEmailUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceEmailUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceEmailUpdateResponseData struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceEmailUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceEmailUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceEmailGetResponse struct {
	Data UserPreferenceEmailGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceEmailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceEmailGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceEmailGetResponseData struct {
	Bounced             bool `json:"bounced" api:"required"`
	DirectNotifications bool `json:"directNotifications" api:"required"`
	RelevantActivity    bool `json:"relevantActivity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced             respjson.Field
		DirectNotifications respjson.Field
		RelevantActivity    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceEmailGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceEmailGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceEmailUpdateParams struct {
	DirectNotifications param.Opt[bool] `json:"directNotifications,omitzero"`
	RelevantActivity    param.Opt[bool] `json:"relevantActivity,omitzero"`
	paramObj
}

func (r UserPreferenceEmailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserPreferenceEmailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPreferenceEmailUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
