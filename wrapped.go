// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// WrappedService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWrappedService] method instead.
type WrappedService struct {
	options []option.RequestOption
}

// NewWrappedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWrappedService(opts ...option.RequestOption) (r WrappedService) {
	r = WrappedService{}
	r.options = opts
	return
}

// Get all wrapped data for a specific user across all years
//
// Source file:
// `api-server/src/controllers/wrapped/v1/get-wrapped-by-user.controller.ts`
func (r *WrappedService) GetByUser(ctx context.Context, userID string, opts ...option.RequestOption) (res *[]WrappedGetByUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("wrapped/v1/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WrappedGetByUserResponse struct {
	ID    string  `json:"id" api:"required"`
	Order float64 `json:"order" api:"required"`
	Type  string  `json:"type" api:"required"`
	Year  float64 `json:"year" api:"required"`
	Data  any     `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Order       respjson.Field
		Type        respjson.Field
		Year        respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WrappedGetByUserResponse) RawJSON() string { return r.JSON.raw }
func (r *WrappedGetByUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
