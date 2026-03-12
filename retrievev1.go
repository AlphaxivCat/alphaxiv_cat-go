// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-go/internal/requestconfig"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/AlphaxivCat/alphaxiv_cat-go/packages/param"
)

// RetrieveV1Service contains methods and other services that help with interacting
// with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetrieveV1Service] method instead.
type RetrieveV1Service struct {
	options []option.RequestOption
}

// NewRetrieveV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRetrieveV1Service(opts ...option.RequestOption) (r RetrieveV1Service) {
	r = RetrieveV1Service{}
	r.options = opts
	return
}

// Get the top AI papers
//
// Source file:
// `api-server/src/controllers/retrieve/v1/retrieve-top-papers.controller.ts`
func (r *RetrieveV1Service) GetTopPapers(ctx context.Context, query RetrieveV1GetTopPapersParams, opts ...option.RequestOption) (res *RetrieveV1GetTopPapersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "retrieve/v1/top-papers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type RetrieveV1GetTopPapersResponse = any

type RetrieveV1GetTopPapersParams struct {
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	Skip  param.Opt[string] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RetrieveV1GetTopPapersParams]'s query parameters as
// `url.Values`.
func (r RetrieveV1GetTopPapersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
