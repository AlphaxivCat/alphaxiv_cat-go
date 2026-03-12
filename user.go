// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/param"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options     []option.RequestOption
	V3          UserV3Service
	Preferences UserPreferenceService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	r.V3 = NewUserV3Service(opts...)
	r.Preferences = NewUserPreferenceService(opts...)
	return
}

// Get private notes for a user with pagination
//
// Source file:
// `api-server/src/controllers/v1/users/get-private-notes.controller.ts`
func (r *UserService) GetPrivateNotes(ctx context.Context, uid string, query UserGetPrivateNotesParams, opts ...option.RequestOption) (res *UserGetPrivateNotesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if uid == "" {
		err = errors.New("missing required uid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/notes", url.PathEscape(uid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Calculates and updates weekly reputation weights for users
//
// Source file:
// `api-server/src/controllers/v1/users/weigh-weekly-reputation.controller.ts`
func (r *UserService) WeighWeeklyReputation(ctx context.Context, opts ...option.RequestOption) (res *UserWeighWeeklyReputationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/users/weigh-weekly-reputation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type UserGetPrivateNotesResponse struct {
	Data UserGetPrivateNotesResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetPrivateNotesResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetPrivateNotesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetPrivateNotesResponseData struct {
	Comments []any   `json:"comments" api:"required"`
	Total    float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comments    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetPrivateNotesResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserGetPrivateNotesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserWeighWeeklyReputationResponse struct {
	Data UserWeighWeeklyReputationResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserWeighWeeklyReputationResponse) RawJSON() string { return r.JSON.raw }
func (r *UserWeighWeeklyReputationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserWeighWeeklyReputationResponseData struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserWeighWeeklyReputationResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserWeighWeeklyReputationResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetPrivateNotesParams struct {
	// Items per page (default: 10)
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Page number (default: 0)
	Page param.Opt[string] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserGetPrivateNotesParams]'s query parameters as
// `url.Values`.
func (r UserGetPrivateNotesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
