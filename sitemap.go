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
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/respjson"
)

// SitemapService contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSitemapService] method instead.
type SitemapService struct {
	options []option.RequestOption
}

// NewSitemapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSitemapService(opts ...option.RequestOption) (r SitemapService) {
	r = SitemapService{}
	r.options = opts
	return
}

// Get paginated list of paper versions with AI overviews for sitemap generation
//
// Source file:
// `api-server/src/controllers/v1/sitemaps/get-overviews-for-sitemap.controller.ts`
func (r *SitemapService) ListOverviews(ctx context.Context, query SitemapListOverviewsParams, opts ...option.RequestOption) (res *SitemapListOverviewsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sitemaps/overviews"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get paginated list of public papers for sitemap generation. Uses cursor caching
// for efficient deep pagination.
//
// Source file:
// `api-server/src/controllers/v1/sitemaps/get-papers-for-sitemap.controller.ts`
func (r *SitemapService) ListPapers(ctx context.Context, query SitemapListPapersParams, opts ...option.RequestOption) (res *SitemapListPapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sitemaps/papers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get paginated list of users for sitemap generation
//
// Source file:
// `api-server/src/controllers/v1/sitemaps/get-users-for-sitemap.controller.ts`
func (r *SitemapService) ListUsers(ctx context.Context, query SitemapListUsersParams, opts ...option.RequestOption) (res *SitemapListUsersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sitemaps/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SitemapListOverviewsResponse struct {
	Data SitemapListOverviewsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListOverviewsResponse) RawJSON() string { return r.JSON.raw }
func (r *SitemapListOverviewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListOverviewsResponseData struct {
	Pagination SitemapListOverviewsResponseDataPagination `json:"pagination" api:"required"`
	Papers     []SitemapListOverviewsResponseDataPaper    `json:"papers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Papers      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListOverviewsResponseData) RawJSON() string { return r.JSON.raw }
func (r *SitemapListOverviewsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListOverviewsResponseDataPagination struct {
	Total float64 `json:"total" api:"required"`
	Limit float64 `json:"limit"`
	Page  float64 `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total       respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListOverviewsResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *SitemapListOverviewsResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListOverviewsResponseDataPaper struct {
	ID               string  `json:"id" api:"required" format:"uuid"`
	Language         string  `json:"language" api:"required"`
	UniversalPaperID string  `json:"universal_paper_id" api:"required"`
	UpdatedAt        string  `json:"updated_at" api:"required"`
	VersionOrder     float64 `json:"version_order" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Language         respjson.Field
		UniversalPaperID respjson.Field
		UpdatedAt        respjson.Field
		VersionOrder     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListOverviewsResponseDataPaper) RawJSON() string { return r.JSON.raw }
func (r *SitemapListOverviewsResponseDataPaper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListPapersResponse struct {
	Data SitemapListPapersResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListPapersResponse) RawJSON() string { return r.JSON.raw }
func (r *SitemapListPapersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListPapersResponseData struct {
	Pagination SitemapListPapersResponseDataPagination `json:"pagination" api:"required"`
	Papers     []SitemapListPapersResponseDataPaper    `json:"papers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Papers      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListPapersResponseData) RawJSON() string { return r.JSON.raw }
func (r *SitemapListPapersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListPapersResponseDataPagination struct {
	Limit float64 `json:"limit" api:"required"`
	Page  float64 `json:"page" api:"required"`
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListPapersResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *SitemapListPapersResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListPapersResponseDataPaper struct {
	ID          string `json:"id" api:"required" format:"uuid"`
	UniversalID string `json:"universalId" api:"required"`
	UpdatedAt   string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		UniversalID respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListPapersResponseDataPaper) RawJSON() string { return r.JSON.raw }
func (r *SitemapListPapersResponseDataPaper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListUsersResponse struct {
	Data SitemapListUsersResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListUsersResponse) RawJSON() string { return r.JSON.raw }
func (r *SitemapListUsersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListUsersResponseData struct {
	Pagination SitemapListUsersResponseDataPagination `json:"pagination" api:"required"`
	Users      []SitemapListUsersResponseDataUser     `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListUsersResponseData) RawJSON() string { return r.JSON.raw }
func (r *SitemapListUsersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListUsersResponseDataPagination struct {
	Total float64 `json:"total" api:"required"`
	Limit float64 `json:"limit"`
	Page  float64 `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total       respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListUsersResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *SitemapListUsersResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListUsersResponseDataUser struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SitemapListUsersResponseDataUser) RawJSON() string { return r.JSON.raw }
func (r *SitemapListUsersResponseDataUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SitemapListOverviewsParams struct {
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[string] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SitemapListOverviewsParams]'s query parameters as
// `url.Values`.
func (r SitemapListOverviewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SitemapListPapersParams struct {
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[string] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SitemapListPapersParams]'s query parameters as
// `url.Values`.
func (r SitemapListPapersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SitemapListUsersParams struct {
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[string] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SitemapListUsersParams]'s query parameters as `url.Values`.
func (r SitemapListUsersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
