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
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-go/internal/requestconfig"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/param"
	"github.com/stainless-sdks/alphaxiv_cat-go/packages/respjson"
)

// PaperMetadataService contains methods and other services that help with
// interacting with the alphaxiv_cat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperMetadataService] method instead.
type PaperMetadataService struct {
	options []option.RequestOption
}

// NewPaperMetadataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperMetadataService(opts ...option.RequestOption) (r PaperMetadataService) {
	r = PaperMetadataService{}
	r.options = opts
	return
}

// Get metadata for latest paper version (deprecated, used by browser extension)
//
// Source file:
// `api-server/src/controllers/v2/papers/get-paper-metadata-latest.controller.ts`
func (r *PaperMetadataService) GetLatestMetadata(ctx context.Context, upid string, query PaperMetadataGetLatestMetadataParams, opts ...option.RequestOption) (res *PaperMetadataGetLatestMetadataResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/metadata", url.PathEscape(upid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get metadata for a specific paper version (deprecated, used by browser
// extension)
//
// Source file:
// `api-server/src/controllers/v2/papers/get-paper-metadata.controller.ts`
func (r *PaperMetadataService) GetVersionMetadata(ctx context.Context, versionOrder string, params PaperMetadataGetVersionMetadataParams, opts ...option.RequestOption) (res *PaperMetadataGetVersionMetadataResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Upid == "" {
		err = errors.New("missing required upid parameter")
		return nil, err
	}
	if versionOrder == "" {
		err = errors.New("missing required versionOrder parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/papers/%s/metadata/versions/%s", url.PathEscape(params.Upid), url.PathEscape(versionOrder))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type PaperMetadataGetLatestMetadataResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperMetadataGetLatestMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperMetadataGetLatestMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperMetadataGetVersionMetadataResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperMetadataGetVersionMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperMetadataGetVersionMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperMetadataGetLatestMetadataParams struct {
	PreventTracking param.Opt[string] `query:"preventTracking,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperMetadataGetLatestMetadataParams]'s query parameters as
// `url.Values`.
func (r PaperMetadataGetLatestMetadataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaperMetadataGetVersionMetadataParams struct {
	Upid            string            `path:"upid" api:"required" json:"-"`
	PreventTracking param.Opt[string] `query:"preventTracking,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperMetadataGetVersionMetadataParams]'s query parameters
// as `url.Values`.
func (r PaperMetadataGetVersionMetadataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
