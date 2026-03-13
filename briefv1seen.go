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

// BriefV1SeenService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBriefV1SeenService] method instead.
type BriefV1SeenService struct {
	options []option.RequestOption
}

// NewBriefV1SeenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBriefV1SeenService(opts ...option.RequestOption) (r BriefV1SeenService) {
	r = BriefV1SeenService{}
	r.options = opts
	return
}

// Returns the list of paper group IDs the current user has already viewed as
// briefs
//
// Source file: `api-server/src/controllers/briefs/v1/get-seen.controller.ts`
func (r *BriefV1SeenService) GetSeen(ctx context.Context, opts ...option.RequestOption) (res *BriefV1SeenGetSeenResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "briefs/v1/seen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Records that the current user has viewed a brief for the given paper group
//
// Source file: `api-server/src/controllers/briefs/v1/mark-seen.controller.ts`
func (r *BriefV1SeenService) MarkSeen(ctx context.Context, body BriefV1SeenMarkSeenParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "briefs/v1/seen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type BriefV1SeenGetSeenResponse struct {
	SeenPaperGroupIDs []string `json:"seenPaperGroupIds" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeenPaperGroupIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefV1SeenGetSeenResponse) RawJSON() string { return r.JSON.raw }
func (r *BriefV1SeenGetSeenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefV1SeenMarkSeenParams struct {
	PaperGroupID string `json:"paperGroupId" api:"required" format:"uuid"`
	paramObj
}

func (r BriefV1SeenMarkSeenParams) MarshalJSON() (data []byte, err error) {
	type shadow BriefV1SeenMarkSeenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BriefV1SeenMarkSeenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
