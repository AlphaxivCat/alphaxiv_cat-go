// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserPreferenceService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPreferenceService] method instead.
type UserPreferenceService struct {
	options []option.RequestOption
	Email   UserPreferenceEmailService
}

// NewUserPreferenceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserPreferenceService(opts ...option.RequestOption) (r UserPreferenceService) {
	r = UserPreferenceService{}
	r.options = opts
	r.Email = NewUserPreferenceEmailService(opts...)
	return
}

// Get folder preferences for the authenticated user
//
// Source file:
// `api-server/src/controllers/v2/users/get-folders-preferences.controller.ts`
func (r *UserPreferenceService) GetFoldersPreferences(ctx context.Context, opts ...option.RequestOption) (res *UserPreferenceGetFoldersPreferencesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/users/preferences/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserPreferenceGetFoldersPreferencesResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceGetFoldersPreferencesResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceGetFoldersPreferencesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
