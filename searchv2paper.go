// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apijson"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// SearchV2PaperService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchV2PaperService] method instead.
type SearchV2PaperService struct {
	options []option.RequestOption
}

// NewSearchV2PaperService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSearchV2PaperService(opts ...option.RequestOption) (r SearchV2PaperService) {
	r = SearchV2PaperService{}
	r.options = opts
	return
}

// Search for public and optionally private papers
//
// Source file:
// `api-server/src/controllers/search/v2/search-google-fast.controller.ts`
func (r *SearchV2PaperService) FastSearch(ctx context.Context, query SearchV2PaperFastSearchParams, opts ...option.RequestOption) (res *[]SearchV2PaperFastSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "search/v2/paper/fast"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SearchV2PaperFastSearchResponse struct {
	Link string `json:"link" api:"required"`
	// An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)
	PaperID   string `json:"paperId" api:"required"`
	Snippet   string `json:"snippet" api:"required"`
	Title     string `json:"title" api:"required"`
	IsPrivate bool   `json:"isPrivate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Link        respjson.Field
		PaperID     respjson.Field
		Snippet     respjson.Field
		Title       respjson.Field
		IsPrivate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchV2PaperFastSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchV2PaperFastSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchV2PaperFastSearchParams struct {
	// Any of "true", "false".
	IncludePrivate SearchV2PaperFastSearchParamsIncludePrivate `query:"includePrivate,omitzero" api:"required" json:"-"`
	Q              string                                      `query:"q" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchV2PaperFastSearchParams]'s query parameters as
// `url.Values`.
func (r SearchV2PaperFastSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchV2PaperFastSearchParamsIncludePrivate string

const (
	SearchV2PaperFastSearchParamsIncludePrivateTrue  SearchV2PaperFastSearchParamsIncludePrivate = "true"
	SearchV2PaperFastSearchParamsIncludePrivateFalse SearchV2PaperFastSearchParamsIncludePrivate = "false"
)
