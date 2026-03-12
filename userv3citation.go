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
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// UserV3CitationService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV3CitationService] method instead.
type UserV3CitationService struct {
	options []option.RequestOption
}

// NewUserV3CitationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserV3CitationService(opts ...option.RequestOption) (r UserV3CitationService) {
	r = UserV3CitationService{}
	r.options = opts
	return
}

// Retrieve citation counts by year for a user
//
// Source file:
// `api-server/src/controllers/users/v3/get-citation-graph.controller.ts`
func (r *UserV3CitationService) GetGraph(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3CitationGetGraphResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/citations/graph", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve aggregated citation metrics for a user
//
// Source file:
// `api-server/src/controllers/users/v3/get-citation-summary.controller.ts`
func (r *UserV3CitationService) GetSummary(ctx context.Context, id string, opts ...option.RequestOption) (res *UserV3CitationGetSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/v3/%s/citations/summary", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserV3CitationGetGraphResponse struct {
	Graph []UserV3CitationGetGraphResponseGraph `json:"graph" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Graph       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3CitationGetGraphResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3CitationGetGraphResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3CitationGetGraphResponseGraph struct {
	Citations float64 `json:"citations" api:"required"`
	Year      float64 `json:"year" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Citations   respjson.Field
		Year        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3CitationGetGraphResponseGraph) RawJSON() string { return r.JSON.raw }
func (r *UserV3CitationGetGraphResponseGraph) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3CitationGetSummaryResponse struct {
	Citations UserV3CitationGetSummaryResponseCitations `json:"citations" api:"required"`
	HIndex    UserV3CitationGetSummaryResponseHIndex    `json:"h_index" api:"required"`
	I10Index  UserV3CitationGetSummaryResponseI10Index  `json:"i10_index" api:"required"`
	UpdatedAt string                                    `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Citations   respjson.Field
		HIndex      respjson.Field
		I10Index    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3CitationGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *UserV3CitationGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3CitationGetSummaryResponseCitations struct {
	All       float64 `json:"all" api:"required"`
	Since2020 float64 `json:"since_2020" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Since2020   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3CitationGetSummaryResponseCitations) RawJSON() string { return r.JSON.raw }
func (r *UserV3CitationGetSummaryResponseCitations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3CitationGetSummaryResponseHIndex struct {
	All       float64 `json:"all" api:"required"`
	Since2020 float64 `json:"since_2020" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Since2020   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3CitationGetSummaryResponseHIndex) RawJSON() string { return r.JSON.raw }
func (r *UserV3CitationGetSummaryResponseHIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserV3CitationGetSummaryResponseI10Index struct {
	All       float64 `json:"all" api:"required"`
	Since2020 float64 `json:"since_2020" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Since2020   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserV3CitationGetSummaryResponseI10Index) RawJSON() string { return r.JSON.raw }
func (r *UserV3CitationGetSummaryResponseI10Index) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
