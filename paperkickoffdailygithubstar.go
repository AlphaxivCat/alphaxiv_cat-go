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
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// PaperKickoffDailyGitHubStarService contains methods and other services that help
// with interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperKickoffDailyGitHubStarService] method instead.
type PaperKickoffDailyGitHubStarService struct {
	options []option.RequestOption
}

// NewPaperKickoffDailyGitHubStarService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPaperKickoffDailyGitHubStarService(opts ...option.RequestOption) (r PaperKickoffDailyGitHubStarService) {
	r = PaperKickoffDailyGitHubStarService{}
	r.options = opts
	return
}

// Kickoff background job to update daily GitHub stars with max limit
//
// Source file:
// `api-server/src/controllers/v2/papers/kickoff-daily-github-stars-max.controller.ts`
func (r *PaperKickoffDailyGitHubStarService) Update(ctx context.Context, max string, opts ...option.RequestOption) (res *PaperKickoffDailyGitHubStarUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if max == "" {
		err = errors.New("missing required max parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/kickoff-daily-github-stars/%s", url.PathEscape(max))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Kickoff background job to update daily GitHub stars
//
// Source file:
// `api-server/src/controllers/v2/papers/kickoff-daily-github-stars.controller.ts`
func (r *PaperKickoffDailyGitHubStarService) KickoffDailyGitHubStars(ctx context.Context, opts ...option.RequestOption) (res *PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/papers/kickoff-daily-github-stars"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type PaperKickoffDailyGitHubStarUpdateResponse struct {
	Data PaperKickoffDailyGitHubStarUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffDailyGitHubStarUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffDailyGitHubStarUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffDailyGitHubStarUpdateResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffDailyGitHubStarUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperKickoffDailyGitHubStarUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponse struct {
	Data PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperKickoffDailyGitHubStarKickoffDailyGitHubStarsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
