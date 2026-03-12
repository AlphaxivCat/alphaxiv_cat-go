// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alphaxivcat

import (
	"context"
	"errors"
	"fmt"
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

// PaperIngestService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperIngestService] method instead.
type PaperIngestService struct {
	options []option.RequestOption
}

// NewPaperIngestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperIngestService(opts ...option.RequestOption) (r PaperIngestService) {
	r = PaperIngestService{}
	r.options = opts
	return
}

// Ingest latest paper version if it doesn't exist (deprecated, used by browser
// extension)
//
// Source file:
// `api-server/src/controllers/v2/papers/ingest-paper-latest.controller.ts`
func (r *PaperIngestService) IngestLatest(ctx context.Context, upid string, query PaperIngestIngestLatestParams, opts ...option.RequestOption) (res *PaperIngestIngestLatestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/ingest", url.PathEscape(upid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Ingest a paper version if it doesn't exist (deprecated, used by browser
// extension)
//
// Source file: `api-server/src/controllers/v2/papers/ingest-paper.controller.ts`
func (r *PaperIngestService) IngestVersion(ctx context.Context, versionLabel string, params PaperIngestIngestVersionParams, opts ...option.RequestOption) (res *PaperIngestIngestVersionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	if versionLabel == "" {
		err = errors.New("missing required versionLabel parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/ingest/versions/%s", url.PathEscape(params.Upid), url.PathEscape(versionLabel))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type PaperIngestIngestLatestResponse struct {
	Data PaperIngestIngestLatestResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperIngestIngestLatestResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperIngestIngestLatestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperIngestIngestLatestResponseData struct {
	Authors         []any                                      `json:"authors" api:"required"`
	MaxVersionOrder float64                                    `json:"max_version_order" api:"required"`
	PdfInfo         PaperIngestIngestLatestResponseDataPdfInfo `json:"pdf_info" api:"required"`
	VerifiedAuthors []any                                      `json:"verified_authors" api:"required"`
	PaperGroup      any                                        `json:"paper_group"`
	PaperVersion    any                                        `json:"paper_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authors         respjson.Field
		MaxVersionOrder respjson.Field
		PdfInfo         respjson.Field
		VerifiedAuthors respjson.Field
		PaperGroup      respjson.Field
		PaperVersion    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperIngestIngestLatestResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperIngestIngestLatestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperIngestIngestLatestResponseDataPdfInfo struct {
	FetcherURL string `json:"fetcher_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetcherURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperIngestIngestLatestResponseDataPdfInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperIngestIngestLatestResponseDataPdfInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperIngestIngestVersionResponse struct {
	Data PaperIngestIngestVersionResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperIngestIngestVersionResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperIngestIngestVersionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperIngestIngestVersionResponseData struct {
	Authors         []any                                       `json:"authors" api:"required"`
	MaxVersionOrder float64                                     `json:"max_version_order" api:"required"`
	PdfInfo         PaperIngestIngestVersionResponseDataPdfInfo `json:"pdf_info" api:"required"`
	VerifiedAuthors []any                                       `json:"verified_authors" api:"required"`
	PaperGroup      any                                         `json:"paper_group"`
	PaperVersion    any                                         `json:"paper_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authors         respjson.Field
		MaxVersionOrder respjson.Field
		PdfInfo         respjson.Field
		VerifiedAuthors respjson.Field
		PaperGroup      respjson.Field
		PaperVersion    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperIngestIngestVersionResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperIngestIngestVersionResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperIngestIngestVersionResponseDataPdfInfo struct {
	FetcherURL string `json:"fetcher_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FetcherURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperIngestIngestVersionResponseDataPdfInfo) RawJSON() string { return r.JSON.raw }
func (r *PaperIngestIngestVersionResponseDataPdfInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperIngestIngestLatestParams struct {
	PreventTracking param.Opt[string] `query:"preventTracking,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperIngestIngestLatestParams]'s query parameters as
// `url.Values`.
func (r PaperIngestIngestLatestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperIngestIngestVersionParams struct {
	Upid            string            `path:"upid" api:"required" json:"-"`
	PreventTracking param.Opt[string] `query:"preventTracking,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperIngestIngestVersionParams]'s query parameters as
// `url.Values`.
func (r PaperIngestIngestVersionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
