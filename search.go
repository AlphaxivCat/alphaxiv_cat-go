// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apijson"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// SearchService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	options []option.RequestOption
	V2      SearchV2Service
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.options = opts
	r.V2 = NewSearchV2Service(opts...)
	return
}

// Find the closest matching topics/categories for a given input using AI
//
// Source file: `api-server/src/controllers/v1/search/closest-topic.controller.ts`
func (r *SearchService) ClosestTopic(ctx context.Context, query SearchClosestTopicParams, opts ...option.RequestOption) (res *SearchClosestTopicResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/search/closest-topic"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search for papers using Google and enrich results
//
// Source file: `api-server/src/controllers/v1/search/search-google.controller.ts`
func (r *SearchService) GoogleSearch(ctx context.Context, query SearchGoogleSearchParams, opts ...option.RequestOption) (res *SearchGoogleSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/search/paper"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SearchClosestTopicResponse struct {
	Data []string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchClosestTopicResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchClosestTopicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchGoogleSearchResponse = any

type SearchClosestTopicParams struct {
	// User input to match against categories
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchClosestTopicParams]'s query parameters as
// `url.Values`.
func (r SearchClosestTopicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchGoogleSearchParams struct {
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchGoogleSearchParams]'s query parameters as
// `url.Values`.
func (r SearchGoogleSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
